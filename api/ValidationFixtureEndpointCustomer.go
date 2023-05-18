package api

import (
	"github.com/gin-gonic/gin"
)

// GENERATED INTERFACE. DO NOT EDIT

type ValidationFixtureEndpointCustomer interface {
	ValidationFixtureEndpointCustomerDelete(*gin.Context, *ValidationFixtureEndpointCustomerDeleteParameters, *ValidationFixtureEndpointCustomerDeleteRequestBody) ValidationFixtureEndpointCustomerDeleteResponse
	ValidationFixtureEndpointCustomerGet(*gin.Context, *ValidationFixtureEndpointCustomerGetParameters, *ValidationFixtureEndpointCustomerGetRequestBody) ValidationFixtureEndpointCustomerGetResponse
	ValidationFixtureEndpointCustomerPatch(*gin.Context, *ValidationFixtureEndpointCustomerPatchParameters, *ValidationFixtureEndpointCustomerPatchRequestBody) ValidationFixtureEndpointCustomerPatchResponse
	ValidationFixtureEndpointCustomerPost(*gin.Context, *ValidationFixtureEndpointCustomerPostParameters, *ValidationFixtureEndpointCustomerPostRequestBody) ValidationFixtureEndpointCustomerPostResponse
	InvalidRequest(*gin.Context, error)
}

type UnimplementedValidationFixtureEndpointCustomer struct{}

func (u UnimplementedValidationFixtureEndpointCustomer) ValidationFixtureEndpointCustomerDelete(*gin.Context, *ValidationFixtureEndpointCustomerDeleteParameters, *ValidationFixtureEndpointCustomerDeleteRequestBody) ValidationFixtureEndpointCustomerDeleteResponse {
	return ValidationFixtureEndpointCustomerDelete405Response{}
}
func (u UnimplementedValidationFixtureEndpointCustomer) ValidationFixtureEndpointCustomerGet(*gin.Context, *ValidationFixtureEndpointCustomerGetParameters, *ValidationFixtureEndpointCustomerGetRequestBody) ValidationFixtureEndpointCustomerGetResponse {
	return ValidationFixtureEndpointCustomerGet405Response{}
}
func (u UnimplementedValidationFixtureEndpointCustomer) ValidationFixtureEndpointCustomerPatch(*gin.Context, *ValidationFixtureEndpointCustomerPatchParameters, *ValidationFixtureEndpointCustomerPatchRequestBody) ValidationFixtureEndpointCustomerPatchResponse {
	return ValidationFixtureEndpointCustomerPatch405Response{}
}
func (u UnimplementedValidationFixtureEndpointCustomer) ValidationFixtureEndpointCustomerPost(*gin.Context, *ValidationFixtureEndpointCustomerPostParameters, *ValidationFixtureEndpointCustomerPostRequestBody) ValidationFixtureEndpointCustomerPostResponse {
	return ValidationFixtureEndpointCustomerPost405Response{}
}
func (u UnimplementedValidationFixtureEndpointCustomer) InvalidRequest(c *gin.Context, err error) {
	c.JSON(400, err.Error())
	c.Abort()
}

func RegisterValidationFixtureEndpointCustomerPath(e gin.IRouter, srv ValidationFixtureEndpointCustomer) {

	e.DELETE("/endpoint/{customer}", func(c *gin.Context) {
		params := &ValidationFixtureEndpointCustomerDeleteParameters{}
		params.Customer = c.Param("customer")
		params.Id = c.Query("id")

		if valid, err := params.IsValid(); !valid {
			srv.InvalidRequest(c, err)
			return
		}

		var body *ValidationFixtureEndpointCustomerDeleteRequestBody

		if valid, err := body.IsValid(); !valid {
			srv.InvalidRequest(c, err)
			return
		}
		response := srv.ValidationFixtureEndpointCustomerDelete(c, params, body)
		response.Send(c)
	})

	e.GET("/endpoint/{customer}", func(c *gin.Context) {
		params := &ValidationFixtureEndpointCustomerGetParameters{}
		params.Customer = c.Param("customer")
		params.Id = c.Query("id")

		if valid, err := params.IsValid(); !valid {
			srv.InvalidRequest(c, err)
			return
		}

		var body *ValidationFixtureEndpointCustomerGetRequestBody

		if valid, err := body.IsValid(); !valid {
			srv.InvalidRequest(c, err)
			return
		}
		response := srv.ValidationFixtureEndpointCustomerGet(c, params, body)
		response.Send(c)
	})

	e.PATCH("/endpoint/{customer}", func(c *gin.Context) {
		params := &ValidationFixtureEndpointCustomerPatchParameters{}
		params.Customer = c.Param("customer")
		params.Id = c.Query("id")

		if valid, err := params.IsValid(); !valid {
			srv.InvalidRequest(c, err)
			return
		}

		var body *ValidationFixtureEndpointCustomerPatchRequestBody

		if valid, err := body.IsValid(); !valid {
			srv.InvalidRequest(c, err)
			return
		}
		response := srv.ValidationFixtureEndpointCustomerPatch(c, params, body)
		response.Send(c)
	})

	e.POST("/endpoint/{customer}", func(c *gin.Context) {
		params := &ValidationFixtureEndpointCustomerPostParameters{}
		params.Customer = c.Param("customer")
		params.Id = c.Query("id")

		if valid, err := params.IsValid(); !valid {
			srv.InvalidRequest(c, err)
			return
		}

		var body *ValidationFixtureEndpointCustomerPostRequestBody
		body = &ValidationFixtureEndpointCustomerPostRequestBody{}
		err := c.ShouldBindJSON(body)
		if err != nil {
			srv.InvalidRequest(c, err)
			return
		}

		if valid, err := body.IsValid(); !valid {
			srv.InvalidRequest(c, err)
			return
		}
		response := srv.ValidationFixtureEndpointCustomerPost(c, params, body)
		response.Send(c)
	})
}
