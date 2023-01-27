## Commands

### Create project
```
go mod init examples/microservices
```

### Download mysql
```
go get github.com/go-sql-driver/mysql
```

### Download Gin Framework
```
go get github.com/gin-gonic/gin
```

### Download GORM
```
go get github.com/jinzhu/gorm
```

### Tests
- Install httpie [Link](https://github.com/httpie/httpie)

#### Install on MacOS
```
brew install httpie
```

#### Install on Linux
```
curl -SsL https://packages.httpie.io/deb/KEY.gpg | apt-key add -
curl -SsL -o /etc/apt/sources.list.d/httpie.list https://packages.httpie.io/deb/httpie.list
apt update
apt install httpie
```

#### Test Commands
```
http POST localhost:9090/v1/todo description=deneme1 title=deneme1
http POST localhost:9090/v1/todo description=deneme2 title=deneme2
http POST localhost:9090/v1/todo description=deneme3 title=deneme3
http POST localhost:9090/v1/todo description=deneme4 title=deneme4
http GET localhost:9090/v1/todo
http PUT localhost:9090/v1/todo/1 description=deneme5 title=deneme5
http GET localhost:9090/v1/todo
http DELETE localhost:9090/v1/todo/1
http GET localhost:9090/v1/todo
```

### OpenAPI (Swagger)
```
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/swag/cmd/swag
go get github.com/swaggo/files
go get github.com/swaggo/gin-swagger
swag init
```
