package callisto

import (
	"context"
	"path/filepath"

	"github.com/CoreumFoundation/crust/build/golang"
	"github.com/CoreumFoundation/crust/build/tools"
	"github.com/CoreumFoundation/crust/build/types"
)

const (
	repoPath   = "."
	binaryName = "callisto"
	binaryPath = "bin/" + binaryName
)

// Build builds callisto in docker.
func Build(ctx context.Context, deps types.DepsFunc) error {
	return buildCallisto(ctx, deps, tools.TargetPlatformLinuxLocalArchInDocker)
}

func BuildAMD64(ctx context.Context, deps types.DepsFunc) error {
	return buildCallisto(ctx, deps, tools.TargetPlatformLinuxAMD64InDocker)
}

func BuildARM64(ctx context.Context, deps types.DepsFunc) error {
	return buildCallisto(ctx, deps, tools.TargetPlatformLinuxARM64InDocker)
}

func buildCallisto(ctx context.Context, deps types.DepsFunc, targetPlatform tools.TargetPlatform) error {
	return golang.Build(ctx, deps, golang.BinaryBuildConfig{
		TargetPlatform: targetPlatform,
		PackagePath:    filepath.Join(repoPath, "cmd", "callisto"),
		BinOutputPath:  filepath.Join("bin", ".cache", binaryName, targetPlatform.String(), "bin", binaryName),
		CGOEnabled:     false,
	})
}
