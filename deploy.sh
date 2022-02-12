#!/bin/bash
alias ll="ls -lav"
git pull
go mod tidy
# go run mymod