---
title: magic_update_bgp_settings_request
page_id: schema-magic-update-bgp-settings-request-f07ee2ec
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_update_bgp_settings_request

```yaml
{"type": "object", "properties": {"cloudflare_asn": {"description": "Route advertisements from Cloudflare to ramps in this account will use this ASN.", "type": "integer", "format": "uint32", "example": 13335, "minimum": 1}, "redistribute": {"$ref": "#/components/schemas/magic_update_bgp_settings_redistribute_sources"}}}
```
