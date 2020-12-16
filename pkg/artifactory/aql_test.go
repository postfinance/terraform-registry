package artifactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindItems(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		exp := `items.find({"repo":{"$eq":"linux-generic-local"},"path":{"$eq":"terraform/providers"},"name":{"$eq":"terraform-provider-uam_linux_x86_64-0.0.2.zip"}})`
		act := FindItems("linux-generic-local", "terraform/providers", "terraform-provider-uam_linux_x86_64-0.0.2.zip")
		assert.Equal(t, exp, act.String())
	})

	t.Run("match", func(t *testing.T) {
		exp := `items.find({"repo":{"$eq":"linux-generic-local"},"path":{"$eq":"terraform/providers"},"name":{"$match":"terraform-provider-*"}})`
		act := FindItems("linux-generic-local", "terraform/providers", "terraform-provider-*")
		assert.Equal(t, exp, act.String())
	})
}
