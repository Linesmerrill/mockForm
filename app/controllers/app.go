package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"encoding/json"
	"mockForm/app"
	"reflect"
	// "database/sql"
	// _ "github.com/lib/pq"
)

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
	c.everythingFromDB()
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

func (c App) everythingFromDB(){
	rows, err := app.DB.Query("select * from users");
	if err != nil {
		fmt.Println("Beluga Whales: ", err);
	}

	defer rows.Close()
  for rows.Next() {
          var name string
					var id int
					var email string
          if err := rows.Scan(&id, &name, &email); err != nil {
                  fmt.Println(err)
          }
          fmt.Printf("Name: %s, Email: %s", name, email)
  }
  if err := rows.Err(); err != nil {
          fmt.Println(err)
  }
	fmt.Println("Narwhal: ", reflect.TypeOf(rows))
}
