package redis

import (
	"context"
	tests "github.com/chenbihao/gob/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestRedisService_Load(t *testing.T) {
	container := tests.InitBaseConfLogContainer()

	Convey("test get client", t, func() {
		redis, err := NewRedisService(container)
		So(err, ShouldBeNil)
		service, ok := redis.(*RedisService)
		So(ok, ShouldBeTrue)
		client, err := service.GetClient(WithConfigPath("redis.write"))
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)

		ctx := context.Background()
		pong, err := client.Ping(ctx).Result()
		So(err, ShouldBeNil)
		So(pong, ShouldNotBeNil)

		err = client.Set(ctx, "foo", "bar", 1*time.Hour).Err()
		So(err, ShouldBeNil)
		val, err := client.Get(ctx, "foo").Result()
		So(err, ShouldBeNil)
		So(val, ShouldEqual, "bar")
		err = client.Del(ctx, "foo").Err()
		So(err, ShouldBeNil)
	})
}
