# go-base

Base project for golang

# For develop

This project use air for auto load when has change in project

install it first https://github.com/cosmtrek/air

with go you cant easy install with

```
GO111MODULE=off go get -u github.com/cosmtrek/air
```

**Note**: Make sure that `$GOPATH/bin` was load for excute path

Starting develop

```
make dev
```

Enjoy it :)

# For Production

```
make build
make run
```

**Note**: if you dont need mysql container, please comment out it from `docker-composer.yml`
