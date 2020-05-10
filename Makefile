clean:
	#################################
	######      Go clean       ######
	#################################

	@go mod tidy
	@go vet ./league
	@go fmt ./league