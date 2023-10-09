# gin-cache

[![Release](https://img.shields.io/github/release/zychimne/gin-cache.svg?style=flat-square)](https://github.com/zychimne/gin-cache/releases)
[![doc](https://img.shields.io/badge/go.dev-doc-007d9c?style=flat-square&logo=read-the-docs)](https://pkg.go.dev/github.com/zychimne/gin-cache)
[![goreportcard for gin-cache](https://goreportcard.com/badge/github.com/zychimne/gin-cache)](https://goreportcard.com/report/github.com/zychimne/gin-cache)
![](https://img.shields.io/badge/license-MIT-green)
[![codecov](https://codecov.io/gh/zychimne/gin-cache/branch/main/graph/badge.svg?token=MX8Z4D5RZS)](https://codecov.io/gh/zychimne/gin-cache)

[English](README_ZH.md) | 🇨🇳 中文

一个用于缓存 http 接口内容的 gin 高性能中间件。相比于官方的 gin-contrib/cache，gin-cache 有巨大的性能提升。

# 特性

- 相比于 gin-contrib/cache，性能提升巨大。
- 同时支持本机内存和 redis 作为缓存后端。
- 支持用户根据请求来指定 cache 策略。
- 使用 singleflight 解决了缓存击穿问题。
- 仅缓存 http 状态码为 2xx 的回包

# 用法

## 安装

```
go get -u github.com/zychimne/gin-cache
```

## 例子

## 使用本地缓存

```go
package main

import (
	"time"

	"github.com/zychimne/gin-cache"
	"github.com/zychimne/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	memoryStore := persist.NewMemoryStore(1 * time.Minute)

	app.GET("/hello",
		cache.CacheByRequestURI(memoryStore, 2*time.Second),
		func(c *gin.Context) {
			c.String(200, "hello world")
		},
	)

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
```

### 使用 redis 作为缓存

```go
package main

import (
	"time"

	"github.com/zychimne/gin-cache"
	"github.com/zychimne/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	app := gin.New()

	redisStore := persist.NewRedisStore(redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	}))

	app.GET("/hello",
		cache.CacheByRequestURI(redisStore, 2*time.Second),
		func(c *gin.Context) {
			c.String(200, "hello world")
		},
	)
	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
```

# 压测

```
wrk -c 500 -d 1m -t 5 http://127.0.0.1:8080/hello
```

## MemoryStore

![MemoryStore QPS](https://www.cyhone.com/img/gin-cache/memory_cache_qps.png)

## RedisStore

![RedisStore QPS](https://www.cyhone.com/img/gin-cache/redis_cache_qps.png)
