# Go REST API Quick Setup

## 1. Create Project

Run once:

```powershell
go mod init restAPI
```

## 2. Install Swagger

```powershell
go get github.com/swaggo/swag
go get github.com/swaggo/http-swagger
go get github.com/swaggo/files
```

```powershell
go install github.com/swaggo/swag/cmd/swag@latest
```

## 3. Generate Swagger Docs

Run after adding or changing Swagger comments:

```powershell
swag init
```

## 4. Run API

```powershell
go run .
```

## 5. Open

API:

```text
http://localhost:8080/shop
```

Swagger:

```text
http://localhost:8080/docs/index.html
```

## Normal Workflow

If you change Swagger comments:

```powershell
swag init
go run .
```

If you only change Go code:

```powershell
go run .
```
# shop
