#!/bin/bash

set -eux

tag=$(git tag)
apl_dir="janken/$tag"
mkdir -p "$apl_dir"
cp -r ../../app/{bin,conf,shell,template} "$apl_dir/"
tar czf janken.tar.gz janken

mkdir -p assets
mv janken.tar.gz assets/
rm -rf janken
