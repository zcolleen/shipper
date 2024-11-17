//go:build tools
// +build tools

package shipper

import (
	_ "github.com/gojuno/minimock/v3/cmd/minimock"
	_ "github.com/vektra/mockery/v2"
	_ "go.uber.org/mock/gomock"
)
