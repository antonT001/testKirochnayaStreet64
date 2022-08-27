#!/bin/bash
sleep 15

cd /var/www/service/
go run cmd/service/main.go &

while true ; do sleep 5; done;
