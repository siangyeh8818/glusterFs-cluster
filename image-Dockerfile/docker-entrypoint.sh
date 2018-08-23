#!/bin/bash

#mkdir -p /12345
sleep 2s
echo "run init script"
#systemctl status glusterd.service 
go run /Golang/src/glusterfs-init/main.go
