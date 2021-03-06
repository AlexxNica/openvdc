// +build linux

package lxc

import (
	"testing"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/axsh/openvdc/hypervisor"
	"github.com/axsh/openvdc/model"
	"github.com/stretchr/testify/assert"
	"github.com/Sirupsen/logrus"
	lxc "gopkg.in/lxc/go-lxc.v2"
)

func TestProviderRegistration(t *testing.T) {
	p, _ := hypervisor.FindProvider("lxc")
	if p == nil {
		t.Error("lxc provider is not registered.")
	}
}

func TestLXCHypervisorDriver(t *testing.T) {
	t.Skipf("Currently skipping this test because it requires too many outside dependencies. Will rewrite as integration test later.")

	p, _ := hypervisor.FindProvider("lxc")
	lxc, _ := p.CreateDriver("lxc-test")
	err := lxc.CreateInstance(&model.Instance{}, &model.LxcTemplate{})
	if err != nil {
		t.Error(err)
	}
	err = lxc.StartInstance()
	if err != nil {
		t.Error(err)
	}
	err = lxc.StopInstance()
	if err != nil {
		t.Error(err)
	}
	err = lxc.DestroyInstance()
	if err != nil {
		t.Error(err)
	}
}

const lxcConfTemplate = `
# Template used to create this container: /usr/share/lxc/templates/lxc-download
# Parameters passed to the template: --dist ubuntu --release trusty --arch amd64
# For additional config options, please look at lxc.container.conf(5)

# Distribution configuration
lxc.include = /usr/share/lxc/config/ubuntu.common.conf
lxc.arch = x86_64

# Container specific configuration
lxc.rootfs = /var/lib/lxc/i-0000000001/rootfs
lxc.utsname = i-0000000001

# Network configuration
lxc.network.type = veth
lxc.network.flags = up
lxc.network.link = virbr0
`

func TestLXCHypervisorDriver_modifyConf(t *testing.T) {
	assert := assert.New(t)
	lxcpath, err := ioutil.TempDir("/var/tmp", "")
	defer os.RemoveAll(lxcpath)
	c, err := lxc.NewContainer("lxc-test", lxcpath)
	assert.NoError(err)
	lxcdrv := &LXCHypervisorDriver{
		log: logrus.NewEntry(logrus.New()),
		container: c,
		template: lxc.BusyboxTemplateOptions,
	}
	os.MkdirAll(filepath.Join(lxcpath, "lxc-test"), 0755)
	ioutil.WriteFile(filepath.Join(lxcpath, "lxc-test", "config"), []byte(lxcConfTemplate), 0644)
	err = lxcdrv.modifyConf(&model.LxcTemplate{
		Vcpu: 1,
		MemoryGb: 256,
		Interfaces: []*model.LxcTemplate_Interface{
			&model.LxcTemplate_Interface{
				Type: "veth",
				Ipv4Addr: "192.168.1.1",
			},
			&model.LxcTemplate_Interface{
				Type: "veth",
				Macaddr: "xx:xx:xx:44:55:66",
				Ipv4Addr: "192.168.1.2",
			},
		},
	})
	assert.NoError(err)
	net_type := c.ConfigItem("lxc.network.type")
	assert.NotZero(len(net_type), "lxc.network.type does not apper")
	net_ipv4 := c.ConfigItem("lxc.network.ipv4")
	assert.NotZero(len(net_ipv4), "lxc.network.ipv4 does not apper")
	net_hwad := c.ConfigItem("lxc.network.hwaddr")
	assert.NotZero(len(net_hwad), "lxc.network.hwaddr does not apper")
}
