// DO NOT EDIT: auto generated from ddwrap

package externalexample

import (
	"context"
	"io"

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
		tracer.ServiceName(w.serviceName),
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
	out0 := w.inner.HelloListArg(in0, in1...)
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
		tracer.ServiceName(w.serviceName),
		tracer.ResourceName(resourceName),
	)
}

func (w *ddExample2Wrapper) AsyncHello(in0 context.Context) (string, error) {
	span, _ := w.startSpan(context.Background(), "AsyncHello")
	defer span.Finish()
	out0, out1 := w.inner.AsyncHello(in0)
	return out0, out1
}

type ddExample3Wrapper struct {
	serviceName   string
	operationName string
	inner         Example3
}

func NewDDExample3Wrapper(inner Example3, serviceName, operationName string) *ddExample3Wrapper {
	return &ddExample3Wrapper{
		serviceName:   serviceName,
		operationName: operationName,
		inner:         inner,
	}
}
func (w *ddExample3Wrapper) startSpan(ctx context.Context, resourceName string) (tracer.Span, context.Context) {
	return tracer.StartSpanFromContext(
		ctx,
		w.operationName,
		tracer.ServiceName(w.serviceName),
		tracer.ResourceName(resourceName),
	)
}

func (w *ddExample3Wrapper) ExternalArgsHello(in0 context.Context, in1 io.Reader) (string, error) {
	span, _ := w.startSpan(context.Background(), "ExternalArgsHello")
	defer span.Finish()
	out0, out1 := w.inner.ExternalArgsHello(in0, in1)
	return out0, out1
}

func (w *ddExample3Wrapper) InternalPublicStructHello(in0 context.Context, in1 PublicStruct) (string, error) {
	span, _ := w.startSpan(context.Background(), "InternalPublicStructHello")
	defer span.Finish()
	out0, out1 := w.inner.InternalPublicStructHello(in0, in1)
	return out0, out1
}

type ddShouldFailExternallyExampleWrapper struct {
	serviceName   string
	operationName string
	inner         ShouldFailExternallyExample
}

func NewDDShouldFailExternallyExampleWrapper(inner ShouldFailExternallyExample, serviceName, operationName string) *ddShouldFailExternallyExampleWrapper {
	return &ddShouldFailExternallyExampleWrapper{
		serviceName:   serviceName,
		operationName: operationName,
		inner:         inner,
	}
}
func (w *ddShouldFailExternallyExampleWrapper) startSpan(ctx context.Context, resourceName string) (tracer.Span, context.Context) {
	return tracer.StartSpanFromContext(
		ctx,
		w.operationName,
		tracer.ServiceName(w.serviceName),
		tracer.ResourceName(resourceName),
	)
}

func (w *ddShouldFailExternallyExampleWrapper) ShouldFail(in0 privateStruct) (string, error) {
	span, _ := w.startSpan(context.Background(), "ShouldFail")
	defer span.Finish()
	out0, out1 := w.inner.ShouldFail(in0)
	return out0, out1
}
