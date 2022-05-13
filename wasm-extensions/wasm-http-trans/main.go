// Author: yangzq80@gmail.com
// Date: 2022/5/11
//
package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"strings"
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
	return &pluginContext{}
}

type pluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
	// the remaining token for rate limiting, refreshed periodically.
	remainToken int

	lastRefillNanoSec int64
}

// Override types.DefaultPluginContext.
func (p *pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	proxywasm.LogCriticalf("wasm-http-trans-------------> NewHttpContext")
	return &httpHeaders{contextID: contextID, pluginContext: p}
}

type httpHeaders struct {
	// Embed the default http context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultHttpContext
	contextID     uint32
	pluginContext *pluginContext
}

func (ctx *httpHeaders) OnHttpRequestBody(bodySize int, endOfStream bool) types.Action {
	proxywasm.LogCriticalf("wasm-http-trans--------> OnHttpRequestBody")
	originalBody, err := proxywasm.GetHttpRequestBody(0, bodySize)
	if err != nil {
		proxywasm.LogErrorf("failed to get request body: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("----OnHttpRequestBody-----original request body: %s", string(originalBody))

	return types.ActionContinue
}
func (ctx *httpHeaders) OnHttpResponseBody(bodySize int, endOfStream bool) types.Action {
	proxywasm.LogCriticalf("wasm-http-trans-----------> OnHttpResponseBody")
	originalBody, err := proxywasm.GetHttpResponseBody(0, bodySize)
	if err != nil {
		proxywasm.LogErrorf("failed to get response body: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("----OnHttpResponseBody-----original request body: %s", string(originalBody))

	return types.ActionContinue
}

func (ctx *httpHeaders) OnHttpResponseHeaders(int, bool) types.Action {
	contentType, _ := proxywasm.GetHttpResponseHeader("content-type")
	proxywasm.LogCriticalf("contentType:%v", contentType)
	if strings.Contains(contentType, "application/xml") {
		proxywasm.SendHttpResponse(200, [][2]string{{"Content-Type", "application/json; charset=utf-8"}},
			[]byte("{\"Name\":\"Join------\",\"Content\":\"hello\"}"), -1)
	}
	if strings.Contains(contentType, "application/json") {
		proxywasm.SendHttpResponse(200, [][2]string{{"Content-Type", "application/xml; charset=utf-8"}},
			[]byte("<Message><Name>Join</Name><Content>hello</Content></Message>"), -1)
	}
	return types.ActionContinue
}
