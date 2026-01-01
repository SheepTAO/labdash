package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type SystemInfo struct {
	Hostname string  `json:"hostname"`
	OS       string  `json:"os"`
	Kernel   string  `json:"kernel"`
	Uptime   string  `json:"uptime"`
	LoadAvg  float64 `json:"loadAvg"`
}

func GetStaticSystemInfo() SystemInfo {
	info := SystemInfo{}

	// Get Hostname
	info.Hostname, _ = os.Hostname()

	// Get Kernel Version (uname -r)
	out, _ := exec.Command("uname", "-r").Output()
	info.Kernel = strings.TrimSpace(string(out))

	// Get OS Version Name
	info.OS = "Linux"
	if data, err := os.ReadFile("/etc/os-release"); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "PRETTY_NAME=") {
				// Remove quotes
				info.OS = strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), "\"")
				break
			}
		}
	}

	return info
}

// Get Uptime
func GetUptime() string {
	data, _ := os.ReadFile("/proc/uptime")
	parts := strings.Fields(string(data))
	if len(parts) > 0 {
		seconds, _ := strconv.ParseFloat(parts[0], 64)
		days := int(seconds) / 86400
		hours := int(seconds) % 86400 / 3600
		mins := int(seconds) % 3600 / 60
		return strconv.Itoa(days) + "d " + strconv.Itoa(hours) + "h " + strconv.Itoa(mins) + "m"
	}
	return "--"
}

// Get System Load Average (1 min)
func GetLoadAvg() float64 {
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return 0.0
	}

	parts := strings.Fields(string(data))
	if len(parts) > 0 {
		load, _ := strconv.ParseFloat(parts[0], 64)
		return load
	}
	return 0.0
}
