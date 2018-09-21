# vault-pwgen-plugin
### Vault plugin to support password generation using [EFF's diceware list](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases).

## Installation

https://www.vaultproject.io/docs/plugin/index.html

```sh
# Example Setup
vault write sys/plugins/catalog/pwgen-plugin \
  sha_256="$SHASUM" \
  command="vault-pwgen-plugin"
vault secrets enable -path=diceware -plugin-name=pwgen-plugin plugin
```

## Usage

via Vault CLI:
```
vault write diceware/pwgen count=2
Key         Value
---         -----
password    abmatchlessabrocklike
```
via REST API:
```
curl --data '{"count": 6}' "$VAULT_ADDR/v1/diceware/pwgen"
{
  "request_id": "c4f15a2b-d487-46ae-2a63-e14ccb49f1ec",
  "lease_id": "",
  "renewable": false,
  "lease_duration": 0,
  "data": {
    "password": "abdisabledabmatchlessabrocklikeabreverendabatlasabhandwrite"
  },
  "wrap_info": null,
  "warnings": null,
  "auth": null
}
```

## License

MIT
