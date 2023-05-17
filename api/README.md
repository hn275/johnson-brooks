# API

### Seed data

```sh
export MONGODB_URL="mongodb://root:root@localhost:27017/"
go run ./scripts/seeds/main.go
```

### Accessing db

Either with Docker

```sh
docker exec -it johnson-brooks-db mongosh --username root --password root
```

Or with Mongo Compass with the connection string

```
mongodb://root:root@localhost:27017/
```

### Scripts

```sh
# any of these commands will start the docker instance of mongo image

# running tests
sh run.sh test ./<dir name to test>

# mock data
sh run.sh mock
```
