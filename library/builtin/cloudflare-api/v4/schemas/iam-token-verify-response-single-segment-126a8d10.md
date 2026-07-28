---
title: iam_token_verify_response_single_segment
page_id: schema-iam-token-verify-response-single-segment-126a8d10
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_token_verify_response_single_segment

```yaml
{"allOf": [{"$ref": "#/components/schemas/iam_api-response-single"}, {"properties": {"result": {"properties": {"expires_on": {"$ref": "#/components/schemas/iam_expires_on"}, "id": {"$ref": "#/components/schemas/iam_token_identifier"}, "not_before": {"$ref": "#/components/schemas/iam_not_before"}, "status": {"$ref": "#/components/schemas/iam_token_status"}}, "required": ["id", "status"]}}, "type": "object"}]}
```
