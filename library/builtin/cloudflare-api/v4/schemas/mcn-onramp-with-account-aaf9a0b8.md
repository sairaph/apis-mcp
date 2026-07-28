---
title: mcn_onramp_with_account
page_id: schema-mcn-onramp-with-account-aaf9a0b8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_onramp_with_account

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/mcn_onramp"}, {"properties": {"account_id": {"$ref": "#/components/schemas/mcn_account_id"}}, "required": ["account_id"], "type": "object"}]}
```
