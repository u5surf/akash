
minikube_docker(){
  ( eval minikube docker-env; docker "$@" )
}

minikube_image_exists() {
  minikube_docker image inspect $1
}

minikube_wait_image() {
  for count in {0..5}; do
    if minikube_image_exists $1; then
      return 0
    fi
    sleep 1
  done
  return 1
}

check_order(){
  akash query order | jq -e ".items[].id | select(.deployment==\"$1\")"
}

check_lease(){
  akash query order | jq -e ".items[].id | select(.deployment==\"$1\")"
}
