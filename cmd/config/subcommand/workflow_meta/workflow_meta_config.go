package workflow_meta

//go:generate pflags Config --default-var DefaultConfig --bind-default-var

var (
	DefaultConfig = &Config{}
)

// Config commandline configuration
type Config struct {
	Version string `json:"version" pflag:",version of the workflow to be fetched."`
}
