---
title: iam_resources
page_id: schema-iam-resources-c8de8a2a
path: schemas
description: A list of resource names that the policy applies to.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_resources

A list of resource names that the policy applies to.

```yaml
{"description": "A list of resource names that the policy applies to.", "example": {"object": {"summary": "Nested object value", "value": {"com.cloudflare.api.account.eb78d65290b24279ba6f44721b3ea3c4": {"com.cloudflare.api.account.zone.*": "*"}}}, "string": {"summary": "Single string value", "value": {"com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43": "*"}}}, "oneOf": [{"$ref": "#/components/schemas/iam_resources_type_object_string"}, {"$ref": "#/components/schemas/iam_resources_type_object_nested"}], "x-auditable": true}
```
