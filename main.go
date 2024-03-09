package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	script := "./runCreateBlog.sh"
	cmd := exec.Command("bash", script)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to start command: %s", err)
	}

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("Command finished with error: %s", err)
	}

}
