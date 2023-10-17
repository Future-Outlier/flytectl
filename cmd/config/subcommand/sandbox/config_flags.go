// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package sandbox

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (Config) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (Config) mustJsonMarshal(v interface{}) string {
	raw, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(raw)
}

func (Config) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.StringVar(&DefaultConfig.DeprecatedSource, fmt.Sprintf("%v%v", prefix, "source"), DefaultConfig.DeprecatedSource, "deprecated,  path of your source code,  please build images with local daemon")
	cmdFlags.StringVar(&DefaultConfig.Version, fmt.Sprintf("%v%v", prefix, "version"), DefaultConfig.Version, "Version of flyte. Only supports flyte releases greater than v0.10.0")
	cmdFlags.StringVar(&DefaultConfig.Image, fmt.Sprintf("%v%v", prefix, "image"), DefaultConfig.Image, "Optional. Provide a fully qualified path to a Flyte compliant docker image.")
	cmdFlags.BoolVar(&DefaultConfig.Prerelease, fmt.Sprintf("%v%v", prefix, "pre"), DefaultConfig.Prerelease, "Optional. Pre release Version of flyte will be used for sandbox.")
	cmdFlags.BoolVar(&DefaultConfig.DisableAgent, fmt.Sprintf("%v%v", prefix, "disable-agent"), DefaultConfig.DisableAgent, "Optional. Disable the agent service.")
	cmdFlags.StringSliceVar(&DefaultConfig.Env, fmt.Sprintf("%v%v", prefix, "env"), DefaultConfig.Env, "Optional. Provide Env variable in key=value format which can be passed to sandbox container.")
	cmdFlags.Var(&DefaultConfig.ImagePullPolicy, fmt.Sprintf("%v%v", prefix, "imagePullPolicy"), "Optional. Defines the image pull behavior [Always/IfNotPresent/Never]")
	cmdFlags.StringVar(&DefaultConfig.ImagePullOptions.RegistryAuth, fmt.Sprintf("%v%v", prefix, "imagePullOptions.registryAuth"), DefaultConfig.ImagePullOptions.RegistryAuth, "The base64 encoded credentials for the registry.")
	cmdFlags.StringVar(&DefaultConfig.ImagePullOptions.Platform, fmt.Sprintf("%v%v", prefix, "imagePullOptions.platform"), DefaultConfig.ImagePullOptions.Platform, "Forces a specific platform's image to be pulled.'")
	cmdFlags.BoolVar(&DefaultConfig.Dev, fmt.Sprintf("%v%v", prefix, "dev"), DefaultConfig.Dev, "Optional. Only start minio and postgres in the sandbox.")
	cmdFlags.BoolVar(&DefaultConfig.DryRun, fmt.Sprintf("%v%v", prefix, "dryRun"), DefaultConfig.DryRun, "Optional. Only print the docker commands to bring up flyte sandbox/demo container.This will still call github api's to get the latest flyte release to use'")
	return cmdFlags
}
