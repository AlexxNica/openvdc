// +build acceptance

package tests

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestLXCInstance(t *testing.T) {
	stdout, _ := RunCmdAndReportFail(t, "openvdc", "run", "centos/7/lxc", `{"interfaces":[{"type":"veth", "bridge":"linux"}]}`)
	instance_id := strings.TrimSpace(stdout.String())

	_, _ = RunCmdAndReportFail(t, "openvdc", "show", instance_id)
	RunCmdAndReportFail(t, "openvdc", "wait", instance_id, "RUNNING")
	RunSshWithTimeoutAndReportFail(t, executor_lxc_ip, "sudo lxc-info -n "+instance_id, 10, 5)
	//TODO: Run only once after we've waited for RUNNING
	_, _ = RunCmdWithTimeoutAndReportFail(t, 10, 5, "openvdc", "destroy", instance_id)
	RunCmdAndReportFail(t, "openvdc", "wait", instance_id, "TERMINATED")

	//TODO: Don't rely on the standard 'command failed' error.
	//It's more clear to say "container didn't get destroyed after calling openvdc destroy"
	RunSshWithTimeoutAndExpectFail(t, executor_lxc_ip, "sudo lxc-info -n "+instance_id, 10, 5)
}

func TestLXCInstance_NICx2(t *testing.T) {
	stdout, _ := RunCmdAndReportFail(t, "openvdc", "run", "centos/7/lxc",
		`{"interfaces":[{"type":"veth", "bridge":"linux"}, {"type":"veth", "bridge":"linux"}]}`)
	instance_id := strings.TrimSpace(stdout.String())

	RunCmdAndReportFail(t, "openvdc", "show", instance_id)
	RunCmdAndReportFail(t, "openvdc", "wait", instance_id, "RUNNING")
	RunSshWithTimeoutAndReportFail(t, executor_lxc_ip, "sudo lxc-info -n "+instance_id, 10, 5)
	stdout, _, err := RunSsh(executor_lxc_ip, "/usr/sbin/brctl show br0")
	if err != nil {
		t.Error(err)
	}
	/* Expected Output
	bridge name bridge id         STP  enabled interfaces
	br0         8000.fe0e49305657	no		       i-0000000001_00
	                                           i-0000000001_01
	*/
	lines := bufio.NewScanner(stdout)
	lines.Scan() // Skip "brctl show" header
	lines.Scan()
	line_if00 := bufio.NewScanner(bytes.NewReader(lines.Bytes()))
	line_if00.Split(bufio.ScanWords)
	// "br0          8000.080027a02faf       no              i-00000000_00"
	line_if00.Scan()
	line_if00.Scan()
	line_if00.Scan()
	line_if00.Scan()
	t.Log(line_if00.Text())
	if line_if00.Text() != instance_id+"_00" {
		t.Errorf("Interface %s is not attached", instance_id+"_00")
		if testing.Verbose() {
			t.Log(stdout.String())
		}
	}
	lines.Scan()
	line_if01 := bufio.NewScanner(bytes.NewReader(lines.Bytes()))
	line_if01.Split(bufio.ScanWords)
	// "                                    i-00000000_01"
	line_if01.Scan()
	t.Log(line_if01.Text())
	if line_if01.Text() != instance_id+"_01" {
		t.Errorf("Interface %s is not attached", instance_id+"_01")
		if testing.Verbose() {
			t.Log(stdout.String())
		}
	}

	RunCmdWithTimeoutAndReportFail(t, 10, 5, "openvdc", "destroy", instance_id)
	RunCmdAndReportFail(t, "openvdc", "wait", instance_id, "TERMINATED")
}
