package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type OsListInfo struct {
	OsList []OsInfo `yaml:"os_list"`
}

type OsInfo struct {
	Version      string   `yaml:"version"`
	IsoPath      string   `yaml:"iso_path"`
	PkgMgr       string   `yaml:"pkg_mgr"`
	PkgInstaller string   `yaml:"pkg_installer"`
	Name         string   `yaml:"name"`
	Arch         string   `yaml:"arch"`
	Type         string   `yaml:"type"`
	Id           string   `yaml:"id"`
	PkgList      []string `yaml:"pkg_list"`
}

var DefaultConfig = []byte(`os_list:
  - name: rocky8
    arch: x86_64
    type: linux
    id: rocky
    version: 8.6
    iso_path: rockylinux8-amd64.iso
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

  - name: debian13
    arch: x86_64
    type: linux
    id: debian
    version: 13
    iso_path: debian13-amd64.iso
    pkg_mgr: dpkg
    pkg_installer: apt
    pkg_list:
      - socat
`)

func LoadConfig(configPath string) []OsInfo {
	var osInfo OsListInfo

	data, err := os.ReadFile(configPath)
	if err != nil {
		data = DefaultConfig
	}

	err = yaml.Unmarshal(data, &osInfo)
	if err != nil {
		log.Fatal(err)
	}
	return osInfo.OsList
}
