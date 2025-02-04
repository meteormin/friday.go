#!/bin/bash

BASEDIR=$(dirname "$0")

tag=$1

imageName="prs-api"
imagePath="$BASEDIR/../.docker"
echo "[docker save] tag=$tag"
echo "imageName=$imageName"

docker save -o $imagePath/prs-api-$tag.tar $imageName:$tag