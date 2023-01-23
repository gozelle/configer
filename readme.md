# Configer

简单的配置文件加载器。

## 特性

* 仅支持 toml 文件格式；
* 仅支持使用指针和切片类型接受配置项；
* 可自动校验配置项是否为空；
* 可绑定默认值。

## 用法

Configer 只校验指针类型的配置项；如果配置项为指针类型，并且没有定义 default 值，则认为该项是必填配置项。

## 示例

结构体:

```go
type Config struct {
	Address  *string   `toml:"address"`                  // 必填项  （指针类型，未定义 default 标签）
	LogLevel *string   `toml:"log_level" default:"warn"` // 非必填项 (指针类型，定义了 default 标签)
	Database *Database `toml:"database"`                 // 必填项  （struct 类型，不支持 default，且其 struct 中未定义 default 标签）
	Redis    *Redis    `toml:"redis"`                    // 非必填项（struct 中定义了 default 标签，将会自动填充值）
	Nodes   []*Node    `toml:"nodes"`                    // 必填项  （切片型为必填项, struct 中的 default 只对子项默认值填充）
}

type Database struct {
	DSN  *string  `toml:"dsn"`
}

type Redis struct {
	Host *string  `toml:"host" default:"127.0.0.1:5736"`
}

type Node struct {
	Name *string `toml:"name" default:"node"`
}
```

## 支持配置类型

* int*、uint*、float*
* struct
* slice

## 使用

```go
conf := &Config{}
err = configer.Load(data,conf)

err = configer.LoadFile(file,conf)
```