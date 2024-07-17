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
