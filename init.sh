#!/bin/sh
go build ./
kill $(pgrep rms)
./rms > /dev/null 2>&1 & 
disown
echo "running on prccedd id :" + $(pgrep rms)

