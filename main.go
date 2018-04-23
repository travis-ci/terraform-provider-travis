package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/travis-ci/terraform-provider-travis/travis"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: travis.Provider})
}
