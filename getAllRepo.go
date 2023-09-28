package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"time"
)

type RepoInfos struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Full_name string `json:"full_name"`
	Owner     struct {
		Login string `json:"login"`
		Id    int    `json:"id"`
	} `json:"owner"`
	Description string    `json:"description"`
	Url         string    `json:"clone_url"`
	Api_url     string    `json:"url"`
	Language    string    `json:"language"`
	Updated_at  time.Time `json:"updated_at"`
}

func getAllRepo(infos Config) []RepoInfos {
	url := fmt.Sprintf("https://api.github.com/%s/%s/repos", infos.UserType, infos.UserName)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	result := parseReposInfos(string(body))
	return result
}

func parseReposInfos(body string) []RepoInfos {
	var result []RepoInfos
	err := json.Unmarshal([]byte(body), &result)
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Updated_at.After(result[j].Updated_at)
	})
	return result
}
