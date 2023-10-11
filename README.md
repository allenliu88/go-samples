# Go HTTP Server for Static Files

## Local Run

```shell
go run main.go -port=8080

INFO[0000] port: 8080                                   
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.index (3 handlers)
[GIN-debug] GET    /albums                   --> main.getAlbums (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
```

## Build & Run

For Linux

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/go-samples-linux-amd64 main.go

chmod +x go-samples-linux-amd64
nohup ./go-samples-linux-amd64 -port=8080 > http.log 2>&1 &
```

For Windows

```shell
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/go-samples-windows-amd64.exe main.go

chmod +x go-samples-windows-amd64.exe
nohup ./go-samples-windows-amd64.exe -port=8080 > http.log 2>&1 &
```

## Docker

```shell
docker buildx build --platform linux/amd64 -t allen88/go-samples:1.0.0 -f Dockerfile .
docker push allen88/go-samples:1.0.0
```

## Kubernetes

```shell
kubectl apply -f ./Deployment.yaml
```

## License

Copyright &copy; 2023 Allen

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
