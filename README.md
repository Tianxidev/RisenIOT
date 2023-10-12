# RisenIOT
物联网云平台


<div align="center">
<img src="https://img.shields.io/badge/RisenIOT-Beta-brightgreen?style=flat-square&logo=github" />
<img src="https://img.shields.io/badge/license-Mit-yellowgreen?style=flat-square&logo=github" />
<img src="https://img.shields.io/github/languages/top/Ritusan/color-library?color=blue" />
</div>

## 注意事项

1. 开发环境在 Windows 下，使用 GoLand 开发，go 版本需要 1.21.0 或以上, 如您在 Linux 下开发，需要自行修正编译脚本;
2. 本项目启用了 CGO，需要安装 gcc 环境，如您在 Windows 下开发，需要自行安装 gcc 环境;

## 运行

前端
```bash
cd frontend
npm i
npm run dev
```

后端
```bash
// Windows 平台
go mod tidy
go mod vendor
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64
SET GO111MODULE=on
go run ./cmd/app
```

## 编译


前端
```bash
cd frontend
npm i
npm run build
```

后端
```bash
// Windows 平台
go mod tidy
go mod vendor
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64
SET GO111MODULE=on
go build -o RisenIOT.exe ./cmd/app

// 交叉编译 Linux 平台
go mod tidy
go mod vendor
SET CGO_ENABLED=1 
SET GOOS=linux
SET GOARCH=amd64
SET GO111MODULE=on
go build -o RisenIOT ./cmd/app
```

## 功能支持

- [ ] HTTP API 接入协议支持
- [ ] 标准 MQTT 协议支持


