package requirements

type requirementFunc func() error

var requirementFuncs = []requirementFunc{}

func Register(f requirementFunc) {
	requirementFuncs = append(requirementFuncs, f)
}

func Check() []error {
	requirementErrs := []error{}
	for _, f := range requirementFuncs {
		if err := f(); err != nil {
			requirementErrs = append(requirementErrs, err)
		}
	}

	return requirementErrs
}
