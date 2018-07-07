#!/usr/bin/env bats

@test "transfer: check source key balance" {
  balance=$(./run.sh query | jq '.balance')
  [ "$balance" -eq 1000000000 ]
}

@test "transfer: send tokens" {
  run ./run.sh send
  [ "$status" -eq 0 ]
}

@test "transfer: check target account balance" {
  balance=$(./run.sh query other | jq '.balance')
  [ "$balance" -eq 100 ]
}

@test "transfer: check source account balance" {
  balance=$(./run.sh query | jq '.balance')
  [ "$balance" -eq 999999900 ]
}
