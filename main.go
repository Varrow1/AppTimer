package main

import (
	"fmt"
	"log"
	"os/exec"
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

	pid, err := launchApp(appName)
	if err != nil {
		log.Fatalf("Failed to launch %s: %v\n", appName, err)
	}

	// Start the countdown
	for i := countdown; i > 0; i-- {
		fmt.Printf("Time remaining: %d minutes\n", i)
		time.Sleep(1 * time.Minute) // Wait for 1 Minute
	}

	fmt.Println("Time's up! Killing", appName)
	killAppByPID(pid)
}

func launchApp(appName string) (int, error) {
	cmd := exec.Command(appName)
	err := cmd.Start()
	if err != nil {
		return 0, fmt.Errorf("failed to start the application: %w", err)
	}

	pid := cmd.Process.Pid
	fmt.Printf("%s launched with PID %d\n", appName, pid)
	return pid, nil
}

func killAppByPID(pid int) error {
	cmd := exec.Command("kill", fmt.Sprintf("%d", pid))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to kill process with PID %d: %w", pid, err)
	}

	fmt.Printf("Process with PID %d has been killed.\n", pid)
	return nil
}
