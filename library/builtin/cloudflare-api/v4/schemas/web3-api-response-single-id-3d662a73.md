---
title: web3_api-response-single-id
page_id: schema-web3-api-response-single-id-3d662a73
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# web3_api-response-single-id

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/web3_api-response-common"}, {"properties": {"result": {"type": "object", "nullable": true, "properties": {"id": {"$ref": "#/components/schemas/web3_identifier"}}, "required": ["id"]}}}]}
```
