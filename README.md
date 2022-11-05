# 个人博客

## 前言

- 基于Golang的Gin库+VVT(Vue3+Vite+TypeScript)开发的个人博客


### Vue前端

#### 安装

```
npm install
```

#### 运行

```
npm run dev
```

#### 打包编译
```
npm run build
```

### Go后端

#### 安装
```
go install
```

#### 运行

```
go run main.go
```

#### 打包编译

```
go build main.go
```

#### 数据库和端口配置
- 目录 admin/configs/application.yml
```yaml
server:
  port: 8080
datasource:
  driverName: mysql
  host: 127.0.0.1
  port: 3306
  database: blog
  username: root
  password: admin
  charset: utf8mb4
  parseTime: true
  loc: Local
admin:
  username: admin
  password: admin
```