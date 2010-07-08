include $(GOROOT)/src/Make.$(GOARCH)
.SUFFIXES: .go .$(O)

.go.$(O):
	$(GC) $<

GOFILES=now.go\
		serve.go\
		serve-to-humans.go\
		serve-to-machines.go

lib: $(GOFILES)
	$(GC) $^

clean::
	rm -f *.6
	rm -f serve
