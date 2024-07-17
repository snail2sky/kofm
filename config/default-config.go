package config

var DefaultConfig = []byte(
	`# other iso can be find https://github.com/kubesphere/kubekey/releases
os_list:
  - name: rocky8
    arch: x86_64
    type: linux
    id: rocky
    version: 8.6
    iso_path: el8-amd64.iso
    pkg_mgr: rpm
    pkg_installer: dnf
    pkg_list:
      - socat
      - sudo
      - curl
      - openssl
      - ebtables
      - ipset
      - ipvsadm
      - conntrack
      - keepalived
      - haproxy

  - name: rocky9
    arch: x86_64
    type: linux
    id: rocky
    version: 9.4
    iso_path: el9-amd64.iso
    pkg_mgr: rpm
    pkg_installer: dnf
    pkg_list:
      - socat
      - sudo
      - curl
      - openssl
      - ebtables
      - ipset
      - ipvsadm
      - conntrack
      - keepalived
      - haproxy

  - name: ubuntu2004
    arch: x86_64
    type: linux
    id: ubuntu
    version: 20.04
    iso_path: ubuntu2004-amd64.iso
    pkg_mgr: dpkg
    pkg_installer: apt
    pkg_list:
      - socat
      - conntrack
`)
