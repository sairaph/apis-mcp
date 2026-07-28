---
title: iam_collection_membership_response_with_policies
page_id: schema-iam-collection-membership-response-with-policies-1b1285d9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_collection_membership_response_with_policies

```yaml
{"allOf": [{"$ref": "#/components/schemas/iam_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/iam_membership-with-policies"}}}, "type": "object"}], "title": "memberships_with_policies"}
```
