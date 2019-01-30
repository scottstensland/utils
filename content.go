
package utils

import (
	"github.com/kelseyhightower/envconfig"
)

type Env_Vars struct {

	Tls_Cert	string
	Tls_Key		string
}

func Get_env( env_var_prefix string) ( Env_Vars, error) {

    var curr_spec Env_Vars

    err_spec := envconfig.Process(env_var_prefix, &curr_spec)
    if err_spec != nil {
		return curr_spec, err_spec
    }

	return curr_spec, nil
}

