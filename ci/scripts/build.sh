#!/bin/bash -eux


pushd dis-design-system-go
  make build
  TAG=`git tag --points-at HEAD | grep ^v | head -n 1`
  
  if [ -z "${TAG}" ]; then
    echo "No tag found, using short commit hash"
    TAG=`git rev-parse --short HEAD`
  fi

popd

mkdir build/$TAG
cp -a dis-design-system-go/dist/assets/. build/$TAG
