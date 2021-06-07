// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package register

import (
	"fmt"

	"github.com/spf13/pflag"
)

// GetPFlagSet will return strongly types pflags for all fields in FilesConfig and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg FilesConfig) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("FilesConfig", pflag.ExitOnError)
	cmdFlags.StringVarP(&DefaultFilesConfig.Version, fmt.Sprintf("%v%v", prefix, "version"), "v", DefaultFilesConfig.Version, "version of the entity to be registered with flyte.")
	cmdFlags.BoolVarP(&DefaultFilesConfig.ContinueOnError, fmt.Sprintf("%v%v", prefix, "continueOnError"), "c", DefaultFilesConfig.ContinueOnError, "continue on error when registering files.")
	cmdFlags.BoolVarP(&DefaultFilesConfig.Archive, fmt.Sprintf("%v%v", prefix, "archive"), "a", DefaultFilesConfig.Archive, "pass in archive file either an http link or local path.")
	cmdFlags.StringVarP(&DefaultFilesConfig.AssumableIamRole, fmt.Sprintf("%v%v", prefix, "assumableIamRole"), "i", DefaultFilesConfig.AssumableIamRole, " Custom assumable iam auth role to register launch plans with.")
	cmdFlags.StringVarP(&DefaultFilesConfig.K8ServiceAccount, fmt.Sprintf("%v%v", prefix, "k8ServiceAccount"), "k", DefaultFilesConfig.K8ServiceAccount, " custom kubernetes service account auth role to register launch plans with.")
	cmdFlags.StringVarP(&DefaultFilesConfig.OutputLocationPrefix, fmt.Sprintf("%v%v", prefix, "outputLocationPrefix"), "l", DefaultFilesConfig.OutputLocationPrefix, " custom output location prefix for offloaded types (files/schemas).")
	return cmdFlags
}