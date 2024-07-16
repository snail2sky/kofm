package app

import (
	"os"
	"strings"
)

func LoadOSInfo() map[string]string {
	data, err := os.ReadFile("/etc/os-release")
	var osInfoMap = make(map[string]string)
	hostname, _ := os.Hostname()
	osInfoMap["HOSTNAME"] = hostname
	if err != nil {
		return nil
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.Contains(line, "=") {
			kvPair := strings.Split(line, "=")
			key := strings.TrimSpace(kvPair[0])
			value := strings.TrimSpace(kvPair[1])
			osInfoMap[key] = value
		}
	}
	return osInfoMap
}

