#!/bin/bash

export ENV_ROOTDIR="$(cd "$(dirname $(readlink -f "$0"))/.." && pwd -P)"
export NODE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TMP_ROOT="${NODE_DIR}/tmp_root"

. "${ENV_ROOTDIR}/config.source"
. "${NODE_DIR}/vmspec.conf"
. "${ENV_ROOTDIR}/ind-steps/common.source"

mesos_agent=true

IND_STEPS=(
    "box"
    "ssh"
    "epel"
    "lxc"
    "mesosphere"
    "mesos"
    "zookeeper"
)

build "${IND_STEPS[@]}"

install_openvdc_yum_repo

# This is not part of the ind-steps because we don't want OpenVDC installed in
# the cached images. We want a clean cache without OpenVDC so we can install a
# different version to test every the CI runs.
#
# The CLI could be installed in any node but let's do it here. This is the best
# place because we want to make sure everything works with other addresses than
# 127.0.0.1
install_openvdc_yum_repo
install_yum_package_over_ssh "openvdc-executor openvdc-cli"
