#!/bin/bash -eux

go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.5.0

pushd dis-design-system-go
  make lint
popd
