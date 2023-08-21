#!/usr/bin/env bash
if [ -f cloud-disk ]; then
  rm -f cloud-disk
fi

export GOOS=linux

res=$(go env | grep GOOS=)
echo "$res"

go build -o cloud-disk core.go

export GOOS=windows

res=$(go env | grep GOOS=)
echo "$res"
