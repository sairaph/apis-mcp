---
title: workers_placement-regions-response
page_id: schema-workers-placement-regions-response-ca0af99d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_placement-regions-response

```yaml
{"type": "object", "properties": {"providers": {"description": "List of cloud providers with their available regions.", "type": "array", "items": {"$ref": "#/components/schemas/workers_placement-provider"}}}, "required": ["providers"]}
```
