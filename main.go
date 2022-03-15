package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"log"
)

func main() {
	log.Println("hello tinyGo")
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct{}

func (*vmContext) OnVMStart(vmConfigurationSize int) types.OnVMStartStatus {
	data, err := proxywasm.GetVMConfiguration()
	if err != nil {
		proxywasm.LogCriticalf("error reading vm configuration: %v", err)
	}

	proxywasm.LogInfof("vm config: %s", string(data))
	return types.OnVMStartStatusOK
}

// Implement types.VMContext.
func (*vmContext) NewPluginContext(uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
}

// Override types.DefaultPluginContext.
func (ctx pluginContext) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	data, err := proxywasm.GetPluginConfiguration()
	if err != nil {
		proxywasm.LogCriticalf("error reading plugin configuration: %v", err)
	}

	proxywasm.LogInfof("plugin config: %s", string(data))
	return types.OnPluginStartStatusOK
}
