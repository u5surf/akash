#!/usr/bin/env bash

set -exo pipefail

HYPERKUBE_VERSION=1.10.1
PUPERNETES_VERSION=0.5.0

install(){
  # install systemd
  # XXX: only works from travis.yml
  # sudo apt-get install -yq systemd

  # download, install pupernetes, kubectl, helm
  curl -L https://github.com/DataDog/pupernetes/releases/download/v${PUPERNETES_VERSION}/pupernetes -o ./pupernetes
  chmod +x ./pupernetes

  # download, install kubectl
  sudo curl -L https://storage.googleapis.com/kubernetes-release/release/v${HYPERKUBE_VERSION}/bin/linux/amd64/kubectl -o /usr/local/bin/kubectl
  sudo chmod +x /usr/local/bin/kubectl

  # download, install helm
  curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get | sudo bash
}

boot(){
  # install kube
  sudo ./pupernetes daemon run /tmp/sandbox -v 4 --job-type systemd --bind-address 0.0.0.0:8989 --hyperkube-version $HYPERKUBE_VERSION

  # wait for system ready
  sudo ./pupernetes wait --wait-timeout 5m

  # initialize helm
  helm init

  kubectl create -f "$(dirname $0)/../_run/multi/rbac.yml"
  kubectl create -f "$(dirname $0)/ingress.yml"
}

install
boot
