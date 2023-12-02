# admins
Back end del microservicio de administradores.
## Agregar esto a tu bashrc(si no lo tenes, es para ejecutar cosas de go.):
`export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin`


## Link del swagger
`http://localhost:8080/swagger/index.html`

# Link Mongo:
`mongodb+srv://admin:<password>@taller-admins.ez0xrnf.mongodb.net/?retryWrites=true&w=majority`

# tests:
Con el docker-compose levantado:
```bash
docker-compose up
```
y con DB_URI exportado:
```bash
export DB_URI=mongodb://localhost:27017
```
Para correr los tests:
```bash
go test -v ./...
```
Coverage:
```bash
bash test_coverage.sh
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

## golang-jwt:
```bash
go get -u github.com/golang-jwt/jwt/v5
```
## bcrypt:
```bash
go get -u golang.org/x/crypto/bcrypt
```
(linux)
## go-imports:
```bash
sudo apt install golang-golang-x-tools 
```
## golangci-lint:
```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
```
Para que no te moleste con el volumen del docker-compose:
```bash
sudo chmod +r mongodb-data/*
```
## Staticcheck:
```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
```
Se corre con
`saticcheck ./...`
## Uptrace:
```bash
go get github.com/uptrace/uptrace-go
```
## OpenTelemetry:
```bash
go get go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin
```


## Swagger:
Segui los pasos de aca:
https://github.com/swaggo/gin-swagger

Para actualizar la documentacion:
```bash
swag init
```