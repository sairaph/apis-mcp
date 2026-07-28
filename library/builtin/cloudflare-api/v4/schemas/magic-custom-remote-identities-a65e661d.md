---
title: magic_custom_remote_identities
page_id: schema-magic-custom-remote-identities-a65e661d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_custom_remote_identities

```yaml
{"type": "object", "properties": {"fqdn_id": {"description": "A custom IKE ID of type FQDN that may be used to identity the IPsec tunnel. The\ngenerated IKE IDs can still be used even if this custom value is specified.\n\nMust be of the form `<custom label>.<account ID>.custom.ipsec.cloudflare.com`.\n\nThis custom ID does not need to be unique. Two IPsec tunnels may have the same custom\nfqdn_id. However, if another IPsec tunnel has the same value then the two tunnels\ncannot have the same cloudflare_endpoint.", "type": "string"}}}
```
