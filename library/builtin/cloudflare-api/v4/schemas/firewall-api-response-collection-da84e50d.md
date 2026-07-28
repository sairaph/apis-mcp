---
title: firewall_api-response-collection
page_id: schema-firewall-api-response-collection-da84e50d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_api-response-collection

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/firewall_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"type": "object"}, "nullable": true}, "result_info": {"$ref": "#/components/schemas/firewall_result_info"}}, "type": "object"}]}
```
