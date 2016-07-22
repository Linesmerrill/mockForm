package controllers

import (
	"github.com/revel/revel"
	"fmt"
)

type App struct {
	*revel.Controller
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
	return c.Render()
}

func (c App) RouteToVerifyPage() revel.Result {
	fmt.Println(c)
	return c.Redirect(App.Verify)
}
