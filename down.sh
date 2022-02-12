#!/bin/bash
alias ll="ls -lav"
git pull
echo "PULL"
go mod tidy
echo "tidy"
go run mymod