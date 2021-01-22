#!/bin/sh 
# This file will build the binary for linux_i386 & linux_amd64
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
go get github.com/karalabe/xgo
xgo --targets="darwin/amd64,linux/amd64,linux/386" github.com/iamstefin/glitzy