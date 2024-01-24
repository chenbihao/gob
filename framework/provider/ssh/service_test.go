package ssh

import (
	tests "github.com/chenbihao/gob/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSSHService_Load(t *testing.T) {
	container := tests.InitBaseConfLogContainer()

	Convey("test get client", t, func() {
		hadeRedis, err := NewSSHService(container)
		So(err, ShouldBeNil)
		service, ok := hadeRedis.(*SSHService)
		So(ok, ShouldBeTrue)
		client, err := service.GetClient(WithConfigPath("ssh.web-01"))
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)
		session, err := client.NewSession()
		So(err, ShouldBeNil)
		out, err := session.Output("pwd")
		So(err, ShouldBeNil)
		So(out, ShouldNotBeNil)
		session.Close()
	})
}
