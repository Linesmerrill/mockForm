package app

import (
	"github.com/revel/revel"
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
)
var DB *sql.DB

func InitDB() {
    connstring := fmt.Sprintf("user=%s dbname=%s sslmode=disable", "localadmin", "mockDB")

    var err error
    DB, err = sql.Open("postgres", connstring)
    if err != nil {
        revel.INFO.Println("DB Error", err)
    }
		err = DB.Ping();
		if err != nil {
			fmt.Println(err);
		}
		// fmt.Println(err);
    revel.INFO.Println("DB Connected")

		// DB.Query("CREATE TABLE users (id serial, name text, email text)")
		// DB.Query("INSERT INTO users (name, email) VALUES ('Bob', 'Bob@fake.net')")
		// row, err := DB.Query("SELECT * FROM users")
		// revel.INFO.Println(row);
}

func init() {
	fmt.Println("HERE IN init")
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
