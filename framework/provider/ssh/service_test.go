package ssh

import (
	tests "github.com/chenbihao/gob/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSSHService_Load(t *testing.T) {
	container := tests.InitBaseConfLogContainer()

	Convey("test get client", t, func() {
		ssh, err := NewSSHService(container)
		So(err, ShouldBeNil)
		service, ok := ssh.(*SSHService)
		So(ok, ShouldBeTrue)
		client, err := service.GetClient(WithConfigPath("ssh.web-pwd"))
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)
		session, err := client.NewSession()
		So(err, ShouldBeNil)
		out, err := session.Output("pwd")
		So(err, ShouldBeNil)
		So(out, ShouldNotBeNil)
		session.Close()

		client, err = service.GetClient(WithConfigPath("ssh.web-key"))
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)
		session, err = client.NewSession()
		So(err, ShouldBeNil)
		out, err = session.Output("pwd")
		So(err, ShouldBeNil)
		So(out, ShouldNotBeNil)
		session.Close()
	})
}
