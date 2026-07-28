---
title: iam_collection_member_response_with_policies
page_id: schema-iam-collection-member-response-with-policies-2d3da849
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_collection_member_response_with_policies

```yaml
{"allOf": [{"$ref": "#/components/schemas/iam_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/iam_member_with_policies"}}}, "type": "object"}], "title": "Members with Policies"}
```
