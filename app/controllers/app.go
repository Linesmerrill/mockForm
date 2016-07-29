package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"encoding/json"
	// "database/sql"
)
//
// type DB struct{
// 	*app.DB
// }

type App struct {
	*revel.Controller
}

type Form struct {
	ServiceName string `json:"serviceName"`
	BackendServeImp string `json:"backendServeImp"`
	BackendServePubCert string `json:"backendServePubCert"`
}

func (c App) Index() revel.Result {

	// database := c.everythingFromDB()
	return c.Render()
}


func (c App) Verify() revel.Result {
	params := c.Params.Form

	for k := range c.Params.Form {
		if(c.VerifyUserInput(params[k][0])){
			return c.Redirect("/")
		}
	}
	data, _ := json.Marshal(c.Params.Form)
	fmt.Println(string(data))
	return c.Render(params)
}

func (c App) VerifyUserInput(form string) bool{
		c.Validation.Required(form)
		if(c.Validation.HasErrors()){
			return true
		}
		return false
}

// func (c App) everythingFromDB() string{
// 	return DB.Query("select * from users");
// }
