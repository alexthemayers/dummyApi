// GENERATED CODE. DO NOT EDIT

package api

import (
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
)

var _ = fmt.Sprint("")

type ValidationFixtureEndpointCustomerPostParameters struct {
	Customer string `json:"customer"`
	Id       string `json:"id"`
}

func (p *ValidationFixtureEndpointCustomerPostParameters) IsValid() (bool, error) {

	return true, nil
}

type ValidationFixtureEndpointCustomerPostRequestBody ValidationFixtureEndpointCustomerPostRequestBodyModel

func (p *ValidationFixtureEndpointCustomerPostRequestBody) IsValid() (bool, error) {
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

type ValidationFixtureEndpointCustomerPostResponse interface {
	Send(c *gin.Context)
	GetStatus() int
}

type ValidationFixtureEndpointCustomerPost201Response struct{}

func (*ValidationFixtureEndpointCustomerPost201Response) IsValid() (bool, error) {
	return true, nil
}

func (r ValidationFixtureEndpointCustomerPost201Response) GetStatus() int {
	return 201
}

func (r ValidationFixtureEndpointCustomerPost201Response) Send(c *gin.Context) {
	c.String(201, fmt.Sprint(r))
}

type ValidationFixtureEndpointCustomerPost405Response struct{}

func (r ValidationFixtureEndpointCustomerPost405Response) GetStatus() int {
	return 405
}
func (r ValidationFixtureEndpointCustomerPost405Response) Send(c *gin.Context) {
	c.JSON(405, r)
}
