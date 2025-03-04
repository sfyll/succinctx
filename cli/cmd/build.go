package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the circuit and generate artifacts",
	Run: func(cmd *cobra.Command, args []string) {
		buildCLI()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func buildCLI() {
	// Check for existence of initialized project
	if !isProjectInitialized() {
		fmt.Println("Project not initialized. Please run 'succinct init' first.")
		return
	}

	// Build and run the generated main.go file
	if err := buildCircuit(); err != nil {
		fmt.Printf("Failed to build and run the project: %v\n", err)
		return
	}

	fmt.Println("Circuit built and artifacts have been successfully generated.")
}

func isProjectInitialized() bool {
	// Check for specific files or directories that indicate the project is initialized
	if _, err := os.Stat("circuit/main.go"); os.IsNotExist(err) {
		return false
	}
	return true
}

func buildCircuit() error {
	args := []string{"run", "./circuit"}
	buildCmd := exec.Command("go", args...)
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		return fmt.Errorf("failed to build the circuit: %w", err)
	}

	return nil
}
