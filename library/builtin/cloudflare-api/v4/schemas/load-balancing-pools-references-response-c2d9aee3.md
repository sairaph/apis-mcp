---
title: load-balancing_pools-references-response
page_id: schema-load-balancing-pools-references-response-c2d9aee3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_pools-references-response

```yaml
{"allOf": [{"$ref": "#/components/schemas/load-balancing_api-response-common"}, {"properties": {"result": {"description": "List of resources that reference a given pool.", "type": "array", "items": {"properties": {"reference_type": {"type": "string", "enum": ["*", "referral", "referrer"], "x-auditable": true}, "resource_id": {"type": "string", "x-auditable": true}, "resource_name": {"type": "string", "x-auditable": true}, "resource_type": {"type": "string", "x-auditable": true}}, "type": "object"}, "example": [{"reference_type": "referrer", "resource_id": "699d98642c564d2e855e9661899b7252", "resource_name": "www.example.com", "resource_type": "load_balancer"}, {"reference_type": "referral", "resource_id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "resource_name": "Login page monitor", "resource_type": "monitor"}]}}, "type": "object"}]}
```
