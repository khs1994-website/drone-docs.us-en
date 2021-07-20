---
date: 2000-01-01T00:00:00+00:00
title: Docker
title_in_header: Example Docker Configuration
author: bradrydzewski
weight: 1
toc: false
---

This guide covers configuring continuous integration pipelines for projects that have a Docker dependency. If you're new to Drone please read our Tutorial and build configuration guides first.

# Basic Example

In the below example we demonstrate a pipeline that connects to the host machine Docker daemon by mounting a volume. For security reasons, only trusted repositories can mount volumes. Furthermore, mounting the host machine Docker socket is highly insecure, and should only be used in trusted environments. _Unlike docker pipelines, on kubernetes you cannot mount files or sockets, you need to mount folders._

```
---
kind: pipeline
type: kubernetes
name: default

steps:
- name: test
  image: docker:dind
  volumes:
  - name: dockersock
    path: /var/run/
  commands:
  - docker ps -a

volumes:
- name: dockersock
  host:
    path: /var/run/

...
```
