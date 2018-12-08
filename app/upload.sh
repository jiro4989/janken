#!/bin/bash

set -eux

tar czf janken.tar.gz bin conf shell template
scp janken.tar.gz sandbox-web:/tmp/
