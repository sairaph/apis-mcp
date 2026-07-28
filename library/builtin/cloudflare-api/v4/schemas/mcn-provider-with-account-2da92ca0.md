---
title: mcn_provider_with_account
page_id: schema-mcn-provider-with-account-2da92ca0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_provider_with_account

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/mcn_provider"}, {"properties": {"account_id": {"$ref": "#/components/schemas/mcn_account_id"}}, "required": ["account_id"], "type": "object"}]}
```
