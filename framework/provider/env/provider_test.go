package env

import (
	"testing"

	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/provider/app"
	tests "github.com/chenbihao/gob/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGobEnvProvider(t *testing.T) {
	Convey("test gob env normal case", t, func() {
		basePath := tests.BasePath
		c := framework.NewGobContainer()
		sp := &app.GobAppProvider{BaseFolder: basePath}

		err := c.Bind(sp)
		So(err, ShouldBeNil)

		sp2 := &EnvProvider{}
		err = c.Bind(sp2)
		So(err, ShouldBeNil)

		envServ := c.MustMake(contract.EnvKey).(contract.Env)
		So(envServ.AppEnv(), ShouldEqual, "dev")
		// So(envServ.Get("DB_HOST"), ShouldEqual, "127.0.0.1")
		// So(envServ.AppDebug(), ShouldBeTrue)
	})
}
