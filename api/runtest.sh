export MONGODB_URL="mongodb://root:root@127.0.0.1:27017/"

go test -v ./mod/... -coverprofile .test-coverage.txt && rm .test-coverage.txt
