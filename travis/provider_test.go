package travis

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	p := Provider()
	assert.NotNil(t, p)

	sp := p.(*schema.Provider)
	assert.NotNil(t, sp)

	assert.NotNil(t, sp.DataSourcesMap)

	for _, key := range []string{
		"travis_expanded_cidr",
	} {
		dp, ok := sp.DataSourcesMap[key]
		assert.True(t, ok)
		assert.NotNil(t, dp)
	}
}
