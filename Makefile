build: res
	@go build -o vuego-test main.go

res:
	@cd resource && parcello -r;
