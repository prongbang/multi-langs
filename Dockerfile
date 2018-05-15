FROM golang:1.9.3-alpine3.7

WORKDIR /go/src/multi-langs
COPY . .

# RUN apt-get install git

# RUN go get -d -v github.com/kardianos/govendor
# RUN govendor -d -v github.com/dgrijalva/jwt-go
# RUN govendor -d -v github.com/jinzhu/gorm
# RUN govendor -d -v github.com/labstack/echo
# RUN govendor -d -v github.com/labstack/echo-contrib/session
# RUN govendor -d -v github.com/labstack/echo/middleware
# RUN govendor -d -v golang.org/x/net/websocket

# RUN go install main.go
# RUN go build -o main . 

# CMD ["go", "run", "main.go"]

# FROM golang:onbuild