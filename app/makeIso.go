package app

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/snail2sky/kofm/config"
	"github.com/snail2sky/kofm/lib"
	"log"
	"os"
	"os/exec"
	"path"
)

type PkgManager interface {
	Install(pkgs ...string)
	DownloadPkg(rootDir string)
	MakeRepo(dir string) error
	GetPkgDir() string
	SetInitializer(*Initializer)
}

type Initializer struct {
	WorkerDir string
	*config.OsInfo
	PkgManager
}

func NewInitializer(workerDir string, osList []config.OsInfo) *Initializer {
	osInfo := GetOSConfig(osList)
	pkgMgr := NewPkgManager(osInfo)

	return &Initializer{
		WorkerDir:  workerDir,
		OsInfo:     osInfo,
		PkgManager: pkgMgr,
	}
}

func GetOSConfig(osList []config.OsInfo) *config.OsInfo {
	info, _ := host.Info()
	for _, osInfo := range osList {
		log.Printf("%#v\n", osInfo)
		log.Printf("%#v\n", info)
		if osInfo.Version == info.PlatformVersion && osInfo.Id == info.Platform {
			return &osInfo
		}
	}
	return nil
}

func NewPkgManager(osInfo *config.OsInfo) PkgManager {
	var pkgMgr PkgManager
	info, _ := host.Info()
	if osInfo == nil {
		log.Fatal("Not supported OS:", info.Platform, info.PlatformVersion)
	}
	switch osInfo.PkgInstaller {
	case "yum":
		pkgMgr = &RpmPkgMgr{
			pkgDir:       "rpms",
			pkgInstaller: "yum",
		}
	case "dnf":
		pkgMgr = &RpmPkgMgr{
			pkgDir:       "rpms",
			pkgInstaller: "dnf",
		}
	case "apt":
	case "apt-get":
		pkgMgr = &DpkgMgr{
			pkgDir:       "dpkgs",
			pkgInstaller: "apt",
		}
	}
	return pkgMgr
}

func (i *Initializer) Mkdir(dirs ...string) {
	log.Println("mkdir", dirs)
	lib.Mkdir(dirs...)
}

func (i *Initializer) GetKK(kkZone string) {
	oldPwd, _ := os.Getwd()
	defer os.Chdir(oldPwd)
	err := os.Chdir(i.WorkerDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("check kk exists")
	_, err = os.Stat("kk")
	if err == nil {
		log.Println("The kk already exists. No repeated installation kk")
		return
	}
	log.Println("installing kk")
	cmd := exec.Command("sh", "-c", "curl -sfL https://get-kk.kubesphere.io | VERSION=v3.0.13 sh -")
	cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", "KKZONE", kkZone))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))
}

func (i *Initializer) MkISO() {
	oldPwd, _ := os.Getwd()
	defer os.Chdir(oldPwd)
	err := os.Chdir(i.WorkerDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("mkiso", i.OsInfo.IsoPath)
	cmd := exec.Command("sh", "-c", fmt.Sprintf("mkisofs -R -o %s %s", i.OsInfo.IsoPath, i.GetPkgDir()))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))
}

type RpmPkgMgr struct {
	pkgDir       string
	pkgInstaller string
	*Initializer
}

type DpkgMgr struct {
	pkgDir       string
	pkgInstaller string
	*Initializer
}

func (i *RpmPkgMgr) Install(pkgs ...string) {
	for index, pkg := range pkgs {
		log.Println(index, "installing", pkg)
		err := exec.Command("yum", "-y", "install", pkg).Run()
		if err != nil {
			log.Printf("Error installing %s: %s\n", pkg, err)
		}
	}
}

func (i *RpmPkgMgr) DownloadPkg(rootDir string) {
	oldPwd, _ := os.Getwd()
	defer os.Chdir(oldPwd)
	err := os.Chdir(path.Join(rootDir, i.GetPkgDir()))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("downloading", i.Initializer.PkgList)
	for _, pkg := range i.Initializer.PkgList {
		log.Println("downloading", pkg)
		output, _ := exec.Command("repotrack", pkg).CombinedOutput()
		log.Println(string(output))

	}
}
func (i *RpmPkgMgr) MakeRepo(dir string) error {
	log.Println("making repo", dir)
	output, err := exec.Command("createrepo", dir).CombinedOutput()
	log.Println(string(output))
	return err
}

func (i *RpmPkgMgr) SetInitializer(initializer *Initializer) {
	i.Initializer = initializer
}

func (i *RpmPkgMgr) GetPkgDir() string {
	return i.pkgDir
}

func (i *DpkgMgr) Install(pkgs ...string) {
	for _, pkg := range pkgs {
		err := exec.Command("apt", "-y", "install", pkg).Run()
		if err != nil {
			log.Printf("Error installing %s: %s\n", pkg, err)
		}
	}
}
func (i *DpkgMgr) DownloadPkg(rootDir string) {
	log.Println(rootDir)
}
func (i *DpkgMgr) MakeRepo(dir string) error {
	log.Println(dir)
	return nil
}

func (i *DpkgMgr) SetInitializer(initializer *Initializer) {
	i.Initializer = initializer
}

func (i *DpkgMgr) GetPkgDir() string {
	return i.pkgDir
}
