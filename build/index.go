package build

import (
	"context"

	"github.com/CoreumFoundation/callisto/build/callisto"
	"github.com/CoreumFoundation/callisto/build/hasura"
	"github.com/CoreumFoundation/crust/build/crust"
	"github.com/CoreumFoundation/crust/build/golang"
	"github.com/CoreumFoundation/crust/build/types"
)

// Commands is a definition of commands available in build system.
var Commands = map[string]types.Command{
	"build/me":    {Fn: crust.BuildBuilder, Description: "Builds the builder"},
	"build/znet":  {Fn: crust.BuildZNet, Description: "Builds znet binary"},
	"build":       {Fn: callisto.Build, Description: "Builds callisto binary"},
	"build/amd64": {Fn: callisto.BuildAMD64, Description: "Builds callisto binary for arm64 platform"},
	"build/arm64": {Fn: callisto.BuildARM64, Description: "Builds callisto binary for amd64 platform"},
	"images": {Fn: func(ctx context.Context, deps types.DepsFunc) error {
		deps(callisto.BuildDockerImage, hasura.BuildDockerImage)
		return nil
	}, Description: "Builds callisto and hasura docker images"},
	"images/callisto": {Fn: callisto.BuildDockerImage, Description: "Builds callisto image"},
	"images/hasura":   {Fn: hasura.BuildDockerImage, Description: "Builds hasura docker image"},
	"test":            {Fn: golang.Test, Description: "Runs unit tests"},
	"tidy":            {Fn: golang.Tidy, Description: "Runs go mod tidy"},
}
