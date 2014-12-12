# Getting Started

This is a quick walk-through of how to get started with the Google APIs for Go.

# Background

The first thing to understand is that the Google API libraries are auto-generated for each language, including Go, so they may not feel like 100% natural for any language.  The Go versions are pretty natural, but please forgive any small non-idiomatic things.  (Suggestions welcome, though!)

# Installing

Pick an API and a version of that API to install.  You can find the complete list by looking at the <a href='http://code.google.com/p/google-api-go-client/source/browse/'>directories here</a>.

For example, let's install the urlshortener's version 1 API:

```
$ go get code.google.com/p/google-api-go-client/urlshortener/v1
```

Now it's ready for use in your code.

# Using

Once you've installed a library, you import it like this:

```
package main

import (
    "code.google.com/p/google-api-go-client/urlshortener/v1"
)
```

The package name, if you don't override it on your import line, is the name of the API without the version number.  In the case above, just ` urlshortener `.

# Instantiating

Each API has a ` New ` function taking an ` *http.Client ` and returning an API-specific ` *Service `.

You create the service like:

```
    svc, err := urlshortener.New(httpClient)
```

# OAuth HTTP Client

The HTTP client you pass in to the service must be one that automatically adds Google-supported Authorization information to the requests.

The best option is to use http://code.google.com/p/goauth2/ , an OAuth2 library for Go.  You can see how to set use goauth2 with these APIs by checking out the <a href='http://code.google.com/p/google-api-go-client/source/browse/#hg%2Fexamples'>example code</a>.

In summary, you need to create an OAuth config:

```
    var config = &oauth.Config{
        ClientId:     "", // from https://code.google.com/apis/console/
        ClientSecret: "", // from https://code.google.com/apis/console/
        Scope:        urlshortener.UrlshortenerScope,
        AuthURL:      "https://accounts.google.com/o/oauth2/auth",
        TokenURL:     "https://accounts.google.com/o/oauth2/token",
    }
```

Then you need to get an OAuth Token from the user.  This involves sending the user to a URL (at Google) to grant access to your application (either a web application or a desktop application), and then the browser redirects to the website or local application's webserver with the per-user token in the URL.

Once you have that token,

```
    transport := &oauth.Transport{
        Token:     token,
        Config:    config,
        Transport: http.DefaultTransport,
    }

    httpClient := transport.Client()
```

Then you're good to pass that client to the API's ` New ` function.

# Using API Keys

Some APIs require passing API keys from your application.  To do this, you can use <a href='http://godoc.org/code.google.com/p/google-api-go-client/googleapi/transport#APIKey'>transport.APIKey</a>:

```
    client := &http.Client{
        Transport: &transport.APIKey{Key: developerKey},
    }
```

# Using the Service

Each service contains zero or more methods and zero or more sub-services.  The sub-services related to a specific type of "Resource".

Those sub-services then contain their own methods.

For instance, the urlshortener API has just the "Url" sub-service:

```
    url, err := svc.Url.Get(shortURL).Do()
    if err != nil {
        ...
    }
    fmt.Printf("The URL %s goes to %s\n", shortURL, url.LongUrl)
```

For a more complete example, see the http://code.google.com/p/google-api-go-client/source/browse/examples/urlshortener.go in the <a href='http://code.google.com/p/google-api-go-client/source/browse/examples/'>examples directory</a>.