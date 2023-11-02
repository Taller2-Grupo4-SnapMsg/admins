# admins
Back end del microservicio de administradores.

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
sudo snap install golangci-lint   
```