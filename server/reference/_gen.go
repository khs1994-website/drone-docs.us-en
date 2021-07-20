package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	for _, v := range variables {
		name := strings.ToLower(v + ".md")
		name = strings.ReplaceAll(name, "_", "-")
		data := fmt.Sprintf(markdownF, v, v)
		ioutil.WriteFile(name, []byte(data), 0644)
	}
}

var markdownF = `---
date: 2000-01-01T00:00:00+00:00
title: %s
author: bradrydzewski
---



`+"```"+`
%s=
`+"```"+`
`

var variables = []string{
	"DRONE_LICENSE",
	//
	"DRONE_GIT_ALWAYS_AUTH",
	"DRONE_GIT_USERNAME",
	"DRONE_GIT_PASSWORD",
	//
	"DRONE_CRON_DISABLED",
	"DRONE_CRON_INTERVAL" ,
	//
	"DRONE_DATABASE_DRIVER",
	"DRONE_DATABASE_DATASOURCE",
	"DRONE_DATABASE_SECRET",
	//
	"DRONE_JSONNET_ENABLED",
	//
	"DRONE_LICENSE",
	//
	"DRONE_LOGS_DEBUG",
	"DRONE_LOGS_TRACE",
	"DRONE_LOGS_COLOR",
	"DRONE_LOGS_PRETTY",
	"DRONE_LOGS_TEXT",
	//
	"DRONE_PROMETHEUS_ANONYMOUS_ACCESS",
	//
	"DRONE_REPOSITORY_FILTER",
	//
	"DRONE_RPC_SERVER",
	"DRONE_RPC_SECRET",
	"DRONE_RPC_DEBUG",
	"DRONE_RPC_HOST",
	"DRONE_RPC_PROTO",
	"DRONE_RPC_EXTRA_HOSTS",
	//
	"DRONE_AGENTS_ENABLED",
	//
	"DRONE_SERVER_HOST" ,
	"DRONE_SERVER_PORT" ,
	"DRONE_SERVER_PROTO",
	//
	"DRONE_TLS_AUTOCERT",
	"DRONE_TLS_CERT",
	"DRONE_TLS_KEY",
	//
	"DRONE_SERVER_PROXY_HOST",
	"DRONE_SERVER_PROXY_PROTO",
	"DRONE_REGISTRATION_CLOSED",
	//
	"DRONE_AUTHENTICATION_ENDPOINT",
	"DRONE_AUTHENTICATION_TOKEN",
	"DRONE_AUTHENTICATION_SKIP_VERIFY",
	//
	"DRONE_COOKIE_TIMEOUT",
	"DRONE_COOKIE_SECRET",
	"DRONE_COOKIE_SECURE",
	//
	"DRONE_STATUS_DISABLED",
	"DRONE_STATUS_NAME",
	//
	"DRONE_USER_CREATE",
	"DRONE_USER_FILTER",
	//
	"DRONE_WEBHOOK_EVENTS",
	"DRONE_WEBHOOK_ENDPOINT",
	"DRONE_WEBHOOK_SECRET",
	"DRONE_WEBHOOK_SKIP_VERIFY",
	//
	"DRONE_YAML_ENDPOINT",
	"DRONE_YAML_SECRET",
	"DRONE_YAML_SKIP_VERIFY",
	//
	"DRONE_BITBUCKET_CLIENT_ID",
	"DRONE_BITBUCKET_CLIENT_SECRET",
	"DRONE_BITBUCKET_SKIP_VERIFY",
	"DRONE_BITBUCKET_DEBUG",
	//
	"DRONE_GITEA_SERVER",
	"DRONE_GITEA_CLIENT_ID",
	"DRONE_GITEA_CLIENT_SECRET",
	"DRONE_GITEA_SKIP_VERIFY",
	"DRONE_GITEA_SCOPE",
	"DRONE_GITEA_DEBUG",
	//
	"DRONE_GITHUB_SERVER",
	"DRONE_GITHUB_API_SERVER",
	"DRONE_GITHUB_CLIENT_ID",
	"DRONE_GITHUB_CLIENT_SECRET",
	"DRONE_GITHUB_SKIP_VERIFY",
	"DRONE_GITHUB_SCOPE" ,
	"DRONE_GITHUB_USER_RATELIMIT",
	"DRONE_GITHUB_DEBUG",
	//
	"DRONE_GITLAB_SERVER",
	"DRONE_GITLAB_CLIENT_ID",
	"DRONE_GITLAB_CLIENT_SECRET",
	"DRONE_GITLAB_SKIP_VERIFY",
	"DRONE_GITLAB_DEBUG",
	//
	"DRONE_GOGS_SERVER",
	"DRONE_GOGS_SKIP_VERIFY",
	"DRONE_GOGS_DEBUG",
	//
	"DRONE_STASH_SERVER",
	"DRONE_STASH_CONSUMER_KEY",
	"DRONE_STASH_CONSUMER_SECRET",
	"DRONE_STASH_PRIVATE_KEY",
	"DRONE_STASH_SKIP_VERIFY",
	"DRONE_STASH_DEBUG",
	//
	"DRONE_S3_BUCKET",
	"DRONE_S3_PREFIX",
	"DRONE_S3_ENDPOINT",
	"DRONE_S3_PATH_STYLE",
	"DRONE_AZURE_BLOB_CONTAINER_NAME",
	"DRONE_AZURE_STORAGE_ACCOUNT_NAME",
	"DRONE_AZURE_STORAGE_ACCESS_KEY",
	"DRONE_HTTP_ALLOWED_HOSTS",
	"DRONE_HTTP_PROXY_HEADERS",
	"DRONE_HTTP_SSL_REDIRECT",
	"DRONE_HTTP_SSL_TEMPORARY_REDIRECT",
	"DRONE_HTTP_SSL_HOST",
	"DRONE_HTTP_SSL_PROXY_HEADERS",
	"DRONE_HTTP_STS_SECONDS",
	"DRONE_HTTP_STS_INCLUDE_SUBDOMAINS",
	"DRONE_HTTP_STS_PRELOAD",
	"DRONE_HTTP_STS_FORCE_HEADER",
	"DRONE_HTTP_BROWSER_XSS_FILTER",
	"DRONE_HTTP_FRAME_DENY",
	"DRONE_HTTP_CONTENT_TYPE_NO_SNIFF",
	"DRONE_HTTP_CONTENT_SECURITY_POLICY",
	"DRONE_HTTP_REFERRER_POLICY",
}