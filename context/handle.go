package context

import "context"

func SetValue(parent context.Context, key any, val any) context.Context {
	return context.WithValue(parent, key, val)
}

func GetValue(ctx context.Context, key any) any {
	return ctx.Value(key)
}
