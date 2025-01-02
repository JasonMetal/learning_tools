// Go-Kit 是一个用于构建微服务的 Go 语言库和工具集。以下是它的主要特点：
//
// - **模块化设计**：Go-Kit 提供了多个模块，如服务、端点（Endpoint）、传输层（Transport）等，这些模块可以独立使用也可以组合起来构建复杂的服务。
// - **中间件支持**：可以在不同层次（服务层、端点层、传输层）添加中间件来处理横切关注点，例如日志记录、认证授权等。
// - **服务发现与负载均衡**：虽然 Go-Kit 本身不直接提供这些功能，但它很容易与其他相关工具集成以实现这些特性。
// - **监控和追踪**：提供了对分布式系统的监控和追踪的支持，便于调试和性能优化。
//
// 从你提供的代码片段来看：
// - `v1_service.NewService()` 创建了一个新的服务实例。
// - `v1_endpoint.NewEndPointServer(server)` 将服务包装成端点。
// - `v1_transport.NewHttpHandler(endpoints)` 把端点转换为 HTTP 处理程序，从而可以通过 HTTP 协议访问该服务。
//
// 这段代码展示了如何使用 Go-Kit 的核心组件来搭建一个简单的 HTTP 服务器。
package main

import (
	"fmt"
	"github.com/hwholiday/learning_tools/go-kit/v1/v1_endpoint"
	"github.com/hwholiday/learning_tools/go-kit/v1/v1_service"
	"github.com/hwholiday/learning_tools/go-kit/v1/v1_transport"
	"net/http"
)

func main() {
	server := v1_service.NewService()
	endpoints := v1_endpoint.NewEndPointServer(server)
	httpHandler := v1_transport.NewHttpHandler(endpoints)
	fmt.Println("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)
}
