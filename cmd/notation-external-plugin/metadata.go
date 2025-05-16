package main

import (
	"github.com/notaryproject/notation-go/plugin/proto"
)

var (
	Version       = "v0.1.0"
	BuildMetadata = "unreleased"
)

func getVersion() string {
	if BuildMetadata == "" {
		return Version
	}
	return Version + "+" + BuildMetadata
}

func runGetMetadata() *proto.GetMetadataResponse {
	return &proto.GetMetadataResponse{
		Name:                      "external-plugin",
		Description:               "Sign artifacts with any external signer plugin",
		Version:                   getVersion(),
		URL:                       "https://github.com/tuminoid/notation-external-plugin",
		SupportedContractVersions: []string{proto.ContractVersion},
		Capabilities:              []proto.Capability{proto.CapabilitySignatureGenerator},
	}
}
