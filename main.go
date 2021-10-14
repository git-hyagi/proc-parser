package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	var statFile *os.File
	var err error

	// If a file is passed as argument it will be used
	// if not, use the current system /proc/stat
	if len(os.Args) == 2 {
		statFile, err = os.Open(os.Args[1])
	} else {
		statFile, err = os.Open("/proc/stat")
	}

	if err != nil {
		log.Fatal(err)
	}

	defer statFile.Close()

	// regular expression for the stat file cpu time line format
	r := regexp.MustCompile(`^(?P<core>cpu[\d]*)\s+(?P<user>\d+)\s+(?P<nice>\d+)\s+(?P<system>\d+)\s+(?P<idle>\d+)\s+(?P<iowait>\d+)\s+(?P<irq>\d+)\s+(?P<softirq>\d+)\s+(?P<steal>\d+)\s+(?P<guest>\d+)\s+(?P<guestNice>\d+)$`)

	scanner := bufio.NewScanner(statFile)

	for scanner.Scan() {

		// if didn't match the regex, abort execution
		if match := r.FindStringSubmatch(scanner.Text()); len(match) > 0 {
			cpuCore := match[1]
			cpuUser, _ := strconv.Atoi(match[2])
			cpuNice, _ := strconv.Atoi(match[3])
			cpuSystem, _ := strconv.Atoi(match[4])
			cpuIdle, _ := strconv.Atoi(match[5])
			cpuIOWait, _ := strconv.Atoi(match[6])
			cpuIRQ, _ := strconv.Atoi(match[7])
			cpuSoftIRQ, _ := strconv.Atoi(match[8])
			cpuSteal, _ := strconv.Atoi(match[9])
			cpuGuest, _ := strconv.Atoi(match[10])
			cpuGuestNice, _ := strconv.Atoi(match[11])
			cpuTotal := cpuUser + cpuNice + cpuSystem + cpuIdle + cpuIOWait + cpuIRQ + cpuSoftIRQ + cpuSteal + cpuGuest + cpuGuestNice

			if cpuCore == "cpu" {
				fmt.Print("%Cpu(s): ")
			} else {
				fmt.Printf("%%%v: ", cpuCore)
			}

			fmt.Printf("%.1f us, ", float32(cpuUser)/float32(cpuTotal)*100)
			fmt.Printf("%.1f sy, ", float32(cpuSystem)/float32(cpuTotal)*100)
			fmt.Printf("%.1f ni, ", float32(cpuNice)/float32(cpuTotal)*100)
			fmt.Printf("%.1f id, ", float32(cpuIdle)/float32(cpuTotal)*100)
			fmt.Printf("%.1f wa, ", float32(cpuIOWait)/float32(cpuTotal)*100)
			fmt.Printf("%.1f hi, ", float32(cpuIRQ)/float32(cpuTotal)*100)
			fmt.Printf("%.1f si, ", float32(cpuSoftIRQ)/float32(cpuTotal)*100)
			fmt.Printf("%.1f st\n", float32(cpuSteal)/float32(cpuTotal)*100)

		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
