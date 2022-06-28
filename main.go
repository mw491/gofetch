package main

import (
	"bufio"
	"fmt"
	"math"
	"os/exec"
	"strings"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/gookit/color"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
)

var pengiun string = `    .--.    
   |^ ^ |   
   | <  |    
  //   \ \  
 (|     | ) 
/'|_   _/'\ 
\___)=(___/ 
`

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getInfo() []string {
	user, err := exec.Command("whoami").Output()
	checkErr(err)
	hostname, err := exec.Command("hostname").Output()
	checkErr(err)
	userathostname := color.Gray.Sprintf("%s@%s\n", strings.Replace(string(user), "\n", "", -1), strings.Replace(string(hostname), "\n", "", -1))

	platform, _, version, err := host.PlatformInformation()
	checkErr(err)

	// kernel, err := exec.Command("uname", "-r").Output()
	kernel, err := host.KernelVersion()
	checkErr(err)

	kernelarch, err := host.KernelArch()
	checkErr(err)

	cpuinfo, err := cpu.Info()
	checkErr(err)

	diskusage, err := disk.Usage("/")
	checkErr(err)
	diskusedpercent := math.Round(diskusage.UsedPercent)
	diskfree := humanize.Bytes(diskusage.Free)

	// uptime, err := exec.Command("uptime", "--pretty").Output()
	uptime, err := host.Uptime()
	checkErr(err)
	fmtuptime, err := time.ParseDuration(fmt.Sprint(uptime) + "s")
	checkErr(err)

	distro := color.Blue.Sprintf("distro:\t%v\n", color.Normal.Sprintf("%v %v", platform, version))
	kernel_display := color.Blue.Sprintf("kernel:\t%v\n", color.Normal.Sprintf("%s", kernel))
	arch := color.Blue.Sprintf("arch:\t%v\n", color.Normal.Sprint(kernelarch))
	cpu := color.Blue.Sprintf("cpu:\t%v\n", color.Normal.Sprint(cpuinfo[0].ModelName))
	df := color.Blue.Sprintf("disk:\t%v\n", color.Normal.Sprintf("%v%% used. %v free", diskusedpercent, diskfree))
	uptime_display := color.Blue.Sprintf("uptime:\t%v\n", color.Normal.Sprintf("%s", fmtuptime.String()))

	return []string{userathostname, kernel_display, arch, distro, cpu, df, uptime_display}
}

func printascii() {
	info := getInfo()

	scanner := bufio.NewScanner(strings.NewReader(pengiun))
	for i := 0; scanner.Scan(); i++ {
		// color.Yellow.Println(scanner.Text())
		color.Yellow.Printf("%v\t%v", scanner.Text(), info[i])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println()
	printascii()
	fmt.Println()
}
