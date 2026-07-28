---
title: access_reusable_policy_resp
page_id: schema-access-reusable-policy-resp-de873073
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_reusable_policy_resp

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_policy_resp"}, {"properties": {"app_count": {"$ref": "#/components/schemas/access_app_count"}, "reusable": {"type": "boolean", "enum": [true]}}}]}
```
