server: serve.6 next.6
	6l -o serve serve.6

serve.6: serve.go next.6
	6g serve.go

next.6: next.go
	6g next.go

clean::
	rm -f *.6
	rm -f serve
