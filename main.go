package main

// imports
import (
	banners "GoScan/bannerArt"
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

//main logic
func main() {
	banners.PrintWelcomeBanner()
	startTime := time.Now()

	//flag options
	//ipPtr := flag.String("ip", "x.x.x.x", "[REQUIRED] target IP(s) to be scanned.")

	//flag.Parse()

	//mkdir(*ipPtr)     //working mkdir
	//pingsweep(*ipPtr) //working pingsweep
	//cat(*ipPtr, "pingsweep-alive-hosts", ".txt")
	//grepAliveHosts(cat(*ipPtr, "pingsweep-alive-hosts", ".txt"))
	//verifyIP(*ipPtr)
	//output := grep("./127.0.0.1/pingsweep-alive-hosts.txt", []byte("Status: Up"))
	//fmt.Printf("total %d\n", output)
	elapsedExecutionTime := time.Since(startTime)
	fmt.Println("Total Execution time: ", elapsedExecutionTime.Seconds(), "seconds. Now that's fast!")
}

// OS commands to execute
func mkdir(dir string) {
	app := "mkdir"
	arg0 := "./" + dir

	cmd := exec.Command(app, arg0)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func pingsweep(ipPtr string) {
	app := "nmap"
	arg0 := "-sn"
	arg1 := "-vv"
	arg2 := "-T4"
	arg3 := ipPtr
	arg4 := "-oG"
	arg5 := "./" + ipPtr + "/pingsweep-alive-hosts.txt"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//THIS IS GETTING NMAP OUTPUT
}

func verifyIP(dir string) {

	fileHandle, _ := os.Open("./" + dir + "/pingsweep-alive-hosts.txt")
	defer fileHandle.Close()
	s := bufio.NewScanner(fileHandle)

	fmt.Println("[+] The following hosts were verified to be alive:")
	for s.Scan() {
		if strings.Contains(s.Text(), "Status: Up") {
			println(s.Text())
		}
		//getting the lines with Up status
	}
}

func grep(file string, pattern []byte) int64 {
	patCount := int64(0)
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if bytes.Contains(scanner.Bytes(), pattern) {
			patCount++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return patCount
}
