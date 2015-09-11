//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"strings"
	"os/exec"
	"io/ioutil"
	"path/filepath"
	"encoding/json"
	"github.com/opencontainers/specs"
	"github.com/appc/spec/schema/types"
	"github.com/appc/spec/schema"
)

func runOCI2ACI(path string, flagDebug bool) error {
	if flagDebug {
		InitDebug()
	}
	if bValidate := validateOCIProc(path); bValidate != true {
                fmt.Println("Conversion stop.")
                return nil
        }

	dirWork := createWorkDir()

	convertProc(path, dirWork)

	err := buildACI(dirWork)
	if err != nil {
		if debugEnabled {
			fmt.Println("Generate aci image failed!")
		}
	} else {
		if debugEnabled {
                        fmt.Println("aci image generated successfully.")
                }
	}

	return nil
}

func createWorkDir() string {
	idir, err := ioutil.TempDir("", "oci2aci")
        if err != nil {
                return ""
        }
        rootfs := filepath.Join(idir, "rootfs")
        os.MkdirAll(rootfs, 0755)

	data := []byte{}
	if err := ioutil.WriteFile(filepath.Join(idir, "manifest"), data, 0644); err != nil {
		return ""
	}
	return idir
}

func genManifest(path string) *schema.ImageManifest {
	runtimePath := path + "/runtime.json"
	configPath := path + "/config.json"

	runtime, err := ioutil.ReadFile(runtimePath)
	if err != nil {
		return nil
	}
	
	config, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil
	}
	
	var spec specs.Spec
	err = json.Unmarshal(config, &spec)
	if err != nil {
		return nil
	}

	var runSpec specs.RuntimeSpec
	err = json.Unmarshal(runtime, &runSpec)
	if err != nil {
		return nil
	}

	m := new(schema.ImageManifest)

	// Assemble "acKind", "acVersion" and "name" fields
	m.ACVersion = schema.AppContainerVersion
	m.ACKind = "ImageManifest"
	m.Name = "example"
	
	// Assemble "app" field
	app := new(types.App)

	app.Exec = spec.Process.Args
	app.User = string(spec.Process.User.UID)
	app.Group = string(spec.Process.User.GID)

	event := new(types.EventHandler)
	event.Name = "pre-start"
	for index := range runSpec.Hooks.Prestart {
		event.Exec = append(event.Exec, runSpec.Hooks.Prestart[index].Path)
		event.Exec = append(event.Exec, runSpec.Hooks.Prestart[index].Args...)
		event.Exec = append(event.Exec, runSpec.Hooks.Prestart[index].Env...)
        }
	app.EventHandlers = append(app.EventHandlers, *event)

	event = new(types.EventHandler)
        event.Name = "post-stop"
        for index := range runSpec.Hooks.Poststop {
                event.Exec = append(event.Exec, runSpec.Hooks.Poststop[index].Path)
                event.Exec = append(event.Exec, runSpec.Hooks.Poststop[index].Args...)
                event.Exec = append(event.Exec, runSpec.Hooks.Poststop[index].Env...)
        }
        app.EventHandlers = append(app.EventHandlers, *event)
	
	app.WorkingDirectory = spec.Process.Cwd

	env := new(types.EnvironmentVariable)
	for index := range spec.Process.Env {
                s := strings.Split(spec.Process.Env[index], "=")
                env.Name = s[0]
		env.Value = s[1]
		app.Environment = append(app.Environment, *env)
        }
	
	port := new(types.Port)
	port.Name = types.ACName("test")
	port.Protocol = "tcp"
	port.Port = 4000
	port.Count = 0
	port.SocketActivated = true
	app.Ports = append(app.Ports, *port)

	for index := range spec.Mounts {
		mount := new(types.MountPoint)
		mount.Name = types.ACName(spec.Mounts[index].Name)
		mount.Path = spec.Mounts[index].Path
		app.MountPoints = append(app.MountPoints, *mount)
	}
	
	m.App = app
	
	// Assemble "labels" field
	label := new(types.Label)
        label.Name = types.ACIdentifier("version")
        label.Value = spec.Version
        m.Labels = append(m.Labels, *label)

	label = new(types.Label)
	label.Name = types.ACIdentifier("os")
	label.Value = spec.Platform.OS
	m.Labels = append(m.Labels, *label)

	label = new(types.Label)
	label.Name = types.ACIdentifier("arch")
	label.Value = spec.Platform.Arch
        m.Labels = append(m.Labels, *label)

	
	return m
}

func convertProc(srcPath, dstPath string) error {
	src, _ := filepath.Abs(srcPath)
	src += "/rootfs"
	if err := run(exec.Command("cp", "-rf", src, dstPath)); err != nil {
		return err
	}

	m := genManifest(srcPath)
	
	bytes, err := json.Marshal(m)
	if err != nil {
		if debugEnabled {		
			fmt.Println("json Marshal exec failed!")
		}
		return err
	}
	
	manifestPath := dstPath + "/manifest"

	ioutil.WriteFile(manifestPath, bytes, 0644)
	return nil
}
