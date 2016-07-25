package controllers

import (
	"github.com/revel/revel"
	"fmt"
	// "reflect"
	// "net/http"
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


func (c App) Verify() revel.Result {
	params := c.Params.Form

	for k := range c.Params.Form {
		if(c.VerifyUserInput(params[k][0])){
			return c.Redirect("/")
		}
	}
	return c.Render(params)
}

func (c App) Blah() revel.Result {
	fmt.Println("AHHHHHHHHH STUFFFFFF", c.Params)
	return c.Render()
}

func (c App) VerifyUserInput(form string) bool{
		c.Validation.Required(form)
		if(c.Validation.HasErrors()){
			return true
		}
		return false
}
