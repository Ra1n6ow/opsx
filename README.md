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