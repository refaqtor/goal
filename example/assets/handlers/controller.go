// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	c0 "github.com/anonx/sunplate/example/assets/handlers/github.com/anonx/sunplate/middleware/template"
	contr "github.com/anonx/sunplate/example/controllers"

	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/strconv"
)

// Controller is automatically generated from a controller
// that was found at "github.com/anonx/sunplate/example/controllers/init.go".
//
// Controller is a struct that should be embedded into every controller
// of your app to make methods provided by middlewares available.
type Controller struct {
}

// New allocates (github.com/anonx/sunplate/example/controllers).Controller controller,
// initializes its parents; then returns the controller.
func (t Controller) New() *contr.Controller {
	c := &contr.Controller{}
	c.Middleware = c0.Middleware{}.New()
	return c
}

// Before executes magic actions of embedded controllers, and
// calls (github.com/anonx/sunplate/example/controllers).Controller.Before with arguments
// that are extracted from r.Form and converted to appropriate types.
func (t Controller) Before(c *contr.Controller, w http.ResponseWriter, r *http.Request) a.Result {
	// Execute magic Before actions of embedded controllers.
	if r := (c0.Middleware{}.After(c.Middleware, w, r)); r != nil {
		return r
	}
	// Call magic Before action of (github.com/anonx/sunplate/example/controllers).Controller.
	if r := c.Before( // Parameters should be binded.
	); r != nil {
		return r
	}
	// Continue execution chain.
	return nil
}

// After executes magic actions of embedded controllers.
func (t Controller) After(c *contr.Controller, w http.ResponseWriter, r *http.Request) a.Result {
	// Execute magic After actions of embedded controllers.
	if r := (c0.Middleware{}.After(c.Middleware, w, r)); r != nil {
		return r
	}
	// Continue execution chain.
	return nil
}

func init() {
	_ = strconv.MeaningOfLife
}
