package waiter

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type (
	WaitFunc func(ctx context.Context) error

	Waiter interface {
		Add(fns ...WaitFunc)
		Wait() error
		Context() context.Context
		CancelFunc() context.CancelFunc
	}

	waiter struct {
		ctx    context.Context
		fns    []WaitFunc
		cancel context.CancelFunc
	}
	waiterCfg struct {
		parentCtx context.Context
	}
)

func New(opts ...func(ctx context.Context)) Waiter {
	parentCtx := context.Background()
	for _, opt := range opts {
		opt(parentCtx)
	}

	w := &waiter{
		fns: []WaitFunc{},
	}

	w.ctx, w.cancel = context.WithCancel(parentCtx)

	return w
}

func (w *waiter) Add(fns ...WaitFunc) {
	w.fns = append(w.fns, fns...)
}

func (w waiter) Wait() error {
	g, ctx := errgroup.WithContext(w.ctx)

	g.Go(func() error {
		<-ctx.Done()
		w.cancel()
		return nil
	})

	for _, fn := range w.fns {
		fn := fn
		g.Go(func() error { return fn(ctx) })
	}

	return g.Wait()
}

func (w waiter) Context() context.Context {
	return w.ctx
}

func (w waiter) CancelFunc() context.CancelFunc {
	return w.cancel
}
