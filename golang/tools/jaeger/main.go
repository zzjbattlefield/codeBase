package main

//链路追踪
import (
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {
	//先配置jaeger配置
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "192.168.58.130:6831",
		},
		ServiceName: "goShop",
	}
	//新建一条链路 让log打印的终端
	trancer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	//新建一个span(此时开始计算开始时间)
	spanMain := trancer.StartSpan("main")
	time.Sleep(100 * time.Millisecond)
	//必须要Finish才会计算结束时间(在上层的链路一定要最后finsh)
	defer spanMain.Finish()
	//添加opentracing.ChildOf表示span1是spanMain的子调用
	span1 := trancer.StartSpan("go-grpc-web", opentracing.ChildOf(spanMain.Context()))
	time.Sleep(50 * time.Millisecond)
	defer span1.Finish()
}
