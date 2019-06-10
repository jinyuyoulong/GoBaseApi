package session

import (
	bservice "project-api/src/bootstrap/service"

	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
)

func SessionGet(ctx context.Context, key string) string {
	var value string
	container := bservice.GetDi().Container
	container.Invoke(func(sess *sessions.Sessions) {
		value = sess.Start(ctx).GetString(key)
	})
	return value
}
func SessionSet(ctx context.Context, key string, value string) {
	container := bservice.GetDi().Container
	container.Invoke(func(sess *sessions.Sessions) {
		sess.Start(ctx).Set(key, value)
	})
}

func SessionDelete(ctx context.Context, key string) {
	container := bservice.GetDi().Container
	container.Invoke(func(sess *sessions.Sessions) {
		sess.Start(ctx).Delete(key)
	})
}
