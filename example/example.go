//go:generate go run github.com/Armatorix/ddwrap -p example example.go internal_gen.go
//go:generate go run github.com/Armatorix/ddwrap -p externalexample example.go externalexample/external_gen.go
package example

import (
	"context"
	"io"
)

type Example interface {
	Hello() string
	HelloListArg(args []string, argsList ...string) string
	HelloListRet() []string
	HelloFnRetAndIn(inFn func() func() string) func() string
}

type Example2 interface {
	AsyncHello(ctx context.Context) (string, error)
}

type PublicStruct struct {
}

type privateStruct struct {
}

type Example3 interface {
	ExternalArgsHello(ctx context.Context, reader io.Reader) (string, error)
	InternalPublicStructHello(ctx context.Context, reader PublicStruct) (string, error)
}

type ShouldFailExternallyExample interface {
	ShouldFail(ex privateStruct) (string, error)
}
