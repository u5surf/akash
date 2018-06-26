#!/usr/bin/env bash

set -exo pipefail

export DOMAIN=127.0.0.1.nip.io

(
  cd $(dirname $0)../_run/multi
  make install-all-baremetal
)

(
  cd $(dirname $0)../_integration
  make kube-run
)
