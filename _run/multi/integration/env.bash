export DATA_DIR=$BATS_TEST_DIRNAME/../data
export AKASH_DATA=$DATA_DIR/client

nodeuri(){
  echo "http://node-0.$(minikube ip).nip.io:80"
}

akash(){
  AKASH_NODE=$(nodeuri) $BATS_TEST_DIRNAME/../../../akash "$@"
}
