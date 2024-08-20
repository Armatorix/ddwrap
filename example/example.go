package example

import "context"

type Example interface {
	Hello() string
	HelloListArg(args []string, argsList ...string) string
	HelloListRet() []string
	HelloFnRetAndIn(inFn func() func() string) func() string
}

type Example2 interface {
	AsyncHello(ctx context.Context) (string, error)
}
