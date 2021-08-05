#!/bin/bash
set -ex
cd $( dirname $0 )

(
  cd ..
  go build
  go install
)

go generate .
