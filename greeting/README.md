## greeting api using go-gin-gonic and opentracing

#### Environment variables
- PORT : specify what port to listen on
- MESSAGE : the string that will be returned along with the hostname
- REMOTE : URL of the remote service that will be invoked, expecting a plain text response - e.g. http://localhost:8080/greetings
- COLLECTOR : address of the jaeger collector - defaults to localhost:6831
- GIN_MODE : set to 'release' to disable extra log entries

#### Compile statically to include all the native libraries needed
```
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o greeting
```

#### Build the docker image
```
docker build -t name/greeting:go .
```


