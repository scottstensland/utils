
package utils

import (
	"github.com/kelseyhightower/envconfig"
)

type Env_Vars struct {

	Tls_Cert	string
	Tls_Key		string
}

func Get_env( env_var_prefix string) (ptr_curr_spec  *Env_Vars, err_spec error) {

    err_spec = envconfig.Process(env_var_prefix, ptr_curr_spec)
    if err_spec != nil {
		return nil, err_spec
    }

	return ptr_curr_spec, nil
}

