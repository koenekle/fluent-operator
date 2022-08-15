package plugins

import (
	"github.com/fluent/fluent-operator/apis/fluentbit/v1alpha2/plugins/params"
)

type Common struct {
	// A user-friendly alias name for this plugin.
	// Used in metrics for distinction of each configured plugin.
	Alias string `json:"alias,omitempty"`
}

func (c *Common) CommonParams(_ SecretLoader) (*params.KVs, error) {
	kvs := params.NewKVs()
	if c.Alias != "" {
		kvs.Insert("Alias", c.Alias)
	}
	return kvs, nil
}
