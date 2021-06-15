#!/bin/bash

set -e

if [ $# -eq 0 ] ; then
	echo "Usage: ./update.sh v#.#.#"
	exit
fi

VERSION=$1

docker build -t red-builder -f ./Dockerfile.all .

ID=$(docker create red-builder)

rm -rf out
mkdir -p out
docker cp $ID:/build/pkg/linux-amd64/redirect ./out/redirect-$VERSION-linux-amd64
docker cp $ID:/build/pkg/linux-i386/redirect ./out/redirect-$VERSION-linux-386
docker cp $ID:/build/pkg/linux-arm6/redirect ./out/redirect-$VERSION-linux-arm6
docker cp $ID:/build/pkg/linux-arm7/redirect ./out/redirect-$VERSION-linux-arm7
docker cp $ID:/build/pkg/linux-arm64/redirect ./out/redirect-$VERSION-linux-arm64
docker cp $ID:/build/pkg/darwin-amd64/redirect ./out/redirect-$VERSION-darwin-amd64
docker cp $ID:/build/pkg/win-amd64/redirect.exe ./out/redirect-$VERSION-windows-amd64.exe

docker rm -f $ID
docker rmi red-builder

docker buildx build --push --platform linux/arm/v7,linux/arm64/v8,linux/amd64 --tag halverneus/https-redirect:$VERSION .
docker buildx build --push --platform linux/arm/v7,linux/arm64/v8,linux/amd64 --tag halverneus/https-redirect:latest .

echo "Done"