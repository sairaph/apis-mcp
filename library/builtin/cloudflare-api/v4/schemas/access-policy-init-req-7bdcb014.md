---
title: access_policy_init_req
page_id: schema-access-policy-init-req-7bdcb014
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_policy_init_req

```yaml
{"type": "object", "properties": {"policies": {"type": "array", "items": {"oneOf": [{"$ref": "#/components/schemas/access_policy_req"}, {"description": "The UUID of the reusable policy you wish to test", "example": "f1a8b3c9d4e5f6789a0b1c2d3e4f5678a9b0c1d2e3f4a5b67890c1d2e3f4b5a6", "type": "string"}]}}}}
```
