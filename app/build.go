package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Builder struct {
	manifest string
	output   string
	name     string
	kkZone   string
	kk       string
}

func NewBuilder(manifest, output, kkZone, kk string) *Builder {
	return &Builder{
		manifest: manifest,
		output:   output,
		kkZone:   kkZone,
		kk:       kk,
	}
}

func (b *Builder) Build() {
	cmd := exec.Command(b.kk, "artifact", "export", "-m", b.manifest, "-o", b.output, "--debug", "--yes")
	cmd.Env = append(os.Environ(), fmt.Sprintf("KKZONE=%s", b.kkZone))
	log.Println("building artifact", cmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("artifact exited with error:", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("artifact exited with error:", err)
	}
	if err != nil {
		log.Println(err)
		return
	}
	defer stdout.Close()
	defer stderr.Close()

	if err := cmd.Start(); err != nil {
		log.Println(err)
		return
	}

	stdoutBuf := make([]byte, 1024)
	for {
		n, err := stdout.Read(stdoutBuf)

		if err != nil {
			break
		}
		fmt.Print(string(stdoutBuf[:n]))
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
		return
	}
}
