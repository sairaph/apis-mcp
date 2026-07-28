---
title: dns-settings_soa-response
page_id: schema-dns-settings-soa-response-a5aa142e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-settings_soa-response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-settings_soa-base"}, {"required": ["mname", "rname", "refresh", "retry", "expire", "min_ttl", "ttl"], "type": "object"}]}
```
