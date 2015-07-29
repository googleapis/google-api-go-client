# Google APIs Client Library for Go

## Status
[![Build Status](https://travis-ci.org/google/google-api-go-client.png)](https://travis-ci.org/google/google-api-go-client)

These are auto-generated Go libraries from the Google Discovery Service's JSON description files of the available "new style" Google APIs.

Due to the auto-generated nature of this collection of libraries, complete APIs or specific versions can appear or go away without notice.
As a result, you should always locally vendor any API(s) that your code relies upon.

Announcement email:
http://groups.google.com/group/golang-nuts/browse_thread/thread/6c7281450be9a21e

Getting started documentation:

   https://github.com/google/google-api-go-client/blob/master/GettingStarted.md

In summary:

```
$ go get google.golang.org/api/storage/v1
$ go get google.golang.org/api/tasks/v1
$ go get google.golang.org/api/moderator/v1
... etc ...
```

For docs, see e.g.:

   https://godoc.org/google.golang.org/api/storage/v1

The package of a given import is the second-to-last component, before the version number.

For examples, see:

   https://github.com/google/google-api-go-client/tree/master/examples

For support, use the golang-nuts@ mailing list:

   https://groups.google.com/group/golang-nuts

## Application Default Credentials Example

The Application Default Credentials provide a simple way to get authorization
credentials for use in calling Google APIs.

They are best suited for cases when the call needs to have the same identity and
authorization level for the application independent of the user.
This is the recommended approach to authorize calls to Cloud APIs, particularly
when you're building an application that uses Google Compute Engine.

```
import (
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func ExampleDefaultClient() {
	client, err := google.DefaultClient(oauth2.NoContext,
		"https://www.googleapis.com/auth/devstorage.full_control")
	if err != nil {
		log.Fatal(err)
	}
	client.Get("...")
}
```
