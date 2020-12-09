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

// curl -u sauterm -X POST -H "Content-Type: text/plain" -d 'items.find({"repo":{"$eq":"linux-generic-local"},"path":{"$eq":"terraform/providers"},"name":{"$match":"terraform-provider-*"}})' https://repo.pnet.ch/artifactory/api/search/aql

// FindItems build a items.find AQL expression
func FindItems(repo, path, name string) AQL {
	q := []string{
		buildExpression("repo", repo),
		buildExpression("path", path),
		buildExpression("name", name),
	}

	return AQL{
		query: fmt.Sprintf(`items.find({%s})`, strings.Join(q, ",")),
	}
}

func buildExpression(name, expr string) string {
	op := "$eq"
	if strings.Contains(expr, "*") {
		op = "$match"
	}

	return fmt.Sprintf("%q:{%q:%q}", name, op, expr)
}
