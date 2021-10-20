#!/bin/bash -x

COMP_NAME="wechatbot"
BUILD_PATH=$(cd `dirname $0`;pwd)
WECHATBOT_PATH=$BUILD_PATH/..
DOCKER_TAG="wechatbot"
DOCKER_IMAGE="wechatbot.tar"

# build wechatbot
cd $WECHATBOT_PATH
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -o $COMP_NAME

# build docker image
cd $BUILD_PATH
mv $WECHATBOT_PATH/$COMP_NAME $BUILD_PATH
docker build -t $DOCKER_TAG .
docker save -o $DOCKER_IMAGE $DOCKER_TAG
rm $BUILD_PATH/$COMP_NAME