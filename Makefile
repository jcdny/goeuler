include $(GOROOT)/src/Make.inc

TARG=tables
GOFILES=\
	doc.go\
	factor.go\
	gcdlcm.go\
	init.go\
	primes.go\
	permute.go\
	readers.go\
	util.go\

include $(GOROOT)/src/Make.pkg

.PHONY: gofmt
gofmt:
	gofmt -w -s $(GOFILES)

