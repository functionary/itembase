package itembase

import "golang.org/x/oauth2"

// A Config structure is used to configure an itembase Client instance.
type Config struct {
	// ClientID is the OAuth2 application ID for a registered itembase app.
	// See oauth2.Config.
	ClientID string

	// ClientSecret is the application's OAuth2 secret credential.
	// See oauth2.Config.
	ClientSecret string

	// Scopes specify requested OAuth2 permissions, as defined by the itembase
	// API. See oauth2.Config.
	Scopes []string

	// A TokenHandler provides handlers for lifecycle events of OAuth2 tokens.
	TokenHandler ItembaseTokens

	// Production may be set to false to put a Client into sandbox mode.
	Production bool

	// RedirectURL is the URL to redirect users after requesting OAuth2
	// permission grants from itembase. See oauth2.Config.
	RedirectURL string
}

// A Client retrieves data from the itembase API. Use itembase.New to create an
// instance of the default implementation.
//
// TODO: document each method
type Client interface {
	// Returns the absolute URL path for the client
	URL() string

	// Gets the value referenced by the client and unmarshals it into
	// the passed in destination.
	GetInto(destination interface{}) error

	// Paginates through all possible values from client, and unmarshals
	// into the passed in destination
	GetAllInto(destination interface {
		Add(interface{})
	}) error

	// Gets values referenced by the client, and returns them as generic interface(!)
	Get() (destination interface{}, err error)

	Me() (destination User, err error)
	Activate() (destination interface{}, err error)

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

	SaveToken(userID string, token *oauth2.Token) (err error)
	GetCachedToken(userID string) (token *oauth2.Token, err error)
	GiveTokenPermissions(authURL string) (authcode string, err error)

	HandleOAuthCode(authcode string) (*oauth2.Token, error)
	GetUserIDForToken(token *oauth2.Token) (string, error)
}

// API is the internal interface for interacting with Itembase. The internal
// implementation of this interface is responsible for all HTTP operations that
// communicate with Itembase.
//
// Users of this library can implement their own API-conformant types for
// testing purposes. To use your own test API type, pass it in to the NewClient
// function.
type API interface {
	// Call is responsible for performing HTTP transactions such as GET, POST,
	// PUT, PATCH, and DELETE. It is used to communicate with Itembase by all
	// of the Client methods.
	//
	// Arguments are as follows:
	//  - `method`: The http method for this call
	//  - `path`: The full itembase url to call
	//	- `auth`: TODO
	//  - `body`: Data to be marshalled to JSON (it's the responsibility of Call to do the marshalling and unmarshalling)
	//  - `params`: Additional parameters to be passed to itembase
	//  - `dest`: The object to save the unmarshalled response body to.
	//    It's up to this method to unmarshal correctly, the default implementation just uses `json.Unmarshal`
	Call(method, path, auth string, body interface{}, params map[string]string, dest interface{}) error
}

// ItembaseTokens is a container struct holding handler functions for events in
// an OAuth2 token's lifecycle.
type ItembaseTokens struct {
	TokenLoader      TokenLoader
	TokenSaver       TokenSaver
	TokenPermissions TokenPermissions
}

// A TokenSaver is called at points during OAuth2 authorization flow when an
// application might wish to persist the given token to a data store or cache.
type TokenSaver func(userID string, token *oauth2.Token) (err error)

// A TokenLoader is called at points during OAuth2 authorization flow when an
// application might wish to retrieve a persisted token from a data store.
type TokenLoader func(userID string) (token *oauth2.Token, err error)

// A TokenPermissions handler is called at points during OAuth2 authorization
// flow when a grantor might have granted new permissions for an authorization,
// such as new scopes.
type TokenPermissions func(authURL string) (authcode string, err error)
