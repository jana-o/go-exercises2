package saying

import "fmt"

func Greet(s string) string {
	return fmt.Sprint("Welcome my dear ", s)
}

///godoc -http=:8080

// go test
// go test -bench .
// go test -cover
// go test -coverprofile c.out
// go tool cover -html=c.out
