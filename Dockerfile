FROM --platform=$BUILDPLATFORM wailsapp/wails:latest AS builder

# 设置工作目录
WORKDIR /app

# 复制项目文件
COPY . .

# 构建应用
RUN if [ "$TARGETOS" = "windows" ]; then \
        wails build -platform windows/amd64; \
    elif [ "$TARGETOS" = "darwin" ]; then \
        wails build -platform darwin/amd64; \
    else \
        wails build -platform linux/amd64; \
    fi

# 使用 scratch 作为最终镜像
FROM scratch

# 复制构建产物
COPY --from=builder /app/build/bin/* /bin/ 