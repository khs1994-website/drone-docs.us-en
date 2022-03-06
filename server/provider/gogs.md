---
date: 2000-01-01T00:00:00+00:00
title: Gogs
author: bradrydzewski
weight: 4
steps: true
toc: true
aliases:
- /installation/providers/gogs
- /install-for-gogs
- /admin/setup-gogs
- /installation/gogs/
- /installation/gogs/single-machine/
- /installation/gogs/multi-machine/
- /installation/gogs/kubernetes/
description: |
  Installation and configure for use with Gogs
---

This article explains how to install the Drone server for Gogs. The server is packaged as a minimal Docker image distributed on [DockerHub](https://hub.docker.com/r/drone/drone).

<div class="alert alert-danger">
Please note we strongly recommend using Gitea with Drone. Gitea has better compatibility with Drone and certain features may not work with Gogs.
</div>

# Download

The Drone server is distributed as a lightweight Docker image. The image is self-contained and does not have any external dependencies.

```
$ docker pull drone/drone:2
```

# Configuration

The Drone server is configured using environment variables. This article references a subset of configuration options, defined below. See [Configuration]({{< relref "../reference/_index.md" >}}) for a complete list of configuration options.

* `DRONE_GOGS_SERVER`
  : Required string value provides your Gogs server address. For example `https://gogs.company.com`.
* `DRONE_GIT_ALWAYS_AUTH`
  : Optional boolean value configures Drone to authenticate when cloning public repositories.
* `DRONE_RPC_SECRET`
  : Required string value provides the drone shared secret. This is used to authenticate the rpc connection to the server. The server and agent must be provided the same secret value.
* `DRONE_SERVER_HOST`
  : Required string value provides your external hostname or IP address. If using an IP address you may include the port. For example `drone.company.com`.
* `DRONE_SERVER_PROTO`
  : Required string value provides your external protocol scheme. This value should be set to http or https. This field defaults to https if you configure ssl or acme.

# Start the Server

The server container can be started with the below command. The container is configured through environment variables. _Remember to replace the placeholder values below with the appropriate values._

{{< highlight bash "linenos=table,hl_lines=3-7" >}}
docker run \
  --volume=/var/lib/drone:/data \
  --env=DRONE_AGENTS_ENABLED=true \
  --env=DRONE_GOGS_SERVER=https://gogs.company.com \
  --env=DRONE_RPC_SECRET=super-duper-secret \
  --env=DRONE_SERVER_HOST=drone.company.com \
  --env=DRONE_SERVER_PROTO=https \
  --publish=80:80 \
  --publish=443:443 \
  --restart=always \
  --detach=true \
  --name=drone \
  drone/drone:2
{{< / highlight >}}

# Install Runners

Once your server is up and running you will need to install runners to execute your build pipelines. See our runner installation documentation for detailed installation instructions. 

{{< link "/runner/overview" "Install Runners" >}}

# Login

You can login to your Drone server by visiting the server address in your browser. You will need to authenticate using your Gogs username and password. _This is required because Gogs does not support oauth._
