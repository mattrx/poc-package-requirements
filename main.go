package main

import (
	"log"
	"os"

	"github.com/mattrx/poc-package-requirements/internal/packageA"
	"github.com/mattrx/poc-package-requirements/internal/packageB"
	"github.com/mattrx/poc-package-requirements/internal/packageC"
	"github.com/mattrx/poc-package-requirements/internal/requirements"
)

func main() {
	log.Println("Starting...")

	setup()

	errs := requirements.Check()
	if len(errs) > 0 {
		for _, err := range errs {
			log.Printf("Requirement not met: %v", err)
		}
		os.Exit(1)
	}

	run()

	log.Println("Done.")
}

func setup() {
	packageA.ConfigurePackage()
	packageB.ConfigurePackage()
	packageC.ConfigurePackage()
}

func run() {
	packageA.Do()
}
