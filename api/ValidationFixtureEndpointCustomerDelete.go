// GENERATED CODE. DO NOT EDIT

package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var _ = fmt.Sprint("")

type ValidationFixtureEndpointCustomerDeleteParameters struct {
	Customer string `json:"customer"`
	Id       string `json:"id"`
}

func (p *ValidationFixtureEndpointCustomerDeleteParameters) IsValid() (bool, error) {

	return true, nil
}

type ValidationFixtureEndpointCustomerDeleteRequestBody struct{}

func (*ValidationFixtureEndpointCustomerDeleteRequestBody) IsValid() (bool, error) {
	return true, nil
}

type ValidationFixtureEndpointCustomerDeleteResponse interface {
	Send(c *gin.Context)
	GetStatus() int
}

type ValidationFixtureEndpointCustomerDelete204Response struct{}

func (*ValidationFixtureEndpointCustomerDelete204Response) IsValid() (bool, error) {
	return true, nil
}

func (r ValidationFixtureEndpointCustomerDelete204Response) GetStatus() int {
	return 204
}

func (r ValidationFixtureEndpointCustomerDelete204Response) Send(c *gin.Context) {
	c.String(204, fmt.Sprint(r))
}

type ValidationFixtureEndpointCustomerDelete405Response struct{}

func (r ValidationFixtureEndpointCustomerDelete405Response) GetStatus() int {
	return 405
}
func (r ValidationFixtureEndpointCustomerDelete405Response) Send(c *gin.Context) {
	c.JSON(405, r)
}
