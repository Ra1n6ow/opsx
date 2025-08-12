# opsx

devops

# air

```shell
air -c .air.toml
```

# LICENSE

根目录 LICENSE 生成

```shell
go install github.com/nishanths/license/v5@latest
$ license -list # 查看支持的代码协议
# 在根目录下执行
$ license -n 'JingFeng Du <jeffduuu@gmail.com>' -o LICENSE mit
```

添加头文件 LICENSE 信息

```shell
go install github.com/marmotedu/addlicense@latest

addlicense -v -f ./scripts/boilerplate.txt --skip-dirs=third_party,_output .
```

# GRPC

## protoc-go-inject-tag

- 在 grpc 生成的代码中，json tag 会自动添加 omitempty，可以通过在 proto 文件添加注释 `// @inject_tag: json:"status"` 的方式，只生成注入的 tag

```shell
go install github.com/favadi/protoc-go-inject-tag@latest

# Makefile 相关
@echo "===========> Inject custom tags"
@protoc-go-inject-tag -input="$(APIROOT)/core/v1/*.pb.go"

# jnject-tag 前
go run examples/client/health/main.go
{"timestamp":"2025-08-11 11:00:47"}
# jnject-tag 后
go run examples/client/health/main.go
{"status":0,"timestamp":"2025-08-11 11:00:47"}
```

> 在上面的方法中，grpc-gateway 仍然不会显示该字段
> 可以将 status 字段改为 optional
> 当为指针时（即使指向零值），字段也会显示在 JSON 中. 当指针为 nil 时，字段不会显示在 JSON 中.

## grpc-gateway

```shell
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.24.0
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.24.0
```

在 RESTful API 中，DELETE 方法通常用于“删除资源”，按 REST 规范应只携带资源标识符（如路径或查询参数），不建议附带请求体。然而，在某些场景下需要通过 DELETE 请求体传递额外信息（如删除选项或多个待删除的资源列表）。此时，可设置 allow_delete_body=true，放宽对 DELETE 请求的限制，允许携带请求体，

```shell
# Makefile
		--grpc-gateway_out=allow_delete_body=true,paths=source_relative:$(APIROOT) \
		--openapiv2_out=$(PROJ_ROOT_DIR)/api/openapi \
		--openapiv2_opt=allow_delete_body=true,logtostderr=true \

# 支持以下接口调用
curl -XDELETE -H"Content-Type: application/json" -d'{"postIDs":["post-w6irkg","post-w6irkb"]}'
```
