# 当前目录
CUR_DIR=./
OUT_DIR=$(CUR_DIR)/bin

# 命令
GO_BUILD = go build -trimpath

.PHONY: build
build:
	$(GO_BUILD) -o $(OUT_DIR)/ .
