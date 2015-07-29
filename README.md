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

Application Default Credentials provide a simplified way to obtain credentials
for authenticating with Google APIs.

The Application Default Credentials authenticate as the application itself,
which make them great for working with Google Cloud APIs like Storage or
Datastore. They are the recommend form of authentication when building
applications that run on Google Compute Engine or Google App Engine.

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
