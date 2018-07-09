#!/usr/bin/env bats

load util

@test "setup: initialize minikube" {
  if minikube status; then
    skip "minikube already running"
  fi

  minikube start --cpus "${CPU:-4}" --memory "${MEMORY:-4096}"

  for count in {0..4}; do
    run minikube status
    if [ "$status" -eq 0 ]; then
      break
    fi
    sleep 1
  done
  [ "$status" -eq 0 ]

  minikube addons enable ingress

  kubectl create -f rbac.yml

  helm init
}

@test "setup: build images" {
  make image-minikube

  minikube_wait_image ovrclk/akash
  minikube_wait_image ovrclk/akashd
}

@test "deploy: create akash configuration" {
  ./run.sh init
}

@test "deploy: install nodes" {
  make helm-install-nodes

  for count in {0..15}; do
    run make helm-check-nodes
    if [ "$status" -eq 0 ]; then
      break
    fi
    sleep 1
  done
  [ "$status" -eq 0 ]
}

@test "deploy: install providers" {
  make helm-install-providers

  for count in {0..15}; do
    run make helm-check-providers
    if [ "$status" -eq 0 ]; then
      break
    fi
    sleep 1
  done
  [ "$status" -eq 0 ]
}
