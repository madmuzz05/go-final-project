# Run
Run your API
```sh
go run api/main.go
```
To test API open Swagger
```sh
http://localhost:8081/api/v1/swagger/index.html#/
```

Enjoy

# Setup
copy and set up your .env based on the .env.example 
```bash
cp internal/config/.env.example internal/config/.env
```

## Install
Validation and download dependency

```bash
go mod tidy
```
## Swagger
### Start using it
1. Add comments to your API source code, [See Declarative Comments Format](https://github.com/swaggo/swag#declarative-comments-format).
2. Download [Swag](https://github.com/swaggo/swag) for Go by using:
```sh
go install github.com/swaggo/swag/cmd/swag@latest
```
To build from source you need [Go](https://golang.org/dl/) (1.17 or newer).

Or download a pre-compiled binary from the [release page](https://github.com/swaggo/swag/releases).
3. export path
```sh
export PATH=$PATH:$(go env GOPATH)/bin
```

4. Run the [Swag](https://github.com/swaggo/swag) in your Go project root folder which contains `main.go` file, [Swag](https://github.com/swaggo/swag) will parse comments and generate required files(`docs` folder and `docs/doc.go`).
```sh
swag init -g routes/router.go -o service/docs
```
