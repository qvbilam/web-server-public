#!/bin/bash
# shellcheck disable=SC2086

tag=$1
imageName=qvbilam/oss-web-server-alpine
originImageName=registry.cn-hangzhou.aliyuncs.com/qvbilam/oss-web

# build image
docker build -t ${imageName} .
# login hub
docker login --username=13501294164 registry.cn-hangzhou.aliyuncs.com
# tag image
docker tag ${imageName} ${originImageName}:${tag}
# push image
docker push ${originImageName}:${tag}