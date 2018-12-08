#!/bin/bash

set -eux

tag=$1
if [ -z "$tag" ]; then
  echo "Neet tag args."
  exit 1
fi

gzfile=janken.tar.gz
release_dir=/opt/release

sudo mv "/tmp/$gzfile" "$release_dir/"
sudo tar xzf "$release_dir/$gzfile" -C "$release_dir"
sudo chown -R janken:janken "$release_dir"

current_dir="$release_dir/janken/$tag"
sudo ln -sfn "$current_dir" /opt/janken
