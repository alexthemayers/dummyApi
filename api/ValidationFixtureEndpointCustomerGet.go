// GENERATED CODE. DO NOT EDIT

package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var _ = fmt.Sprint("")

type ValidationFixtureEndpointCustomerGetParameters struct {
	Customer string `json:"customer"`
	Id       string `json:"id"`
}

func (p *ValidationFixtureEndpointCustomerGetParameters) IsValid() (bool, error) {

	return true, nil
}

type ValidationFixtureEndpointCustomerGetRequestBody struct{}

func (*ValidationFixtureEndpointCustomerGetRequestBody) IsValid() (bool, error) {
	return true, nil
}

type ValidationFixtureEndpointCustomerGetResponse interface {
	Send(c *gin.Context)
	GetStatus() int
}

type ValidationFixtureEndpointCustomerGet200Response ValidationFixtureEndpointCustomerGet200ResponseModel

func (r ValidationFixtureEndpointCustomerGet200Response) GetStatus() int {
	return 200
}

func (r ValidationFixtureEndpointCustomerGet200Response) Send(c *gin.Context) {
	c.JSON(200, r)
}

type ValidationFixtureEndpointCustomerGet405Response struct{}

func (r ValidationFixtureEndpointCustomerGet405Response) GetStatus() int {
	return 405
}
func (r ValidationFixtureEndpointCustomerGet405Response) Send(c *gin.Context) {
	c.JSON(405, r)
}
