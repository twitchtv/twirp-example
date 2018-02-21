package hooks

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/twitchtv/twirp"
)

var ctxKey = new(int)

func LoggingHooks(w io.Writer) *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			startTime := time.Now()
			ctx = context.WithValue(ctx, ctxKey, startTime)
			return ctx, nil
		},
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			svc, _ := twirp.ServiceName(ctx)
			meth, _ := twirp.MethodName(ctx)
			fmt.Fprintf(w, "received req svc=%q method=%q\n", svc, meth)
			return ctx, nil
		},
		ResponseSent: func(ctx context.Context) {
			startTime := ctx.Value(ctxKey).(time.Time)
			svc, _ := twirp.ServiceName(ctx)
			meth, _ := twirp.MethodName(ctx)
			fmt.Fprintf(w, "response sent svc=%q method=%q time=%q\n", svc, meth, time.Since(startTime))
		},
	}
}
