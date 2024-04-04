#!/bin/bash
archs=(amd64 arm64)
oss=(linux darwin)

for arch in "${archs[@]}"
do
  for os in "${oss[@]}"
  do
    env GOOS="${os}" GOARCH="${arch}" go build -o ./dist/gateway-"${os}"-"${arch}"
  done
done
