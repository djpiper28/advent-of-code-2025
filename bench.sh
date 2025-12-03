#!/bin/bash

function bench() {
  pushd "$1" && go build && ./day*
  popd
}

for d in day*; do
  bench "$d"
done
