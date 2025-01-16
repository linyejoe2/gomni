#!/bin/bash

env CGO_ENABLED=0 go build -ldflags "-X github.com/linyejoe2/gomni/cmd.version=$(git describe --tags --abbrev=0)" -o gomni
