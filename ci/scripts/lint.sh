#!/bin/bash -eux

pushd dis-design-system-go
  make prepare-node
  DEFAULT_WORKSPACE=$(pwd)
  export DEFAULT_WORKSPACE
  bash -ex /entrypoint.sh
popd
