// Author: yangzq80@gmail.com
// Date: 2022/5/9
//
package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
	// Embed the default VM context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultVMContext
}

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	proxywasm.LogCriticalf("wasm-tcp-protocol===5===NewPluginContext!")
	return &pluginContext{counter: proxywasm.DefineCounterMetric("proxy_wasm_go.connection_counter")}
}

type pluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
	counter proxywasm.MetricCounter
}

// Override types.DefaultPluginContext.
func (ctx *pluginContext) NewTcpContext(contextID uint32) types.TcpContext {
	return &networkContext{counter: ctx.counter}
}

type networkContext struct {
	// Embed the default tcp context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultTcpContext
	counter proxywasm.MetricCounter
}

// Override types.DefaultTcpContext.
func (ctx *networkContext) OnNewConnection() types.Action {
	proxywasm.LogCritical("new connection!")
	return types.ActionContinue
}

// Override types.DefaultTcpContext.
func (ctx *networkContext) OnDownstreamData(dataSize int, endOfStream bool) types.Action {
	proxywasm.LogCriticalf("wasm-tcp-protocol======OnDownstreamData!")
	if dataSize == 0 {
		return types.ActionContinue
	}

	data, err := proxywasm.GetDownstreamData(0, dataSize)
	if err != nil && err != types.ErrorStatusNotFound {
		proxywasm.LogCriticalf("failed to get downstream data: %v", err)
		return types.ActionContinue
	}

	//proxywasm.ReplaceDownstreamData([]byte("test test\n"))
	proxywasm.LogCriticalf(">>>>>> downstream data received >>>>>>\n%s", string(data))
	return types.ActionContinue
}

// Override types.DefaultTcpContext.
func (ctx *networkContext) OnDownstreamClose(types.PeerType) {
	proxywasm.LogInfo("downstream connection close!")
	return
}

// Override types.DefaultTcpContext.
func (ctx *networkContext) OnUpstreamData(dataSize int, endOfStream bool) types.Action {
	proxywasm.LogCriticalf("wasm-tcp-protocol======OnUpstreamData!")

	data, err := proxywasm.GetUpstreamData(0, dataSize)
	if err != nil && err != types.ErrorStatusNotFound {
		proxywasm.LogCritical(err.Error())
	}

	proxywasm.LogCriticalf("<<<<<< upstream data received <<<<<<\n%s", string(data))
	return types.ActionContinue
}

// Override types.DefaultTcpContext.
func (ctx *networkContext) OnStreamDone() {
	ctx.counter.Increment(1)
	proxywasm.LogInfo("connection complete!")
}
