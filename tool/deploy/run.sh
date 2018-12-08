#!/bin/bash

set -eu

scp assets/janken.tar.gz sandbox-web:/tmp/
scp task/deploy.sh sandbox-web:/tmp/
ssh sandbox-web /tmp/deploy.sh $(git tag)
