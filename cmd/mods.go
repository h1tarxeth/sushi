package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func packages_info() string {
	var package_count string
	pkgCommands := map[string]string{
		"pacman":    "pacman -Q | wc -l",
		"flatpak":   "flatpak list | wc -l",
		"snap":      "snap list | wc -l",
		"dpkg":      "dpkg --get-selections | wc -l",
		"xbps":      "xbps-query -l | awk '{ print $2 }' | xargs -n1 xbps-uhelper getpkgname | wc -l",
		"apt":       "apt list --installed | wc -l",
		"dnf":       "dnf list --installed | grep -c -v '^Installed Packages$'",
		"nix-store": "nix-store -q --references /var/run/current-system/sw | cut -d'-' -f2- | wc -l",
	}

	for pkg, command := range pkgCommands {
		if _, err := exec.LookPath(pkg); err == nil {
			cmd := exec.Command("bash", "-c", command)
			output, _ := cmd.Output()
			package_count += fmt.Sprintf("%s: %s ", pkg, strings.TrimSpace(string(output)))
		}
	}
	return package_count
}

func hostname() string {
	file, _ := os.ReadFile("/etc/hostname")
	hostname := strings.Split(string(file), "\n")
	return hostname[0]
}

func user_info() string {
	return os.Getenv("USER")
}

func shell_info() string {
	shell := strings.SplitN(os.Getenv("SHELL"), "/", 4)
	return shell[3]
}

func mem_info_MiB() string {
	var mem_arr [3]int
	file, _ := os.ReadFile("/proc/meminfo")
	line := strings.SplitN(string(file), "\n", 3)
	for i := 0; i < 3; i++ {
		col_index := strings.Index(line[i], ":")
		kB_index := strings.Index(line[i], "kB")
		men_str := strings.TrimSpace(line[i][col_index+1 : kB_index])
		mem_arr[i], _ = strconv.Atoi(men_str)
		mem_arr[i] /= 1024
	}
	mem := "" + strconv.Itoa(mem_arr[0]-mem_arr[2]) + " MiB | " + strconv.Itoa(mem_arr[0]) + " MiB"
	return mem
}

func cpu_info() string {
	file, _ := os.ReadFile("/proc/cpuinfo")
	line := strings.Split(string(file), "\n")
	cpu_str := strings.TrimSpace(line[4][strings.Index(line[4], ":")+1:])
	return cpu_str
}

func gpu_info() string {
	cmd := exec.Command("bash", "-c", "lspci | grep -E 'VGA|3D' | cut -d ':' -f3 | xargs")
	gpu, _ := cmd.Output()
	return strings.Trim(string(gpu), "\n")
}

func res_info() string {
	cmd := exec.Command("bash", "-c", "xrandr | grep '*' | awk '{print $1}'")
	res, _ := cmd.Output()
	return strings.Trim(string(res), "\n")
}

func ip_info() string {
	cmd := exec.Command("bash", "-c", "ip addr show wlp5s0 | grep 'inet ' | awk '{print $2}'")
	ip, _ := cmd.Output()
	return strings.Trim(string(ip), "\n")
}

func kernal_info() string {
	file, _ := os.ReadFile("/proc/version")
	line := strings.SplitN(string(file), " ", 4)
	return line[2]
}

func os_info() string {
	file, _ := os.ReadFile("/etc/os-release")
	os_name := strings.SplitN(string(file), "\"", 5)
	return os_name[3]
}

func uptime() string {
	file, _ := os.ReadFile("/proc/uptime")
	line := strings.TrimSpace(string(file[:strings.Index(string(file), ".")]))
	time, _ := strconv.Atoi(line)
	time_min := (time % 3600) / 60
	time_hor := (time / 3600) % 24
	time_day := (time / 3600) / 60

	if time_day == 0 && time_hor == 0 {
		return fmt.Sprintf("%vm", time_min)
	} else if time_day == 0 {
		return fmt.Sprintf("%vh %vm", time_hor, time_min)
	} else {
		return fmt.Sprintf("%vd %vh %vmi", time_day, time_hor, time_min)
	}
}

func hostsystem() string {
	name, _ := os.ReadFile("/sys/devices/virtual/dmi/id/product_name")
	version, _ := os.ReadFile("/sys/devices/virtual/dmi/id/product_version")
	model, _ := os.ReadFile("/sys/firmware/devicetree/base/model")
	hostsys := fmt.Sprintf(
		"%s %s %s",
		strings.TrimSpace(string(name)),
		strings.TrimSpace(string(version)),
		strings.TrimSpace(string(model)),
	)
	return hostsys
}
