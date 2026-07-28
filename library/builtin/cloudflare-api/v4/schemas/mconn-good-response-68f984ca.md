---
title: mconn_good_response
page_id: schema-mconn-good-response-68f984ca
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_good_response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/mconn_response"}, {"properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/mconn_coded_message"}, "maxLength": 0}}, "type": "object"}]}
```
