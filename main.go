package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome To Varrow's Cli screen time limiter")

	// Set the countdown duration
	var countdown int
	fmt.Println("input time you want to spent")
	fmt.Scanln(&countdown)
	var appName string
	fmt.Println("Input the app you want to spend time on")
	fmt.Scanln(&appName)

	// Start the countdown
	for i := countdown; i > 0; i-- {
		fmt.Printf("Time remaining: %d seconds\n", i)
		time.Sleep(1 * time.Minute) // Wait for 1 Minute
	}

	fmt.Println("Time's up! Killing", appName)
	killApp(appName)
}

func killApp(appName string) {
	// Find the process ID (PID) of the application
	cmd := exec.Command("pgrep", appName)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to find %s process: %v\n", appName, err)
	}

	// Convert the output to a string and remove any extra whitespace
	pid := strings.TrimSpace(string(output))

	// Kill the process by PID
	cmd = exec.Command("kill", pid)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to kill %s process: %v\n", appName, err)
	}

	fmt.Printf("%s has been killed.\n", appName)
}
