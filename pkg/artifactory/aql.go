package artifactory

import (
	"fmt"
	"strings"
)

// AQL expression
type AQL struct {
	query string
}

func (a AQL) String() string {
	return a.query
}

// Bytes returns the query a byte slice
func (a AQL) Bytes() []byte {
	return []byte(a.query)
}

// FindItems build a items.find AQL expression
func FindItems(repo, path, name string) AQL {
	q := []string{
		buildExpression("repo", repo),
		buildExpression("path", path),
		buildExpression("name", name),
	}

	return AQL{
		query: fmt.Sprintf(`items.find({%s}).include("repo", "path", "name", "sha256")`, strings.Join(q, ",")),
	}
}

func buildExpression(name, expr string) string {
	op := "$eq"
	if strings.Contains(expr, "*") {
		op = "$match"
	}

	return fmt.Sprintf("%q:{%q:%q}", name, op, expr)
}
