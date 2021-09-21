# Envoy proxy as an API gateway. Built with Bazel. Run in Kubernetes.

## Usage

```shell
% grpcurl -plaintext \
  -d '{"name": "Bazel"}' \
  127.0.0.1:55000 svc.ServiceOne.Hello

{
  "body": "Hello, Bazel"
}
```

```shell
% curl -i http://localhost:8080/v1/hello\?name\=Bazel

HTTP/1.1 200 OK
content-type: application/json
x-envoy-upstream-service-time: 2
grpc-status: 0
grpc-message:
content-length: 28
date: Tue, 21 Sep 2021 23:33:27 GMT
server: envoy

{
 "body": "Hello, Bazel"
}
```

NOTE: Until `service-one` bacome available, Envoy will return HTTP 503: `upstream connect error` or `no healthy upstream`
