# Johnson & Brooks

## Web

[TailwindCSS](https://tailwindcss.com/docs/installation)

[Astro](https://docs.astro.build/en/getting-started/)

React (I trust you know how to use this already)

```sh
yarn # install all dependencies
yarn dev # dev server at localhost:3000
```

## API

[Chi](https://github.com/go-chi/chi) - V5

[Golang Mongo Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.11.6/mongo)

```sh
# to build and run the container, remove the flag `--build` other wise
docker compose up --build

# accessing mongosh in the container
docker exec -it johnson-brooks-db mongosh --username root --pasword root
```
