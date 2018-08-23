#!/bin/bash

/usr/sbin/init &

sleep 2s
go run /Golang/src/glusterfs-init/main.go

sleep inf
