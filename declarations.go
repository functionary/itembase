package itembase

import "golang.org/x/oauth2"

type Config struct {
	ClientID     string
	ClientSecret string
	Scopes       []string
	TokenHandler ItembaseTokens
	Production   bool
	RedirectURL  string
}

type Client interface {
	// Returns the absolute URL path for the client
	Url() string

	//Gets the value referenced by the client and unmarshals it into
	// the passed in destination.
	GetInto(destination interface{}) error
	Get() (destination interface{}, err error)

	Me() (destination User, err error)

	// Child returns a reference to the child specified by `path`. This does not
	// actually make a request to itembase, but you can then manipulate the reference
	// by calling one of the other methods (such as `GetInto` or `Get`).
	Child(path string) Client

	Transactions() Client
	Products() Client
	Profiles() Client
	Buyers() Client

	Sandbox() Client

	User(path string) Client

	Select(prop string) Client
	CreatedAtFrom(value string) Client
	CreatedAtTo(value string) Client
	UpdatedAtFrom(value string) Client
	UpdatedAtTo(value string) Client
	Limit(limit uint) Client
	Offset(offset uint) Client

	SaveToken(userId string, token *oauth2.Token) (err error)
	GetCachedToken(userId string) (token *oauth2.Token, err error)
	GiveTokenPermissions(authUrl string) (authcode string, err error)
}

// Api is the internal interface for interacting with Itembase. The internal
// implementation of this interface is responsible for all HTTP operations that
// communicate with Itembase.
//
// Users of this library can implement their own Api-conformant types for
// testing purposes. To use your own test Api type, pass it in to the NewClient
// function.
type Api interface {
	// Call is responsible for performing HTTP transactions such as GET, POST,
	// PUT, PATCH, and DELETE. It is used to communicate with Itembase by all
	// of the Client methods.
	//
	// Arguments are as follows:
	//  - `method`: The http method for this call
	//  - `path`: The full itembase url to call
	//  - `body`: Data to be marshalled to JSON (it's the responsibility of Call to do the marshalling and unmarshalling)
	//  - `params`: Additional parameters to be passed to itembase
	//  - `dest`: The object to save the unmarshalled response body to.
	//    It's up to this method to unmarshal correctly, the default implemenation just uses `json.Unmarshal`
	Call(method, path, auth string, body interface{}, params map[string]string, dest interface{}) error
}

type ItembaseTokens struct {
	TokenLoader      TokenLoader
	TokenSaver       TokenSaver
	TokenPermissions TokenPermissions
}

type TokenSaver func(userId string, token *oauth2.Token) (err error)
type TokenLoader func(userId string) (token *oauth2.Token, err error)
type TokenPermissions func(authUrl string) (authcode string, err error)
