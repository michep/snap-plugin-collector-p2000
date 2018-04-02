package parser

import (
	"fmt"
	"testing"
	"time"
)

func Test_Login(t *testing.T) {
	start := time.Now()
	cl := NewClient("http://172.16.18.31:80", "0e4997806bb599dec1864e034f9e59f9")
	cl.Login()
	fmt.Println(time.Now().Sub(start))
}

func Test_ControllerStatisticsAPI(t *testing.T) {
	start := time.Now()
	cl := NewClient("http://172.16.18.31:80", "0e4997806bb599dec1864e034f9e59f9")
	cl.Login()
	data, _ := cl.GetControllerStatistics()
	fmt.Printf("%+v\n", len(data))
	fmt.Println(time.Now().Sub(start))
}

func Test_HostPortStatisticsAPI(t *testing.T) {
	start := time.Now()
	cl := NewClient("http://172.16.18.31:80", "0e4997806bb599dec1864e034f9e59f9")
	cl.Login()
	data, _ := cl.GetHostPortStatistics()
	fmt.Printf("%+v\n", len(data))
	fmt.Println(time.Now().Sub(start))
}

func Test_DiskStatisticsAPI(t *testing.T) {
	start := time.Now()
	cl := NewClient("http://172.16.18.31:80", "0e4997806bb599dec1864e034f9e59f9")
	cl.Login()
	data, _ := cl.GetDiskStatistics()
	fmt.Printf("%+v\n", len(data))
	fmt.Println(time.Now().Sub(start))
}

func Test_VdiskStatisticsAPI(t *testing.T) {
	start := time.Now()
	cl := NewClient("http://172.16.18.31:80", "0e4997806bb599dec1864e034f9e59f9")
	cl.Login()
	data, _ := cl.GetVdiskStatistics()
	fmt.Printf("%+v\n", len(data))
	fmt.Println(time.Now().Sub(start))
}
