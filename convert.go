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
	"os/exec"
	"io/ioutil"
	"path/filepath"
	"encoding/json"
)

type ImageManifest struct {
	ACVersion     string `json:"acVersion"`
	ACKind        string `json:"acKind"`
	Name          string `json:"name"`
	Labels        string `json:"labels,omitempty"`
	App           string `json:"app,omitempty"`
	Annotations   string `json:"annotations,omitempty"`
	Dependencies  string `json:"dependencies,omitempty"`
	PathWhitelist string `json:"pathWhitelist,omitempty"`
}

func blankImageManifest() *ImageManifest {
	return &ImageManifest{ACKind: "ImageManifest", ACVersion: "0.6.1+git", Name: "test"}
}

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

	buildACI(dirWork)

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

func genManifest() *ImageManifest {
	return blankImageManifest()
}

func convertProc(srcPath, dstPath string) error {
	src := srcPath + "/rootfs"
	if err := run(exec.Command("cp", "-rf", src, dstPath)); err != nil {
		return err
	}

	m := genManifest()
	fmt.Println(m)
	
	bytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	
	manifestPath := dstPath + "/manifest"

	ioutil.WriteFile(manifestPath, bytes, 0644)
	return nil
}
