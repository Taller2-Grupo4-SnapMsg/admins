# admins
Back end del microservicio de administradores.
## Agregar esto a tu bashrc:
`export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin`


## Link del swagger
`http://localhost:8080/swagger/index.html`

# Link Mongo:
`mongodb+srv://admin:<password>@taller-admins.ez0xrnf.mongodb.net/?retryWrites=true&w=majority`

# tests:
```bash
go test -v ./...
```

# format:
Para ver errores de formateo:
```bash
go fmt ./...
```

Para que el format te formatee todo automaticamente:
```bash
go fmt ./... -w
```

# pre-commit:
Si no tenes pre-commit:
```bash
pip install pre-commit
```

Primero lo instalamos en el repo:
```bash
pre-commit install
```
Para ejecutarlo sin commitear:
```bash
pre-commit run --all-files
```

# Instalaciones:
(linux)
## go-imports:
```bash
sudo apt install golang-golang-x-tools 
```
## golangci-lint:
```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
```
## Staticcheck:
```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
```
Se corre con
`saticcheck ./...`

## Swagger:
Segui los pasos de aca:
https://github.com/swaggo/gin-swagger

Para actualizar la documentacion:
```bash
swag init
```