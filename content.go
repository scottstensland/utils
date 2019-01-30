
package utils

import (
	"log"
	"github.com/kelseyhightower/envconfig"
	"fmt"
)

type Env_Vars struct {

	Tls_Cert	string
	Tls_Key		string
}

func get_env( ptr_curr_spec  *Env_Vars, env_var_prefix string) {

    // err := envconfig.Process("GKE", ptr_curr_spec)
    err := envconfig.Process(env_var_prefix, ptr_curr_spec)
    if err != nil {
        log.Fatal(err.Error())
    }

    fmt.Println("about to print env vars")

    format := " Tls_Cert : %v\n Tls_Key : %s\n "

    _, err = fmt.Printf(format, ptr_curr_spec.Tls_Cert, ptr_curr_spec.Tls_Key )

    if err != nil {
        log.Fatal(err.Error())
    }

	// ... confirm env var is populated

    if ptr_curr_spec.Tls_Cert == "" {

        panic("ERROR - failed to see env var populated : Tls_Cert")
    }
}


