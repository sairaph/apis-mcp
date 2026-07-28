---
title: magic_bgp_settings
page_id: schema-magic-bgp-settings-dc6d42e6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_bgp_settings

```yaml
{"type": "object", "properties": {"cloudflare_asn": {"description": "Route advertisements from Cloudflare to ramps in this account will use this ASN.", "type": "integer", "format": "uint32", "example": 13335, "minimum": 1}, "modified_on": {"type": "string", "format": "date-time"}, "redistribute": {"$ref": "#/components/schemas/magic_bgp_redistribute_sources"}}, "required": ["cloudflare_asn", "redistribute"]}
```
