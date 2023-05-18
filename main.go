package main

import (
	"dummyApi/api"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Api struct{}

func (a *Api) ValidationFixtureEndpointCustomerDelete(context *gin.Context, parameters *api.ValidationFixtureEndpointCustomerDeleteParameters, body *api.ValidationFixtureEndpointCustomerDeleteRequestBody) api.ValidationFixtureEndpointCustomerDeleteResponse {
	fmt.Printf("Got DELETE\nBODY   :: %#v\nPARAMS :: %#v\n", *body, *parameters)
	context.JSON(http.StatusNoContent, nil)
	return nil
}

func (a *Api) ValidationFixtureEndpointCustomerGet(context *gin.Context, parameters *api.ValidationFixtureEndpointCustomerGetParameters, body *api.ValidationFixtureEndpointCustomerGetRequestBody) api.ValidationFixtureEndpointCustomerGetResponse {
	fmt.Printf("Got GET\nBODY   :: %#v\nPARAMS :: %#v\n", *body, *parameters)
	context.JSON(http.StatusOK, nil)
	ret := api.ValidationFixtureEndpointCustomerGet200Response{}
	ret = append(ret, api.ValidationFixtureEndpointCustomerGet200ResponseModelItem{
		Id:   345,
		Name: "56789",
		Tag:  "5678",
	})
	return ret
}

func (a *Api) ValidationFixtureEndpointCustomerPatch(context *gin.Context, parameters *api.ValidationFixtureEndpointCustomerPatchParameters, body *api.ValidationFixtureEndpointCustomerPatchRequestBody) api.ValidationFixtureEndpointCustomerPatchResponse {
	fmt.Printf("Got PATCH\nBODY   :: %#v\nPARAMS :: %#v\n", *body, *parameters)
	context.JSON(http.StatusOK, nil)
	return nil
}

func (a *Api) ValidationFixtureEndpointCustomerPost(context *gin.Context, parameters *api.ValidationFixtureEndpointCustomerPostParameters, body *api.ValidationFixtureEndpointCustomerPostRequestBody) api.ValidationFixtureEndpointCustomerPostResponse {
	fmt.Printf("Got POST\nBODY   :: %#v\nPARAMS :: %#v\n", *body, *parameters)
	//resp
	//context.JSON(http.StatusOK, nil)
	return &api.ValidationFixtureEndpointCustomerPost405Response{}
}

func (a *Api) InvalidRequest(ctx *gin.Context, err error) {
	ctx.AbortWithStatus(http.StatusMethodNotAllowed)
}

func (a *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	response := `{"Hello":"World"}`
	written, _ := w.Write([]byte(response))
	fmt.Printf("Wrote %d bytes :: %s\n", written, response)
}

// func main_old() {
// 	port := "88"
// 	fmt.Printf("Listening on port %s\n", port)
// 	err := http.ListenAndServe(":"+port, &Api{})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		os.Exit(1)
// 	}
// }

func main() {
	port := ":87"
	r := gin.Default()
	api.RegisterAPI(r, &Api{})
	err := r.Run(port)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

var _ api.API = &Api{}
