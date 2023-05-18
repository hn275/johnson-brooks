# Johnson & Brooks

## Web `./web`

[TailwindCSS](https://tailwindcss.com/docs/installation)

[Astro](https://docs.astro.build/en/getting-started/)

[React](https://react.dev/reference/react)

```sh
yarn # install all dependencies
yarn dev # dev server at localhost:3000
```

## API `./api`

[Chi](https://github.com/go-chi/chi) - V5

[Golang Mongo Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.11.6/mongo)

### Dev server at `:8080`

```sh
docker compose up
```

### Accessing db (the Mongo image)

```sh
# with Docker shell
docker exec -it johnson-brooks-db mongosh --username root --password root

# with Mongo Compass with the connection string
# mongodb://root:root@localhost:27017/
```

### Scripts

```sh
# any of these commands will start the docker instance of mongo image

# running tests
sh run.sh test ./<dir name to test>

# generate mock data
sh run.sh mock
```
