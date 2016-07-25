package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"net/http"
	// "html"
	// "encoding/json"
)

type App struct {
	*revel.Controller
}

type Surface struct {
    Data string `json:"data"` //since we have to export the field
                              //but want the lowercase letter
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Verify(req *http.Request) revel.Result {
	
	fmt.Println("THIS IS THE PARAMS", c.Params.Form.Get("serviceName"))
	return c.Render()
}

func (c App) RouteToVerifyPage() revel.Result {
	fmt.Println(c)
	return c.Redirect(App.Verify)
}
