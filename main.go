package main

import (
	"log"
	"os"

	"github.com/mattrx/poc-package-requirements/internal/app"
	"github.com/mattrx/poc-package-requirements/internal/requirements"
)

func main() {
	handler := app.Handler{}

	errs := requirements.Check()
	if len(errs) > 0 {
		for _, err := range errs {
			log.Printf("Requirement not met: %v", err)
		}
		os.Exit(1)
	}

	handler.Do()
}
