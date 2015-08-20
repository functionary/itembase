Go Itembase
===========

[![GoDoc][godoc-badge]][godoc]

A Go (golang) REST client library for [the Itembase API]. Supports all entities
from the [Itembase API endpoints]:

  - Buyers
  - Products
  - Store Profiles
  - Transactions

[the Itembase API]: http://business.itembase.com/connect
[Itembase API endpoints]: http://sandbox.api.itembase.io/swagger-ui/index.html

Installation
------------

- Setup your `GOPATH` and workspace. If you are new to Go and you're not sure how
  to do this, read [How to Write Go Code](https://golang.org/doc/code.html).
- Download the package:
```sh
go get github.com/saasbuilders/itembase
```

Usage
-----

High-level examples are shown here, see [the GoDoc][godoc] for the complete API.

### Config

```go
import "github.com/saasbuilders/itembase"

config := itembase.Config{
	ClientID:     "YOUR CLIENT ID",
	ClientSecret: "YOUR CLIENT SECRET",
	Scopes:       []string{"user.minimal", "connection.transaction", "connection.product", "connection.profile", "connection.buyer"},
	TokenHandler: TokenHandler(),
	RedirectURL:  "http://yourredirecturl.com",
	Production:   false,
}
```

### Instantiating

```go
storeRef := itembase.
	New(config, nil).
	User("13ac2c74-7de3-4436-9a6d-2c94dd2b1fd3")

me, err := storeRef.Me()
if err != nil {
	log.Fatal(err)
}

pretty.Println(me)
```

### Queries

```go
var transactions itembase.Transactions
err = storeRef.Transactions().Select("6ee2e2d9f7baea5132ab79b").GetInto(&transactions)

if err != nil {
	log.Fatal(err)
}
pretty.Println(transactions)
```

### Querying Buyers

```go
pretty.Println(storeRef.Buyers().Select("95d5c9ceeaad98706ce").URL())

var buyers itembase.Buyers
err := storeRef.Buyers().GetInto(&buyers)

if err != nil {
	log.Fatal(err)
}
pretty.Println(buyers)
```

### Querying the Store Profile

```go
pretty.Println(storeRef.Profiles().URL())

var profiles itembase.Profiles
err := storeRef.Products().GetInto(&profiles)

if err != nil {
	log.Fatal(err)
}
pretty.Println(profiles)
```

### Querying Store Products

```go
pretty.Println(storeRef.Products().URL())
pretty.Println(storeRef.Products().Select("ee6f8dc930f5bcb671a0").URL())

var products itembase.Products
err := storeRef.Products().GetInto(&products)

if err != nil {
	log.Fatal(err)
}
pretty.Println(products)

```

### Querying Store Transactions

```go
pretty.Println(storeRef.Transactions().URL())
pretty.Println(storeRef.Transactions().Select("6ee2e2d9f7baea5132ab79b").URL())
pretty.Println(storeRef.Transactions().CreatedAtFrom("2015-04-29T08:53:01.738+0200").Limit(2).Offset(6).URL())

var transactions itembase.Transactions
err := storeRef.Transactions().CreatedAtFrom("2015-05-07T09:53:01").Limit(3).Offset(6).GetInto(&transactions)

if err != nil {
	log.Fatal(err)
}
pretty.Println(transactions)

```

### Query Functions

You can stack the different limitation options when it makes sense like so :

- Date specific filters - works for Transactions and Products
```go
storeRef.Transactions().CreatedAtFrom("2015-05-07T09:53:01")
storeRef.Transactions().CreatedAtTo("2015-05-07T09:53:01")
storeRef.Transactions().UpdatedAtFrom("2015-05-07T09:53:01")
storeRef.Transactions().UpdatedAtTo("2015-05-07T09:53:01")
```

- Limits and pagination. Similar to SQL Limit syntax. Works with Transactions and Products
```go
storeRef.Transactions().Limit(10)
storeRef.Transactions().Limit(10).Offset(10)
```

- Selecting specific entries - works with Transactions, Buyers and Products
```go
storeRef.Transactions().Select("6ee2e2d9f7baea5132ab79b")
```

### Token Handlers

You will want to add your own token handlers to save tokens in your own datastore / database. You can define how to retrieve the oauth token for a user, how to save it, and what to do when it expires. Set to nil if you don't want to override the usual functions, which would only make sense for the last one, as the saving and loading should be handled.

```go
func TokenHandler() itembase.ItembaseTokens {
	return itembase.ItembaseTokens{
		GetCachedToken, // How to retrieve a valid oauth token for a user
		SaveToken,      // How to save a valid oauth token for a user
		nil,            // What to do in case of expired tokens
	}
}

func GetCachedToken(userID string) (token *oauth2.Token, err error) {

	// retrieve oauth2.Token from your Database and assign it to &token

	if token == nil {
		err = errors.New("No Refresh Token!")
	}

	return
}

func SaveToken(userID string, token *oauth2.Token) (err error) {

	// save oauth2.Token to your Database for userID

	return
}

func TokenPermissions(authURL string) (authcode string, err error) {
	
	// token expired, offline authURL provided
	// handle the token permission process, and return the new authcode

	return
}

```

Off you go, now enjoy.

Credits
-------

Originally based on the great work of [cosn], [JustinTulloss] and [ereyes01] on
the [Firebase API client].

[cosn]: https://github.com/cosn
[JustinTulloss]: https://github.com/JustinTulloss
[ereyes01]: https://github.com/ereyes01
[Firebase API client]: https://github.com/ereyes01/firebase

[godoc-badge]: http://img.shields.io/badge/godoc-reference-blue.svg?style=flat
[godoc]: https://godoc.org/github.com/saasbuilders/itembase
