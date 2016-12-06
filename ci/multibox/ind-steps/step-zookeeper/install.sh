(
    $starting_step "Install openvdc on ${vm_name}"
    [[ -d "${TMP_ROOT}/opt/mesosphere/zookeper/bin" ]]
    $skip_step_if_already_done; set -ex
    sudo chroot ${TMP_ROOT} /bin/bash -c "yum install -y mesosphere-zookeeper"
) ; prev_cmd_failed
