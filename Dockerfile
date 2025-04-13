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
    libappindicator3-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

# 设置工作目录
WORKDIR /app

# 复制项目文件
COPY . .

# 安装 Wails
RUN go install github.com/wailsapp/wails/v2/cmd/wails@v2.8.0

# 构建应用
RUN mkdir -p build/bin && \
    if [ "$TARGETOS" = "windows" ]; then \
        echo "Building for Windows" && \
        CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o build/bin/app.exe ./main.go; \
    elif [ "$TARGETOS" = "darwin" ]; then \
        echo "Building for macOS" && \
        CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o build/bin/app ./main.go; \
    else \
        echo "Building for Linux" && \
        CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o build/bin/app ./main.go; \
    fi

# 使用 scratch 作为最终镜像
FROM scratch

# 复制构建产物
COPY --from=builder /app/build/bin/* /bin/ 