# CONTRIBUTING

## Working in Docker

```shell
docker compose run --rm dev sh
```

## Building inside the container

```shell
./bin/container-build
```

`is` is now in your `$PATH`.

## Testing inside the container

```shell
apk add bash
go test ./...
```
