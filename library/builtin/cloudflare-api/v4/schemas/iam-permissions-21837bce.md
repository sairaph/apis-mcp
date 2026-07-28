---
title: iam_permissions
page_id: schema-iam-permissions-21837bce
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_permissions

```yaml
{"type": "object", "properties": {"analytics": {"$ref": "#/components/schemas/iam_grants"}, "billing": {"$ref": "#/components/schemas/iam_grants"}, "cache_purge": {"$ref": "#/components/schemas/iam_grants"}, "dns": {"$ref": "#/components/schemas/iam_grants"}, "dns_records": {"$ref": "#/components/schemas/iam_grants"}, "lb": {"$ref": "#/components/schemas/iam_grants"}, "logs": {"$ref": "#/components/schemas/iam_grants"}, "organization": {"$ref": "#/components/schemas/iam_grants"}, "ssl": {"$ref": "#/components/schemas/iam_grants"}, "waf": {"$ref": "#/components/schemas/iam_grants"}, "zone_settings": {"$ref": "#/components/schemas/iam_grants"}, "zones": {"$ref": "#/components/schemas/iam_grants"}}, "example": {"analytics": {"read": true, "write": false}, "zones": {"read": true, "write": true}}}
```
