---
title: iam_create_payload
page_id: schema-iam-create-payload-73d9d56e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_create_payload

```yaml
{"type": "object", "properties": {"condition": {"$ref": "#/components/schemas/iam_condition"}, "expires_on": {"$ref": "#/components/schemas/iam_expires_on"}, "name": {"$ref": "#/components/schemas/iam_name"}, "not_before": {"$ref": "#/components/schemas/iam_not_before"}, "policies": {"$ref": "#/components/schemas/iam_token_policies"}}, "required": ["name", "policies"]}
```
