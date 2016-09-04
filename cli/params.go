package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

const (
	uploadFileKey = "uploadFile"
	envHostKey    = "RIKKA_HOST"
	envPwdKey     = "RIKKA_PWD"
)

var (
	argHost = flag.String("t", "", "rikka host, will override env variable if not empty")
	argPwd  = flag.String("p", "", "rikka password, will override env variable if not empty")
)

func getHost() string {

	var host string

	if *argHost != "" {
		host = *argHost
		l.Info("Get host from argument:", host)
	} else {
		l.Debug("No host argument provided, try get from env")

		host = os.Getenv(envHostKey)
		if host == "" {
			l.Fatal("No", envHostKey, "env variable, I don't know where to upload")
		}
		l.Info("Get Rikka host from env:", host)
	}

	if !strings.HasPrefix(host, "http") {
		host = "http://" + host
		l.Debug("Add scheme http:// for host, becauses:", host)
	}
	l.Debug("Host seems contains scheme, won't process")

	urlStruct, err := url.Parse(host)
	if err != nil || urlStruct.Host == "" || urlStruct.Scheme == "" || urlStruct.Path != "" {
		l.Fatal("Invalid Rikka host:", host)
	}
	l.Debug("Host check passed, struct:", fmt.Sprintf("%+v", *urlStruct))

	return urlStruct.Scheme + "://" + urlStruct.Host
}

func getPassword() string {

	var password string

	if *argPwd != "" {
		password = *argPwd
		l.Info("Get password from argument:", password)
	} else {
		l.Debug("No password argument provided, try get from env")

		password = os.Getenv(envPwdKey)
		if password == "" {
			l.Fatal("No", envPwdKey, "env variable, I don't know password of rikka")
		}
		l.Info("Get password from env variable:", password)
	}

	return password
}

func getParams() map[string]string {
	params := map[string]string{
		"from":     "api",
		"password": getPassword(),
	}

	l.Debug("Build params:", params)

	return params
}
