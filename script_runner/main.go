// main.go
package main

import (
	"flag"
	"fmt"
	"os"
)

// ScriptRegistry maps script names to their Run functions
var ScriptRegistry = map[string]func() error{
	// "pkg1/script1": pkg1.Script1Run,
}

func main() {
	scriptName := flag.String("script", "", "Script to run (e.g., pkg1/script1)")
	flag.Parse()

	if *scriptName == "" {
		fmt.Println("❌ Error: --script flag is required")
		printAvailableScripts()
		os.Exit(1)
	}

	scriptFunc, exists := ScriptRegistry[*scriptName]
	if !exists {
		fmt.Printf("❌ Error: Unknown script '%s'\n", *scriptName)
		printAvailableScripts()
		os.Exit(1)
	}

	fmt.Printf("🚀 Running script: %s\n", *scriptName)
	fmt.Println("─────────────────────────────────────────")

	if err := scriptFunc(); err != nil {
		fmt.Println("─────────────────────────────────────────")
		fmt.Printf("❌ Script failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("─────────────────────────────────────────")
	fmt.Println("✅ Script completed successfully")
}

func printAvailableScripts() {
	fmt.Println("\n📋 Available scripts:")
	for name := range ScriptRegistry {
		fmt.Printf("  • %s\n", name)
	}
	fmt.Println("\nUsage:")
	fmt.Println("  go run main.go --script=pkg1/script1")
}
