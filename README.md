# protologs

Protologs has sample Go code that demonstrates structured and strongly typed logging using [Uber Zap](https://github.com/uber-go/zap) and Protobuf.

This works by using a [zap_marshaler protoc plugin](https://github.com/kazegusuri/go-proto-zap-marshaler) that generates code from Proto messages that implements Zap's [`ObjectMarshaler` interface](https://godoc.org/go.uber.org/zap/zapcore#ObjectMarshaler). Proto generated structs can then be logged with `zap.Object(&someProtoStruct{})`

## Proto definitions

Proto definitions for logs in [/proto](/proto) have base log context fields and service-specific field examples. These files are compiled to the [/pkg/protolog](/pkg/protolog) directory with `protoc`.

```shell script
protoc --zap-marshaler_out=./pkg/protolog --go_out=./pkg/protolog proto/protolog/*.proto
```  

## Examples

`usersvc/main.go` and `paymentsvc/main.go` are examples that both use a common base proto message, but different service-level proto messages.

Run with `go run paymentsvc/main.go` etc. 