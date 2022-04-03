package main

// imports
import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//main logic
func main() {

	//flag options
	ipPtr := flag.String("ip", "x.x.x.x", "[REQUIRED] target IP(s) to be scanned.")

	flag.Parse()

	mkdir(*ipPtr)     //working mkdir
	pingsweep(*ipPtr) //working pingsweep
	//cat(*ipPtr, "pingsweep-alive-hosts", ".txt")
	//grepAliveHosts(cat(*ipPtr, "pingsweep-alive-hosts", ".txt"))
	verifyIP(*ipPtr)
} //end main

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

//func to print ascii banner
func banner() {

	// color variables
	colorReset := "\033[0m" //reset term color
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"
	colorPurple := "\033[35m"
	bold := "\u001b[1m"
	underline := "\u001b[4m"
	reset := "\u001b[0m"
	whitebg := "\u001b[47;1m"

	fmt.Println(colorYellow + "         *                 *                  *              *        ")
	fmt.Println("                             " + bold + whitebg + colorGreen + "Go" + colorYellow + "Scan" + reset + colorYellow + "                    *             *")
	fmt.Println("                        *            *                             " + colorReset + "___")
	fmt.Println(colorReset + "  \033[33m*               \033[33m*                                          \033[31m|     " + colorReset + "| |")
	fmt.Println(colorReset + "        \033[33m*" + "              \033[37m" + colorReset + "_________\033[31m" + colorRed + "##" + "                 \033[33m*" + "        \033[31m/" + colorRed + " \\" + "    " + colorReset + "| |" + reset + "") //working
	fmt.Println(colorReset + "                      " + "\033[33m@" + colorReset + "\\\\\\\\\\\\\\\\\\" + colorRed + "##    \033[33m*     |              " + colorRed + "|" + colorReset + bold + "--o" + reset + colorRed + "|" + colorReset + "===|-|")
	fmt.Println(colorReset + "  \033[33m*                  " + colorYellow + "@@@" + colorReset + "\\\\\\\\\\\\\\\\" + colorRed + "##" + colorReset + "\\" + colorGreen + "       \\" + colorYellow + "|" + colorGreen + "/" + colorYellow + "|" + colorGreen + "/            " + colorRed + "|" + colorReset + bold + "---" + reset + colorRed + "|   " + colorReset + "|" + colorCyan + "j" + colorReset + "|")
	fmt.Println(colorYellow + "                    @@ @@" + colorReset + "\\\\\\\\\\\\\\\\\\\\\\    " + colorGreen + "\\" + colorYellow + "|" + colorGreen + "\\\\" + colorYellow + "|" + colorGreen + "//" + colorYellow + "|" + colorGreen + "/     " + colorYellow + "*   " + colorRed + "/     \\  " + colorReset + "|" + colorCyan + "a" + colorReset + "| ")
	fmt.Println(colorYellow + "             *     @@@@@@@" + colorReset + "\\\\\\\\\\\\\\\\\\\\\\    " + colorGreen + "\\" + colorYellow + "|" + colorGreen + "\\" + colorYellow + "|" + colorGreen + "/" + colorYellow + "|" + colorGreen + "/         " + colorRed + "|  " + "U" + colorRed + "    | " + colorReset + "|" + colorCyan + "k" + colorReset + "| ")
	fmt.Println(colorYellow + "                  @@@@@@@@@----------|    " + colorGreen + "\\\\" + colorYellow + "|" + colorGreen + "//          " + colorRed + "|  " + colorReset + bold + "S    " + reset + colorRed + "|" + colorReset + "=|" + colorCyan + "e" + colorReset + "| ")
	fmt.Println(colorBlue + "       __         " + colorYellow + "@@ @@@ @@__________|     " + colorGreen + "\\" + colorYellow + "|" + colorGreen + "/           " + colorRed + "|  " + colorBlue + "A    " + colorRed + "| " + colorReset + "| | ")
	fmt.Println(colorReset + "  " + colorBlue + "____|_" + colorReset + "@" + colorBlue + "|_       " + colorYellow + "@@@@@@@@@__________|     " + colorGreen + "\\" + colorYellow + "|" + colorGreen + "/           " + colorRed + "|_______| " + colorReset + "|" + colorCyan + "S" + colorReset + "| ")
	fmt.Println(colorReset + "=" + colorBlue + "|__ _____ |" + colorReset + "=     " + colorYellow + "@@@@ " + colorReset + "." + colorYellow + "@@@__________|      |             " + colorRed + "|" + colorYellow + "@" + colorRed + "| |" + colorYellow + "@" + colorRed + "|  " + colorReset + "|_| ")
	fmt.Println(colorReset + colorGreen + "____" + colorReset + "0" + colorGreen + "_____" + colorReset + "0" + colorGreen + "__\\|/__" + colorYellow + "@@@@" + colorReset + "__" + colorYellow + "@@@" + colorYellow + "__________|" + colorGreen + "_\\|/__" + colorYellow + "|" + colorGreen + "___\\|/__\\|/___________" + colorReset + "|" + colorGreen + "_" + colorReset + "|" + colorGreen + "_" + colorReset)
	fmt.Println(colorGreen + "  \\|/                " + colorReset + "/ /       " + colorGreen + "\\|/                              \\|/    ")
	fmt.Println(colorYellow + "+" + colorBlue + "---------------------------------------------------------------------" + colorYellow + "+")
	fmt.Println(colorBlue + "|   " + colorGreen + "Go" + colorYellow + "Scan" + colorReset + "- fast network enumeration (∩｀-´)⊃" + colorYellow + "━   " + colorPurple + "☆" + colorCyan + "ﾟ" + colorYellow + "." + colorPurple + "*" + colorGreen + "･" + colorYellow + "｡" + colorCyan + "ﾟ   " + colorRed + "           " + colorBlue + "|")
	fmt.Println(colorBlue + "|   Written by: " + bold + underline + colorWhite + "Jake Scheetz" + reset + "          " + colorGreen + "<you>          " + colorPurple + "☆" + colorCyan + "ﾟ" + colorYellow + "." + colorPurple + "*" + colorGreen + "･" + colorYellow + "｡" + colorCyan + "ﾟ          " + colorBlue + "|")
	fmt.Println(colorBlue + "|   " + colorCyan + "Twitter" + colorWhite + ": " + colorYellow + "@" + bold + colorWhite + "FindingUrPasswd" + reset + "                 " + colorPurple + "<cyberMagic>    " + colorReset + "(" + colorCyan + "ᵟຶ︵ ᵟຶ" + colorReset + ") " + colorBlue + "|")
	fmt.Println(colorBlue + "|   " + colorRed + "You" + colorReset + "Tube: " + bold + "youtube.com/c/FindingUrPasswd" + reset + "                   " + colorRed + "<badGuys>" + colorBlue + "|")
	fmt.Println(colorYellow + "+" + colorBlue + "---------------------------------------------------------------------" + colorYellow + "+" + colorReset)
}
