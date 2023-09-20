// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package sandbox

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_Config(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_Config(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_Config(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_Config(val, result))
}

func testDecodeRaw_Config(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_Config(vStringSlice, result))
}

func TestConfig_GetPFlagSet(t *testing.T) {
	val := Config{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestConfig_SetFlags(t *testing.T) {
	actual := Config{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_source", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("source", testValue)
			if vString, err := cmdFlags.GetString("source"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.DeprecatedSource)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_version", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("version", testValue)
			if vString, err := cmdFlags.GetString("version"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Version)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_image", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("image", testValue)
			if vString, err := cmdFlags.GetString("image"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Image)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_pre", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("pre", testValue)
			if vBool, err := cmdFlags.GetBool("pre"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.Prerelease)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_enable-agent", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("enable-agent", testValue)
			if vBool, err := cmdFlags.GetBool("enable-agent"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.EnableAgent)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_env", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := join_Config(DefaultConfig.Env, ",")

			cmdFlags.Set("env", testValue)
			if vStringSlice, err := cmdFlags.GetStringSlice("env"); err == nil {
				testDecodeRaw_Config(t, join_Config(vStringSlice, ","), &actual.Env)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_imagePullPolicy", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("imagePullPolicy", testValue)
			if v := cmdFlags.Lookup("imagePullPolicy"); v != nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", v.Value.String()), &actual.ImagePullPolicy)

			}
		})
	})
	t.Run("Test_imagePullOptions.registryAuth", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("imagePullOptions.registryAuth", testValue)
			if vString, err := cmdFlags.GetString("imagePullOptions.registryAuth"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.ImagePullOptions.RegistryAuth)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_imagePullOptions.platform", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("imagePullOptions.platform", testValue)
			if vString, err := cmdFlags.GetString("imagePullOptions.platform"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.ImagePullOptions.Platform)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_dev", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("dev", testValue)
			if vBool, err := cmdFlags.GetBool("dev"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.Dev)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_dryRun", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("dryRun", testValue)
			if vBool, err := cmdFlags.GetBool("dryRun"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.DryRun)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
