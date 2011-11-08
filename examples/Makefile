include $(GOROOT)/src/Make.inc

PREREQ=$(QUOTED_GOROOT)/pkg/$(GOOS)_$(GOARCH)/google-api-go-client.googlecode.com/hg/tasks/v1.a\
	$(QUOTED_GOROOT)/pkg/$(GOOS)_$(GOARCH)/google-api-go-client.googlecode.com/hg/urlshortener/v1.a\

TARG=go-api-demo

GOFILES=main.go\
	debug.go\
	orkut.go\
	tasks.go\
	urlshortener.go\

include $(GOROOT)/src/Make.cmd
