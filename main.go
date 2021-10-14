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

	if len(os.Args) > 1 {
		statFile, err = os.Open(os.Args[1])
	} else {
		statFile, err = os.Open("/proc/stat")
	}

	if err != nil {
		log.Fatal(err)
	}

	defer statFile.Close()

	r := regexp.MustCompile(`^cpu[\d]*\s+(?P<user>\d+)\s+(?P<nice>\d+)\s+(?P<system>\d+)\s+(?P<idle>\d+)\s+(?P<iowait>\d+)\s+(?P<irq>\d+)\s+(?P<softirq>\d+)\s+(?P<steal>\d+)\s+(?P<guest>\d+)\s+(?P<guestNice>\d+)$`)

	scanner := bufio.NewScanner(statFile)

	for scanner.Scan() {
		match := r.FindStringSubmatch(scanner.Text())

		cpuUser, _ := strconv.Atoi(match[1])
		cpuNice, _ := strconv.Atoi(match[2])
		cpuSystem, _ := strconv.Atoi(match[3])
		cpuIdle, _ := strconv.Atoi(match[4])
		cpuIOWait, _ := strconv.Atoi(match[5])
		cpuIRQ, _ := strconv.Atoi(match[6])
		cpuSoftIRQ, _ := strconv.Atoi(match[7])
		cpuSteal, _ := strconv.Atoi(match[8])
		cpuGuest, _ := strconv.Atoi(match[9])
		cpuGuestNice, _ := strconv.Atoi(match[10])
		cpuTotal := cpuUser + cpuNice + cpuSystem + cpuIdle + cpuIOWait + cpuIRQ + cpuSoftIRQ + cpuSteal + cpuGuest + cpuGuestNice

		fmt.Print("%Cpu(s): ")
		fmt.Printf("%.1f us, ", float32(cpuUser)/float32(cpuTotal)*100)
		fmt.Printf("%.1f sy, ", float32(cpuSystem)/float32(cpuTotal)*100)
		fmt.Printf("%.1f ni, ", float32(cpuNice)/float32(cpuTotal)*100)
		fmt.Printf("%.1f id, ", float32(cpuIdle)/float32(cpuTotal)*100)
		fmt.Printf("%.1f wa, ", float32(cpuIOWait)/float32(cpuTotal)*100)
		fmt.Printf("%.1f hi, ", float32(cpuIRQ)/float32(cpuTotal)*100)
		fmt.Printf("%.1f si, ", float32(cpuSoftIRQ)/float32(cpuTotal)*100)
		fmt.Printf("%.1f st\n", float32(cpuSteal)/float32(cpuTotal)*100)
		//fmt.Printf("guest: %f%% ", float32(cpuGuest)/float32(cpuTotal))
		//fmt.Printf("guestNice: %f%%\n", float32(cpuGuestNice)/float32(cpuTotal))
		break
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
