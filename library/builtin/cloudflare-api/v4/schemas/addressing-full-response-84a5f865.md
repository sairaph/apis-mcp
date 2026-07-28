---
title: addressing_full_response
page_id: schema-addressing-full-response-84a5f865
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# addressing_full_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/addressing_api-response-single"}, {"properties": {"result": {"allOf": [{"$ref": "#/components/schemas/addressing_address-maps"}, {"properties": {"ips": {"$ref": "#/components/schemas/addressing_ips"}, "memberships": {"$ref": "#/components/schemas/addressing_memberships"}}, "type": "object"}]}}, "type": "object"}]}
```
