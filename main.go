// main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	scriptPath := flag.String("script", "", "Script to run (e.g., pkg1/script1.go)")
	flag.Parse()

	if *scriptPath == "" {
		fmt.Println("❌ Error: --script flag is required")
		fmt.Println("\nUsage:")
		fmt.Println("  go run main.go --script=pkg1/script1.go")
		os.Exit(1)
	}

	// Ensure .go extension
	if !strings.HasSuffix(*scriptPath, ".go") {
		*scriptPath = *scriptPath + ".go"
	}

	// Check if file exists
	if _, err := os.Stat(*scriptPath); os.IsNotExist(err) {
		fmt.Printf("❌ Error: Script file not found: %s\n", *scriptPath)
		os.Exit(1)
	}

	// Get absolute path
	absPath, err := filepath.Abs(*scriptPath)
	if err != nil {
		fmt.Printf("❌ Error: Failed to resolve path: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("🚀 Running script: %s\n", *scriptPath)
	fmt.Println("─────────────────────────────────────────")

	// Execute: go run <script_path> Run
	// The "Run" argument tells the script to execute its Run() function
	cmd := exec.Command("go", "run", absPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		fmt.Println("─────────────────────────────────────────")
		fmt.Printf("❌ Script failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("─────────────────────────────────────────")
	fmt.Println("✅ Script completed successfully")
}
