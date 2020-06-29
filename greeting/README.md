## Building greeting api

#### Compile statically to include all the native libraries needed
```
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o greeting
```

#### Build the docker image
```
docker build -t name/greeting:go .
```
