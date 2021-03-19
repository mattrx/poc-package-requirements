package packageA

import (
	"errors"
	"log"

	"github.com/mattrx/poc-package-requirements/internal/packageB"
	"github.com/mattrx/poc-package-requirements/internal/requirements"
)

func init() {
	requirements.Register(func() error {
		if !configured {
			return errors.New("packageA not configured")
		}

		return nil
	})
}

var configured = false

func ConfigurePackage() {
	log.Println("packageA configured")
	configured = true
}

func Do() {
	log.Println("packageA does something")
	packageB.Do()
}
