# Datumbazo

Database engine written in Go as a learning exercise. 'Datumbazo' means 'database' in Esperanto.

[Database]("https://icons8.com/icon/Sp5Wxo7V9G0k/database") icon by [Icons8]("https://icons8.com")

## Build environment

- go (Ubuntu: `golang-go`)
- gofmt (Ubuntu: `golang-go`)
- task (Ubuntu: `task` - Snap)

## Building

`task build`, or:

```sh
gofmt -w .
go mod tidy
go mod verify
go test -cover
go build
```

## Environment file

```
DBZO_HOST=localhost
DBZO_PORT=8874
DBZO_USER=user
DBZO_PASSWORD=password
```