package main

import (
	"context"

	_ "github.com/benthosdev/benthos/v4/public/components/io"
	_ "github.com/benthosdev/benthos/v4/public/components/kafka"
	_ "github.com/benthosdev/benthos/v4/public/components/mongodb"
	_ "github.com/benthosdev/benthos/v4/public/components/prometheus"
	"github.com/benthosdev/benthos/v4/public/service"
)

//lint:ignore U1000 Ignore unused function
func main() {
	service.RunCLI(context.Background())
}
