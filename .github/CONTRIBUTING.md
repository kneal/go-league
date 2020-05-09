# Contributing to go-league

Love to accept your contributions to this project! There are just a few small guidelines you need to follow.

## Issues

[Issues](issues/new/) are always welcome!

## Pull Requests

**NOTE: Recommend you start by opening a new issue describing the bug or feature you're intending to fix. Even if you think it's relatively minor, it's helpful to know what people are working on.**

These rules must be followed for any contributions to be merged into master. A Git installation is required.

1. Fork this repo
2. Go get the original code:

  `go get github.com/kneal/go-league`

3. Navigate to the original code:

  `$GOPATH/src/github.com/kneal/leaguego-league`

4. Add a remote branch pointing to your fork:

  `git remote add fork https://github.com/your_fork/go-league`

5. Implement desired changes
6. Validate the changes meet your desired use case
7. Write tests around the changes you made
8. Update documentation
9. Please run the below commands:

```sh
# Generate necessary code
go generate github.com/kneal/go-league/...

# Test the code
go test github.com/kneal/go-league/...

# Format the code
go fmt github.com/kneal/go-league/...

# Vet the code
go vet github.com/kneal/go-league/...
```

10. Push to your fork:

  `git push fork master`

11. Open a pull request. Thank you for your contribution! A dialog will ensue.