package main

import "fmt"

func main() {
	// *** Config parser + watcher ***

	// Generate a list of PetsFiles from the given config directory.
	fmt.Println("DEBUG: * configuration parsing starts *")

	files, err := ParseFiles("sample_pet")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("DEBUG: * configuration parsing ends *")

	// *** Config validator ***
	fmt.Println("DEBUG: * configuration validation starts *")
	globalErrors := CheckGlobalConstraints(files)

	if globalErrors != nil {
		fmt.Println(err)
		// Global validation errors mean we should stop the whole update.
		return
	}

	// Check validation errors in individual files. At this stage, the
	// command in the "pre" validation directive may not be installed yet.
	// Ignore PathErrors for now. Get a list of valid files.
	goodPets := CheckLocalConstraints(files, true)

	fmt.Println("DEBUG: * configuration validation ends *")

	// *** Update visualizer ***
	// Display:
	// - packages to install (XXX: instead of installing all packages one-by-one, we should get a list of package names and then apt install A B C)
	// - pre-update command output
	// - files created/modified
	// - content diff
	// - permissions/owner changes
	// - which post-update commands will be executed
	for _, pet := range goodPets {
		for _, action := range NewPetsActions(pet) {
			action.Visualize()
		}
	}

	// *** Update executor ***
	// Install missing packages
	// Create missing directories
	// Run pre-update command and stop the update if it fails
	// Update files
	// Change permissions/owners
	// Run post-update commands
}
