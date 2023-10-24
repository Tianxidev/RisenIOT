# RisenIOT
物联网云平台


<div align="center">
<img src="https://img.shields.io/badge/RisenIOT-Beta-brightgreen?style=flat-square&logo=github" />
<img src="https://img.shields.io/badge/license-Mit-yellowgreen?style=flat-square&logo=github" />
</div>

## 注意事项

1. 开发环境在 Windows 下，使用 GoLand 开发，go 版本需要 1.21.0 或以上, 如您在 Linux 下开发，需要自行修正编译脚本;

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
go env -w CGO_ENABLED=0
go env -w GOOS=windows
go env -w GOARCH=amd64
go run ./
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
go env -w CGO_ENABLED=0
go env -w GOOS=windows
go env -w GOARCH=amd64
go build -o RisenIOT.exe ./

// 交叉编译 Linux 平台
go mod tidy
go env -w CGO_ENABLED=0
go env -w GOOS=linux
go env -w GOARCH=amd64
go build -o RisenIOT ./
```

## 功能支持

- [ ] HTTP API 接入协议支持
- [X] 标准 MQTT 协议支持

## 设备协议支持

- [X] 云之声灯控协议


## EMQX WebHook 配置

接口: /api/v1/device/data/receive/emqx/webhook

```sql
SELECT
  timestamp,
  topic,
  qos,
  event,
  publish_received_at,
  peerhost,
  clientid,
  payload,
  bin2hexstr(payload) as payload_hexstr
FROM
  "#"
```

