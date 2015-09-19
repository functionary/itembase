// Package itembase gives a thin wrapper around the itembase REST API.
package itembase

import (
	"encoding/json"
)

// Error is a Go representation of the error message sent back by itembase when a
// request results in an error.
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (f *Error) Error() string {
	return f.Message
}

// This is the default implementation
type client struct {
	// root is the client's base URL used for all calls.
	root string
	me   string

	// url is the current url to call
	url string

	// auth is authentication token used when making calls.
	// The token is optional and can also be overwritten on an individual
	// call basis via params.
	auth string

	// user is the current shop we're calling for
	user string

	// production environment vs sandbox
	production bool

	// api is the underlying client used to make calls.
	api API

	params  map[string]string
	options Config
}

// New creates a new instance of the default itembase Client implementation.
//
// The options must be non-nil and must provide all OAuth2 credentials and
// configuration for an application registered with the itembase API.
//
// TODO: always use the default API impl, NewClient allows dependency injection
// needed for testing.
func New(options Config, api API) Client {
	if api == nil {
		api = new(itembaseAPI)
	}

	return &client{options: options, production: options.Production, api: api}
}

// NewClient is an alternative Client constructor intended for testing or
// advanced usage, where a custom API implementation can be injected.
func NewClient(root, auth string, api API) Client {
	if api == nil {
		api = new(itembaseAPI)
	}

	return &client{url: root, root: root, auth: auth, api: api}
}

func (c *client) URL() string {
	return c.url
}

func (c *client) Sandbox() Client {
	c.production = false
	return c
}

func (c *client) User(user string) Client {
	c.auth = c.getUserToken(user).AccessToken
	c.user = user
	c.url = c.root + "/users/" + user
	return c
}

func (c *client) GetInto(destination interface{}) error {
	err := c.api.Call("GET", c.url, c.auth, nil, c.params, &destination)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Get() (destination interface{}, err error) {
	err = c.api.Call("GET", c.url, c.auth, nil, c.params, &destination)
	return
}

func (c *client) Me() (destination User, err error) {
	err = c.api.Call("GET", c.me, c.auth, nil, c.params, &destination)
	return
}

func (c *client) Child(path string) Client {
	c.url = c.url + "/" + path
	return c
}

func (c *client) Transactions() Client {
	c.url = c.root + "/users/" + c.user + "/transactions"
	return c
}

func (c *client) Products() Client {
	c.url = c.root + "/users/" + c.user + "/products"
	return c
}

func (c *client) Buyers() Client {
	c.url = c.root + "/users/" + c.user + "/buyers"
	return c
}

func (c *client) Profiles() Client {
	c.url = c.root + "/users/" + c.user + "/profiles"
	return c
}

// These are some shenanigans, golang. Shenanigans I say.
func (c *client) newParamMap(key string, value interface{}) map[string]string {
	ret := make(map[string]string, len(c.params)+1)
	for key, value := range c.params {
		ret[key] = value
	}
	switch value.(type) {
	case string:
		ret[key] = value.(string)
	default:
		jsonVal, _ := json.Marshal(value)
		ret[key] = string(jsonVal)
	}
	return ret
}

func (c *client) clientWithNewParam(key string, value interface{}) *client {
	c.params = c.newParamMap(key, value)
	return c
}

// Query functions.
func (c *client) Select(prop string) Client {
	c.url = c.url + "/" + prop
	return c
}

func (c *client) CreatedAtFrom(value string) Client {
	return c.clientWithNewParam("created_at_from", value)
}

func (c *client) CreatedAtTo(value string) Client {
	return c.clientWithNewParam("created_at_to", value)
}

func (c *client) UpdatedAtFrom(value string) Client {
	return c.clientWithNewParam("updated_at_from", value)
}

func (c *client) UpdatedAtTo(value string) Client {
	return c.clientWithNewParam("updated_at_to", value)
}

func (c *client) Limit(limit uint) Client {
	return c.clientWithNewParam("document_limit", limit)
}

func (c *client) Offset(offset uint) Client {
	return c.clientWithNewParam("start_at_document", offset)
}
