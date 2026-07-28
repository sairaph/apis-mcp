---
title: magic_ipsec_tunnels_psk_response
page_id: schema-magic-ipsec-tunnels-psk-response-df964dfb
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_ipsec_tunnels_psk_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/magic_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"successfully_applied_psks": {"description": "Map of tunnel IDs to successfully applied PSK details.", "type": "object", "additionalProperties": {"$ref": "#/components/schemas/magic_ipsec_tunnel_applied_psk"}}, "unapplied_psks": {"description": "Map of tunnel IDs to failure reasons for PSKs that could not be applied.", "type": "object", "additionalProperties": {"type": "string"}}}}}}]}
```
