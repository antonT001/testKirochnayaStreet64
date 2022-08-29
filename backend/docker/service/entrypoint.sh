#!/bin/bash
sleep 25

cd /var/www/service/
go run cmd/service/main.go &

while true ; do sleep 5; done;
