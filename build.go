package main

import (
	"log"
	"os/exec"
	"path/filepath"
)

func buildACI(dir string) error {
	imageName, err := filepath.Abs(dir)
        if err != nil {
                log.Fatalf("err: %v", err)
        }
	imageName += ".aci"
	return createACI(dir, imageName)
}

func createACI(dir string, imageName string) error {
	
	if err := run(exec.Command("actool", "build", "-overwrite", dir, imageName)); err != nil {
		return err
	}
	return nil
}
