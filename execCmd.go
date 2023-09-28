package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func clone(response []RepoInfos) {
	folderName := "archive"
	_, err := os.Stat(folderName)
	if !os.IsNotExist(err) {
		err := os.RemoveAll(folderName)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(response); i++ {
		path := fmt.Sprintf("%s/%s", folderName, response[i].Name)
		cmd := exec.Command("git", "clone", response[i].Url, path)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
