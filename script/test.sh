#!/bin/bash

mkdir -p /tmp/vault-plugins
go build -v -o /tmp/vault-plugins/vault-pwgen-plugin

export VAULT_ADDR=http://localhost:8200
export VAULT_TOKEN='root'

echo 'plugin_directory = "/tmp/vault-plugins"' > /tmp/vault.hcl

vault server -dev -dev-root-token-id="$VAULT_TOKEN" -config=/tmp/vault.hcl &
VAULT_PID=$!

SHASUM=$(shasum -a 256 "/tmp/vault-plugins/vault-pwgen-plugin" | cut -d' ' -f1)
vault write sys/plugins/catalog/pwgen-plugin \
  sha_256="$SHASUM" \
  command="vault-pwgen-plugin"

vault secrets enable -path=diceware -plugin-name=pwgen-plugin plugin


for i in $(seq 1 64); do
  curl --data "{\"count\": $i}" "$VAULT_ADDR/v1/diceware/pwgen"
  vault write diceware/pwgen count=$i
done

kill $VAULT_PID
rm -rf /tmp/vault-plugins /tmp/vault.hcl
