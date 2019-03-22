# Build go app for linux (from mac)

```sh
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o web ./main.go
```

# Build/push container

```sh
docker build -t jweissig/istio-demo:v0.1 .
docker push jweissig/istio-demo:v0.1
```
