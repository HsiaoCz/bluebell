.PHONY:all build run gotool clean help
# 这些命令都是我们需要执行的命令

# 定义的变量
BINARY="bluebell"

# make 后面什么都不跟的时候执行的
all: gotool build

#  make build 执行的命令
build:
     CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

# make run 执行的命令
# 加@可以不打印
run:
    @go run ./main.go conf/config.yaml

# 格式化 和检查
gotool:
    go fmt ./
	go vet ./

# 删除当前项目下的文件 不过在删除之前先判断
clean:
    @if [-f ${BINARY} ] ;then rm ${BINARY} ;fi

# help 输出一些命令提示
help:
    @echo "make - 格式化 Go 代码,并编译生成二进制文件"
	@echo "make build - 编译Go 代码,生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件 和vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and  'vet'