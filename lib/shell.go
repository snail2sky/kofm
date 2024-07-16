package lib

import (
	"log"
	"os"
	"os/exec"
)

var pkgMgrCmd = [...]string{
	"yum",
	"dnf",
	"apt",
	"apt-get",
}

func GetPkgMgr() string {
	for _, cmd := range pkgMgrCmd {
		err := exec.Command("which", cmd).Run()
		if err == nil {
			switch cmd {
			case "yum":
			case "dnf":
				return "rpm"
			case "apt":
			case "apt-get":
				return "deb"
			}
		}
	}
	return ""
}

func Mkdir(dirs ...string) {
	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}
