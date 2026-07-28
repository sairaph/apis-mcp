---
title: security-center_valueCountsResponse
page_id: schema-security-center-valuecountsresponse-a390f84a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# security-center_valueCountsResponse

```yaml
{"allOf": [{"$ref": "#/components/schemas/security-center_api-response-common"}, {"properties": {"result": {"anyOf": [{"items": {"properties": {"count": {"type": "integer", "example": 1, "x-auditable": true}, "value": {"type": "string", "x-auditable": true}}, "type": "object"}, "type": "array"}]}}, "type": "object"}]}
```
