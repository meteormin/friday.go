#!/bin/bash

BASEDIR=$(dirname "$0")

tag=$1
buildArgs=$2
platform=$3

dockerImageDir="$BASEDIR/../.docker"
imageName="prs-api"
sourcePath="$BASEDIR/../prs"

echo "[Docker image build] tag=$tag"
echo "buildArgs=$buildArgs"

if [ -z "$platform" ]; then
  os=$(uname -s | tr '[:upper:]' '[:lower:]')
  arch=$(uname -m| tr '[:upper:]' '[:lower:]')

  echo "Detected OS: $os"
  echo "Detected ARCH: $arch"
  platform="linux/$arch"

  echo "Set Default Platform: $platform"
else
  echo "platform=$platform"
fi

if [ -d "$dockerImageDir/../.docker" ]; then
  rm -rf "$dockerImageDir/*.tar"
else
  mkdir "$dockerImageDir"
fi

if [ -z "$tag" ]; then
  tag="latest"
fi

echo "* start docker build"
if [ "$platform" = "" ]; then
  docker build -t $imageName:$tag \
  --build-arg="$buildArgs" \
  -f $sourcePath/Dockerfile $sourcePath
else
  docker buildx build -t $imageName:$tag \
  --platform="$platform" \
  --build-arg="$buildArgs" \
  -f $sourcePath/Dockerfile $sourcePath
fi