#!/bin/bash

until=${1:-"until=2h"}

echo "[Clean up images] '<none>'"
docker image rm $(docker image list -f 'dangling=true' -q --no-trunc)
echo ""

echo "[Clean up builder] until 2h"
docker builder prune --filter $until
echo ""