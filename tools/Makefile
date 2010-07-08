include $(GOROOT)/src/Make.$(GOARCH)
.SUFFIXES: .go .$(O)

.go.6:
	$(GC) $<

all: serve tools

tools: now later

now: now.6
	6l -o now now.6

later: later.6
	6l -o later later.6

serve: serve.6
	6l -o serve serve.6

clean::
	rm -f *.6
	rm -f now
	rm -f later
	rm -f serve
