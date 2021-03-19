package packageB

import (
	"errors"
	"log"

	"github.com/mattrx/poc-package-requirements/internal/packageC"
	"github.com/mattrx/poc-package-requirements/internal/requirements"
)

func init() {
	requirements.Register(func() error {
		if !configured {
			return errors.New("packageB not configured")
		}

		return nil
	})
}

var configured = false

func ConfigurePackage() {
	log.Println("packageB configured")
	configured = true
}

func Do() {
	log.Println("packageB does something")
	packageC.Do()
}
