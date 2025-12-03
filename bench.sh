#!/bin/bash

function bench() {
  pushd "$1" && go build && ./day*
  popd || exit 1
}

for d in day*; do
  bench "$d"
done
