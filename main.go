package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <parent1.sh> <parent2.sh> [--children child1.sh child2.sh ...]")
		os.Exit(1)
	}

	// Split arguments into parents and children
	var parentScripts, childScripts []string
	isChildren := false
	for _, arg := range os.Args[1:] {
		if arg == "--children" {
			isChildren = true
			continue
		}

		if isChildren {
			childScripts = append(childScripts, arg)
		} else {
			parentScripts = append(parentScripts, arg)
		}
	}

	if len(parentScripts) == 0 {
		fmt.Println("Error: At least one parent script is required")
		os.Exit(1)
	}

	// Validate all scripts exist before starting execution
	for _, script := range append(parentScripts, childScripts...) {
		absPath, err := filepath.Abs(script)
		if err != nil {
			fmt.Printf("Failed to get absolute path for '%s': %v\n", script, err)
			os.Exit(1)
		}

		// Check if file exists and is readable
		if _, err := os.Stat(absPath); err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("Script file does not exist: %s\n", absPath)
			} else {
				fmt.Printf("Cannot access script file '%s': %v\n", absPath, err)
			}

			os.Exit(1)
		}
	}

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	errorChan := make(chan error, 1)

	// First run all child scripts in parallel
	for _, script := range childScripts {
		wg.Add(1)
		go func(script string) {
			defer wg.Done()

			absPath, err := filepath.Abs(script)
			if err != nil {
				fmt.Printf("Failed to get absolute path for '%s': %v\n", script, err)
				cancel()
				select {
				case errorChan <- err:
				default:
				}
				return
			}

			cmd := exec.CommandContext(ctx, "bash", absPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Dir = filepath.Dir(absPath)

			err = cmd.Run()
			if err != nil {
				fmt.Printf("Child script '%s' failed with error: %v\n", script, err)
				cancel()
				select {
				case errorChan <- err:
				default:
				}
			}
		}(script)
	}

	// Wait for all child scripts to complete
	childrenDone := make(chan struct{})
	go func() {
		wg.Wait()
		close(childrenDone)
	}()

	// Wait for children to complete or error
	select {
	case <-childrenDone:
		fmt.Println("All child scripts completed successfully")
	case err := <-errorChan:
		fmt.Printf("Exiting due to error in child scripts: %v\n", err)
		os.Exit(1)
	}

	// Now run parent scripts sequentially
	for _, script := range parentScripts {
		absPath, err := filepath.Abs(script)
		if err != nil {
			fmt.Printf("Failed to get absolute path for '%s': %v\n", script, err)
			os.Exit(1)
		}

		cmd := exec.Command("bash", absPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = filepath.Dir(absPath)

		err = cmd.Run()
		if err != nil {
			fmt.Printf("Parent script '%s' failed with error: %v\n", script, err)
			os.Exit(1)
		}
	}

	fmt.Println("All scripts completed successfully")
}
