include $(GOROOT)/src/Make.inc

TARG=github.com/alloy-d/now.go
GOFILES=now.go\
		serve.go\
		serve-to-humans.go\
		serve-to-machines.go\
		dirty-rotten-hacks.go\

include $(GOROOT)/src/Make.pkg

