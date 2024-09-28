package main

import (
	"fmt"
)

func main() {
	rs := "\033[0m"
	bg := "\033[90m"
	rd := "\033[91m"
	gr := "\033[92m"
	ye := "\033[93m"
	bl := "\033[94m"
	ma := "\033[95m"
	cy := "\033[96m"
	wh := "\033[97m"

	// ascii
	ascii := `
 ┓ •
 ┃ ┓ ┏┓ ┓┏ ┓┏
 ┗ ┗ ┛┗ ┗┻ ┛┗  `
	fmt.Println(cy + ascii + rs)
	fmt.Print(" "+ye+user_info()+rs, "@", gr+hostname()+rs, "\n")
	fmt.Println(" Os     ", cy+os_info()+rs)
	fmt.Println(" Kernel ", bl+kernel_info()+rs)
	fmt.Println(" Host   ", ma+hostsystem()+rs)
	fmt.Println(" Mem    ", ye+mem_info_MiB()+rs)
	fmt.Println(" pkgs   ", rd+packages_info()+rs)
	fmt.Println(" Uptime ", gr+uptime()+rs)
	fmt.Println(
		"  " + gr + "  " + ye + "  " + bl + "  " + ma + "  " + cy + "  " + wh + "  " + bg + "  " + rs,
	)
	// other functions
	// cpu_info()
	// gpu_info()
	// shell_info()
	// ip_info()
	// res_info()
}
