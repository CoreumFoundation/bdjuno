module github.com/CoreumFoundation/bdjuno/build

go 1.21

require (
	github.com/CoreumFoundation/coreum-tools v0.4.1-0.20240321120602-0a9c50facc68
	github.com/CoreumFoundation/crust/build v0.0.0-20240409095900-0a67f05983fd
)

require (
	github.com/pkg/errors v0.9.1 // indirect
	github.com/samber/lo v1.39.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	// Make sure to not bump x/exp dependency without cosmos-sdk updated because their breaking change is not compatible
	// with cosmos-sdk v0.47.
	// Details: https://github.com/cosmos/cosmos-sdk/issues/18415
	golang.org/x/exp v0.0.0-20230713183714-613f0c0eb8a1 // indirect
	golang.org/x/mod v0.12.0 // indirect
)
