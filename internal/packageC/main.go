package packageC

import (
	"errors"
	"log"

	"github.com/mattrx/poc-package-requirements/internal/requirements"
)

func init() {
	requirements.Register(func() error {
		if !configured {
			return errors.New("packageC not configured")
		}

		return nil
	})
}

var configured = false

func ConfigurePackage() {
	log.Println("packageC configured")
	configured = true
}

func Do() {
	log.Println("packageC does something")
}
