# Web Server for fiber

### fiber install
```
get -u github.com/gofiber/fiber
```
-----
### fiber github
link : https://github.com/gofiber/fiber
### fiber docs
link : https://docs.gofiber.io
-----
### fiber start !
/helloworld
server start
```
cd helloworld
go run main.go
```
-----
### fiber + entgo + godotenv => login
mysql use

server start
```
cd login-web
sh docker/setDB.sh
go run main.go
```

test
```
// Welcome
curl -u ningning:021023 localhost:3000
curl -u giselle:001030 localhost:3000
curl -u karina:000411 localhost:3000
curl -u winter:010101 localhost:3000

// Unauthorized
curl -u foo:bar localhost:3000
```
-----