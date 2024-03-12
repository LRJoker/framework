package generator

import "framework/generator/gen"

const (
	OpMakeController = "make:controller"
	OpMakeModel      = "make:model"
	OpMakeTransform  = "make:transform"
	OpMakeRoute      = "make:route"
	OpMakeMiddleware = "make:middleware"
	OpMakeMigrate    = "make:migrate"
	OpMakeValidate   = "make:validate"
)

func parse(args []string) error {

	op := args[0]

	fileName := args[1]

	switch op {
	case OpMakeController:
		Make(&gen.Controller{Name: fileName})
	case OpMakeModel:
		Make(&gen.Model{Name: fileName})
	case OpMakeTransform:
		Make(&gen.Transform{Name: fileName})
	case OpMakeRoute:
		Make(&gen.Transform{Name: fileName})
	case OpMakeMiddleware:
		Make(&gen.Transform{Name: fileName})
	case OpMakeMigrate:
		Make(&gen.Transform{Name: fileName})
	case OpMakeValidate:
		Make(&gen.Validate{Name: fileName})
	}

	return nil
}
