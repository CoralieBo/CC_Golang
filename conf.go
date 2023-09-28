package main

import (
	"fmt"
	"log"
	"os"
)

const (
	userNameKey  = "USER_NAME"
	userTypeKey  = "USER_TYPE"
	userTokenKey = "USER_TOKEN"
)

type Config struct {
	UserName  string
	UserType  string
	UserToken string
}

func NewConfig() Config {
	userName, ok := os.LookupEnv(userNameKey)
	if !ok || userName == "" {
		log.Fatal(userNameKey)
	} else {
		fmt.Println(userNameKey, ":", userName)
	}

	userType, ok := os.LookupEnv(userTypeKey)
	if !ok || userType == "" {
		log.Fatal(userTypeKey)
	} else {
		fmt.Println(userTokenKey, ":", userType)
	}

	userToken, ok := os.LookupEnv(userTokenKey)
	if !ok || userToken == "" {
		log.Fatal(userTokenKey)
	} else {
		fmt.Println(userTokenKey, ":", userToken)
	}

	return Config{
		UserName:  userName,
		UserType:  userType,
		UserToken: userToken,
	}
}
