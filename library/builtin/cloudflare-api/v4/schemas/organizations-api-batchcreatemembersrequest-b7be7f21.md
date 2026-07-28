---
title: organizations-api_BatchCreateMembersRequest
page_id: schema-organizations-api-batchcreatemembersrequest-b7be7f21
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# organizations-api_BatchCreateMembersRequest

```yaml
{"type": "object", "properties": {"members": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_CreateSingleMember"}, "maxItems": 10}}, "required": ["members"]}
```
