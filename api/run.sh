export MONGODB_URL="mongodb://root:root@127.0.0.1:27017/"

function test() {
    if [[ -z "${1}" ]];then
        echo "missing arg: directory to run test"
    else
        go test -v $1/... -cover
    fi
}

container=$(docker ps | grep johnson-brooks-db)
[[ -z "${container}" ]] \
    && docker compose up mongo -d &>/dev/null \
    && container=$(docker ps | grep johnson-brooks-db)

case $1 in 
    "mock")
        go run ./scripts/mock/main.go
        ;;
    "test")
        test $2
        ;;
    *)
        echo "usage: [mock|test]"
        ;;
esac

[[ -n "${container}" ]] && docker compose down &> /dev/null
