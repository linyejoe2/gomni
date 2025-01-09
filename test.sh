#!/bin/bash

go build -o gomni

./gomni ssh add 172.16.0.201 -n dae01 -u administrator -p 1qaz@WSX3edc
