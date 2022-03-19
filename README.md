go command
- go test : running unit test
- go test -coverprofile cover.out ./package_name : running test with coverage
- go test -run ^TestDiagonalDifference$ hackerrank/practice : run test on specific testcase
- go clean -testcache
- go tool cover -html=cover.out : view coverage result in html
- go mod init : initialize new module in current directory
