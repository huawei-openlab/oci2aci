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
	"io/ioutil"
	"path/filepath"
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

func convertProc(srcPath, dstPath string) error {
	return nil
}
