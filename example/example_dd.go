package test_package_name

import (
	"context"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type ddExampleWrapper struct {
	serviceName   string
	operationName string
	inner         Example
}

func NewDDExampleWrapper(inner Example, serviceName, operationName string) *ddExampleWrapper {
	return &ddExampleWrapper{
		serviceName:   serviceName,
		operationName: operationName,
		inner:         inner,
	}
}
func (w *ddExampleWrapper) startSpan(ctx context.Context, resourceName string) (tracer.Span, context.Context) {
	return tracer.StartSpanFromContext(
		ctx,
		w.operationName,
		tracer.ServiceName(openaiServiceName),
		tracer.ResourceName(resourceName),
	)
}

func (w *ddExampleWrapper) Hello() string {
	span, _ := w.startSpan(context.Background(), "Hello")
	defer span.Finish()
	out0 := w.inner.Hello()
	return out0
}

func (w *ddExampleWrapper) HelloListArg(in0 []string, in1 ...string) string {
	span, _ := w.startSpan(context.Background(), "HelloListArg")
	defer span.Finish()
	out0 := w.inner.HelloListArg(in0, in1)
	return out0
}

func (w *ddExampleWrapper) HelloListRet() []string {
	span, _ := w.startSpan(context.Background(), "HelloListRet")
	defer span.Finish()
	out0 := w.inner.HelloListRet()
	return out0
}

func (w *ddExampleWrapper) HelloFnRetAndIn(in0 func() func() string) func() string {
	span, _ := w.startSpan(context.Background(), "HelloFnRetAndIn")
	defer span.Finish()
	out0 := w.inner.HelloFnRetAndIn(in0)
	return out0
}

type ddExample2Wrapper struct {
	serviceName   string
	operationName string
	inner         Example2
}

func NewDDExample2Wrapper(inner Example2, serviceName, operationName string) *ddExample2Wrapper {
	return &ddExample2Wrapper{
		serviceName:   serviceName,
		operationName: operationName,
		inner:         inner,
	}
}
func (w *ddExample2Wrapper) startSpan(ctx context.Context, resourceName string) (tracer.Span, context.Context) {
	return tracer.StartSpanFromContext(
		ctx,
		w.operationName,
		tracer.ServiceName(openaiServiceName),
		tracer.ResourceName(resourceName),
	)
}

func (w *ddExample2Wrapper) AsyncHello(in0 context.Context) (string, error) {
	span, _ := w.startSpan(context.Background(), "AsyncHello")
	defer span.Finish()
	out0, out1 := w.inner.AsyncHello(in0)
	return out0, out1
}
