#!/bin/sh
go build ./
kill $(pgrep rms)
./rms > /dev/null 2>&1 & 
echo "running on prccedd id :" + $(pgrep rms)