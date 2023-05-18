// GENERATED CODE. DO NOT EDIT

package api

import "github.com/gin-gonic/gin"

type API interface {
	ValidationFixtureEndpointCustomer
}

func RegisterAPI(e gin.IRouter, srv API) {
	RegisterValidationFixtureEndpointCustomerPath(e, srv)
}

type UnimplementedServer struct {
	UnimplementedValidationFixtureEndpointCustomer
}

func (u UnimplementedServer) InvalidRequest(c *gin.Context, err error) {
	c.JSON(400, err.Error())
	c.Abort()
}
