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
