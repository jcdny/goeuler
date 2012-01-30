include $(GOROOT)/src/Make.inc

TARG=tables
GOFILES=\
	check.go\
	doc.go\
	factor.go\
	gcdlcm.go\
	init.go\
	primes.go\
	permute.go\
	pythagoras.go\
	readers.go\
	util.go\

include $(GOROOT)/src/Make.pkg

.PHONY: gofmt
gofmt:
	gofmt -w -s $(GOFILES)

