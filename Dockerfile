FROM --platform=$BUILDPLATFORM golang:1.22-bullseye AS builder

# 安装必要的构建工具
RUN apt-get update && apt-get install -y \
    gcc \
    g++ \
    make \
    nodejs \
    npm \
    libgtk-3-dev \
    libwebkit2gtk-4.0-dev \
    libayatana-appindicator3-dev \
    && rm -rf /var/lib/apt/lists/*

# 设置工作目录
WORKDIR /app

# 复制项目文件
COPY . .

# 安装 Wails
RUN go version && \
    node --version && \
    npm --version && \
    go install github.com/wailsapp/wails/v2/cmd/wails@latest

# 验证 Wails 安装
RUN which wails && \
    wails version

# 构建应用
RUN wails build -platform $TARGETOS/$TARGETARCH

# 使用 scratch 作为最终镜像
FROM scratch

# 复制构建产物
COPY --from=builder /app/build/bin/* /bin/ 