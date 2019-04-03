package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

const (
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
		l.Debug("Add scheme http:// for host, become:", host)
	} else {
		l.Debug("Host seems contains scheme, won't process")
	}

	if strings.HasSuffix(host, "/") {
		host = host[:len(host)-1]
		l.Debug("Delete extra / for host, become:", host)
	} else {
		l.Debug("No extra / in host, won't process")
	}

	urlObj, err := url.Parse(host)
	if err != nil || urlObj.Host == "" || urlObj.Scheme == "" || urlObj.Path != "" {
		l.Fatal("Invalid Rikka host:", host)
	}
	//noinspection GoNilness because l.Fatal will exit
	l.Debug("Host check passed, struct:", fmt.Sprintf("%+v", *urlObj))

	//noinspection GoNilness, ditto
	return urlObj.Scheme + "://" + urlObj.Host
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
