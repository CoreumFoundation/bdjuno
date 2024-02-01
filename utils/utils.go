package utils

import (
	"context"

	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	"google.golang.org/grpc/metadata"
)

// RemoveDuplicateValues removes the duplicated values from the given slice
func RemoveDuplicateValues(slice []string) []string {
	keys := make(map[string]bool)
	var list []string

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// GetHeightRequestContext adds the height to the context for queries
func GetHeightRequestContext(context context.Context) context.Context {
	return metadata.AppendToOutgoingContext(
		context,
		grpctypes.GRPCBlockHeightHeader,
	)
}
