package api

import (
	"fmt"
	"regexp"
)

// GENERATED MODEL. DO NOT EDIT

type ValidationFixtureEndpointCustomerPostRequestBodyModel struct {
	Description            string `json:"description"`
	Enabled                bool   `json:"enabled"`
	GoPropertiesWithDashes string `json:"go-properties-with-dashes"`
	Id                     string `json:"id"`
	Name                   string `json:"name"`
}

func (p *ValidationFixtureEndpointCustomerPostRequestBodyModel) IsValid() (bool, error) {
	if len(p.Name) < 1 {
		return false, fmt.Errorf("p.Name too short")
	}
	if len(p.Name) > 64 {
		return false, fmt.Errorf("p.Name too long")
	}
	// Regular expressions are checked for compilation during code generation
	// No need to check them here.
	rName, _ := regexp.Compile(`^([a-zA-Z0-9])+([-_ @\.]([a-zA-Z0-9])+)*$`)
	if !rName.MatchString(p.Name) {
		return false, fmt.Errorf("p.Name did not pass validation pattern")
	}

	return true, nil
}
