// GENERATED CODE. DO NOT EDIT

package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var _ = fmt.Sprint("")

type ValidationFixtureEndpointCustomerPatchParameters struct {
	Customer string `json:"customer"`
	Id       string `json:"id"`
}

func (p *ValidationFixtureEndpointCustomerPatchParameters) IsValid() (bool, error) {

	return true, nil
}

type ValidationFixtureEndpointCustomerPatchRequestBody struct{}

func (*ValidationFixtureEndpointCustomerPatchRequestBody) IsValid() (bool, error) {
	return true, nil
}

type ValidationFixtureEndpointCustomerPatchResponse interface {
	Send(c *gin.Context)
	GetStatus() int
}

type ValidationFixtureEndpointCustomerPatch204Response struct{}

func (*ValidationFixtureEndpointCustomerPatch204Response) IsValid() (bool, error) {
	return true, nil
}

func (r ValidationFixtureEndpointCustomerPatch204Response) GetStatus() int {
	return 204
}

func (r ValidationFixtureEndpointCustomerPatch204Response) Send(c *gin.Context) {
	c.String(204, fmt.Sprint(r))
}

type ValidationFixtureEndpointCustomerPatch405Response struct{}

func (r ValidationFixtureEndpointCustomerPatch405Response) GetStatus() int {
	return 405
}
func (r ValidationFixtureEndpointCustomerPatch405Response) Send(c *gin.Context) {
	c.JSON(405, r)
}
