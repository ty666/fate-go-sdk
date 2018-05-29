package fate_go_sdk

import "context"

type fateKey struct{}

func NewContext(ctx context.Context, f *Fate) context.Context {
	return context.WithValue(ctx, fateKey{}, f)
}

func FromContext(ctx context.Context) *Fate {
	return ctx.Value(fateKey{}).(*Fate)
}
