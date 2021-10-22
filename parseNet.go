package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
 I started to write this to be able to parse the /proc/<pid>/net/tcp
 and format the output like the lsof -n  -i TCP -a -p <pid> command.
 But then I realized that the sosreport does not bring /proc/<pid>/net/tcp
 files, so I just give up working on it.
*/

//   0: 00000000:23DC 00000000:0000 0A 00000000:00000000 00:00000000 00000000  1002        0 3564511 1 ffff9c45933c6b40 100 0 0 10 0
func parseNet(pid string) {

	// gather the command name
	cmdFile, err := ioutil.ReadFile("/proc/" + pid + "/comm")
	if err != nil {
		log.Fatal(err)
	}

	// strip the \n from command name
	cmdName := strings.TrimSuffix(string(cmdFile), "\n")

	// tcp connections file
	statFile, err := os.Open("/proc/" + pid + "/net/tcp")

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(statFile)

	// connection status
	status := make(map[string]string)
	status["01"] = "ESTABLISHED"
	status["02"] = "SYN_SENT"
	status["03"] = "SYN_RECV"
	status["04"] = "FIN_WAIT1"
	status["05"] = "FIN_WAIT2"
	status["06"] = "TIME_WAIT"
	status["07"] = "CLOSE"
	status["08"] = "CLOSE_WAIT"
	status["09"] = "LAST_ACK"
	status["0A"] = "LISTEN"
	status["0B"] = "CLOSING"
	status["0C"] = "NEW_SYN_RECV"
	status["0D"] = "MAX_STATES"

	// print header output
	fmt.Println("COMMAND\tPID\tUSER\tNAME")

	// regex for the line format to parse
	r := regexp.MustCompile(`^(?P<general>\s+\d+:\s+)(?P<src4>[[:xdigit:]]{2})(?P<src3>[[:xdigit:]]{2})(?P<src2>[[:xdigit:]]{2})(?P<src1>[[:xdigit:]]{2}):(?P<srcPort>[[:xdigit:]]{4})\s(?P<dst4>[[:xdigit:]]{2})(?P<dst3>[[:xdigit:]]{2})(?P<dst2>[[:xdigit:]]{2})(?P<dst1>[[:xdigit:]]{2}):(?P<dstPort>[[:xdigit:]]{4})\s+(?P<status>[[:xdigit:]]{2})\s+.*?\s+.*?\s+.*?\s+(?P<uid>\d+)(?P<rest>.*)`)

	// go through the file lines
	for scanner.Scan() {
		groups := map[string]string{}
		matches := r.FindStringSubmatch(scanner.Text())

		for i, name := range matches {
			if i == 0 {
				continue
			}
			groups[r.SubexpNames()[i]] = name
		}

		if len(groups) > 0 && groups["status"] == "01" {
			fmt.Printf("%v\t%v\t%v\t%v.%v.%v.%v:%v->%v.%v.%v.%v:%v (%v)\n", cmdName, pid, groups["uid"], convertStrHexToInt(groups["src1"]), convertStrHexToInt(groups["src2"]), convertStrHexToInt(groups["src3"]), convertStrHexToInt(groups["src4"]), convertStrHexToInt(groups["srcPort"]), convertStrHexToInt(groups["dst1"]), convertStrHexToInt(groups["dst2"]), convertStrHexToInt(groups["dst3"]), convertStrHexToInt(groups["dst4"]), convertStrHexToInt(groups["dstPort"]), status[groups["status"]])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func convertStrHexToInt(n string) int {
	num, _ := strconv.ParseInt(n, 16, 64)
	return int(num)
}
