. ../common.sh

nodeuri(){
  echo "http://node-0.${DOMAIN:-$(minikube ip).nip.io}:80"
}

akash(){
  _akash -n "$(nodeuri)" "$@"
}

akashd(){
  _akashd "$@"
}
