#!/usr/bin/env bats

load env
load util

@test "deploy: submit deployment" {
  skip
  akash deployment create deployment.yml -k master -w
}

@test "deploy: query order" {
  did=$(akash query deployment | jq -er '.items[0].address')
  for count in {0..4}; do
    run check_order $did
    if [ "$status" -eq 0 ]; then
      break
    fi
    sleep 1
  done
  [ "$status" -eq 0 ]
}

@test "deploy: query lease" {
  did=$(akash query deployment | jq -er '.items[0].address')
  for count in {0..4}; do
    run check_lease $did
    if [ "$status" -eq 0 ]; then
      break
    fi
    sleep 1
  done
  [ "$status" -eq 0 ]
}
