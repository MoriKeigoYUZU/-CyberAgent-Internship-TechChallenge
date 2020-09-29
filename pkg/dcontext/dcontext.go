package dcontext

import (
	"context"
)

type key string

const (
	userNameKey key = "userName"
)

// SetUserName ContextへユーザNameを保存する
func SetUserName(ctx context.Context, userName string) context.Context {
	return context.WithValue(ctx, userNameKey, userName)
}

// GetUserNameFromContext ContextからユーザNameを取得する
func GetUserNameFromContext(ctx context.Context) string {
	var userName string
	if ctx.Value(userNameKey) != nil {
		userName = ctx.Value(userNameKey).(string)
	}
	return userName
}
