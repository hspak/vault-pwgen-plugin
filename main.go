// Most boilerplate source: https://www.hashicorp.com/blog/building-a-vault-secure-plugin
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/vault/helper/pluginutil"
	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
	"github.com/hashicorp/vault/logical/plugin"
	"github.com/hspak/vault-pwgen-plugin/diceware"
)

type backend struct {
	*framework.Backend
	*diceware.Diceware
}

func main() {
	apiClientMeta := &pluginutil.APIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args[1:])

	tlsConfig := apiClientMeta.GetTLSConfig()
	tlsProviderFunc := pluginutil.VaultPluginTLSProvider(tlsConfig)

	if err := plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: Factory,
		TLSProviderFunc:    tlsProviderFunc,
	}); err != nil {
		log.Fatal(err)
	}
}

func Factory(ctx context.Context, c *logical.BackendConfig) (logical.Backend, error) {
	b := NewBackend(c)
	if err := b.Setup(ctx, c); err != nil {
		return nil, err
	}
	return b, nil
}

func NewBackend(c *logical.BackendConfig) *backend {
	var b backend
	b.Backend = &framework.Backend{
		BackendType: logical.TypeLogical,
		Paths: []*framework.Path{
			&framework.Path{
				Pattern: "pwgen",
				Fields: map[string]*framework.FieldSchema{
					"count": &framework.FieldSchema{
						Type: framework.TypeInt,
					},
				},
				Callbacks: map[logical.Operation]framework.OperationFunc{
					logical.UpdateOperation: b.pathGeneratePass,
				},
			},
		},
		// Deliberately leaving the endpoint unauthenticated -- don't see much reason to add auth
		PathsSpecial: &logical.Paths{
			Unauthenticated: []string{"pwgen"},
		},
	}
	b.Diceware = diceware.NewDiceware(diceware.WordList, diceware.RollList)
	return &b
}

func (b *backend) pathGeneratePass(_ context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	count := d.Get("count").(int)
	pass, err := b.Diceware.GeneratePass(count, diceware.RollCount)
	if err != nil {
		msg := fmt.Sprintf("error with diceware: %s", err.Error())
		return nil, errors.New(msg)
	}
	return &logical.Response{
		Data: map[string]interface{}{
			"password": pass,
		},
	}, nil
}
