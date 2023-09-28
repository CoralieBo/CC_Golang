package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

type Commit struct {
	Sha string `json:"sha"`
}

func pullAllRepos(repos []RepoInfos) {
	for i := 0; i < len(repos); i++ {
		commit := getLastCommit(repos[i].Api_url)
		pullRepo(commit, repos[i].Name)
	}
}

func pullRepo(commit Commit, repoName string) {
	path := fmt.Sprintf("archive/%s", repoName)
	cmd := exec.Command("git", "pull", "origin", commit.Sha)
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func getLastCommit(api_url string) Commit {
	url := fmt.Sprintf("%s/commits", api_url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result []Commit
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(result[0])
	return result[0]
}
