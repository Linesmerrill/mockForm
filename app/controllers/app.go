package controllers

import (
	"github.com/revel/revel"
	"fmt"
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
	// fmt.Println(form)
	// ParseParams(c.Params, c.Request)
	// url.Parse(c.Params)
	// for k,v := range MAPP{
	// 	fmt.Println("Key:", k,"Value:", v)
	// }

	// fmt.Println("THIS IS THE REQUEST", c.Request)
	fmt.Println("THIS IS THE PARAMS", c.Params)
	c.FlashParams()
	params := c.Params.Form.Get("addInstruct")
	return c.Render(params)
}

func (c App) Blah() revel.Result {
	fmt.Println("AHHHHHHHHH STUFFFFFF", c.Params)
	return c.Render()
}
func (c App) RouteToVerifyPage() revel.Result {
	fmt.Println(c)
	return c.Redirect(App.Blah)
}
