package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
)

func try_ping(host string, count int) int {
	cmd := exec.Command("ping", host, "-c", "1", "-s", strconv.Itoa(count), "-M", "do") //"-s", "100", "-M", "do"
	_, err := cmd.Output()
	//fmt.Println(err.Error())
	if err != nil {
		//fmt.Println(err.Error())
		//return
		return 1
	}
	return 0
}

func main() {
	hostPtr := flag.String("host", "ya.ru", "a string")
	flag.Parse()
	host := *hostPtr

	

	if try_ping(host, 0) == 1 {
		fmt.Println("Wrong host - ", host)
		return
	}

	//Find upper border
	
	R := 1
	for try_ping(host, R) == 0 {
		R = R * 2
	}

	L := 0
	for R > L+1 {
		m := (R + L) / 2
		if try_ping(host, m) == 1 {
			R = m
		} else {
			L = m
		}
	}

	fmt.Println("MSU is", L+28) // 28 is a size of header

}
