# Pumpkin Plugin API for Go

This package provides everything needed to write a Pumpkin server plugin compiled to WebAssembly using Go.

## Quick start

1. Initialize your Go module:

```bash
go mod init github.com/yourname/yourplugin
go get github.com/Pumpkin-MC/pumpkin-api-go
```

2. Create your plugin (`main.go`):

```go
package main

import (
	"github.com/Pumpkin-MC/pumpkin-api-go/api"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/pumpkin_plugin_context"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/pumpkin_plugin_logging"
	_ "github.com/Pumpkin-MC/pumpkin-api-go/pkg/wit_exports"
)

type MyPlugin struct {
	api.DefaultPlugin
}

func (p *MyPlugin) Metadata() api.Metadata {
	return api.Metadata{
		Name:    "my-go-plugin",
		Version: "0.1.0",
		Authors: []string{"you"},
	}
}

func (p *MyPlugin) OnLoad(ctx *pumpkin_plugin_context.Context) {
    pumpkin_plugin_logging.Log(pumpkin_plugin_logging.LevelInfo(), "Go plugin loaded!")
}

func init() {
	api.RegisterPlugin(&MyPlugin{})
}

func main() {}
```

3. Build your plugin into a WebAssembly component using TinyGo:

```bash
tinygo build -o my_plugin.wasm -target=wasi main.go
```

Note: You might need to use `wit-component` to turn the core WASM into a component if your build tool doesn't do it automatically.
