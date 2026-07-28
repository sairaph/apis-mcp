---
title: iam_token_base
page_id: schema-iam-token-base-5ee3887b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_token_base

```yaml
{"type": "object", "properties": {"condition": {"$ref": "#/components/schemas/iam_condition"}, "expires_on": {"$ref": "#/components/schemas/iam_expires_on"}, "id": {"$ref": "#/components/schemas/iam_token_identifier"}, "issued_on": {"$ref": "#/components/schemas/iam_issued_on"}, "last_used_on": {"$ref": "#/components/schemas/iam_last_used_on"}, "modified_on": {"$ref": "#/components/schemas/iam_modified_on"}, "name": {"$ref": "#/components/schemas/iam_name"}, "not_before": {"$ref": "#/components/schemas/iam_not_before"}, "policies": {"$ref": "#/components/schemas/iam_token_policies"}, "status": {"$ref": "#/components/schemas/iam_token_status"}}}
```
