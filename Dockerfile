FROM --platform=$BUILDPLATFORM golang:1.21-alpine AS builder

# 安装必要的构建工具
RUN apk add --no-cache gcc musl-dev nodejs npm

# 设置工作目录
WORKDIR /app

# 复制项目文件
COPY . .

# 安装 Wails
RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest && \
    wails doctor

# 构建应用
RUN wails build -platform $TARGETOS/$TARGETARCH

# 使用 scratch 作为最终镜像
FROM scratch

# 复制构建产物
COPY --from=builder /app/build/bin/* /bin/ 