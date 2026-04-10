# Go Framework Playground (Go Zero)

This branch contains a demo project using [go-zero](https://go-zero.dev/).

The key to maintain multiple microservices, whether separate, or aggregated into a monolithic service using [`ServiceGroup`](https://pkg.go.dev/github.com/zeromicro/go-zero@v1.10.1/core/service#ServiceGroup), is to generate multiple zrpc project in the current directory, and move the entrypoint into `cmd/`.

Please check the [Taskfile](./Taskfile.yml) for more detail.