---
title: workers_placement-provider
page_id: schema-workers-placement-provider-722330d0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_placement-provider

```yaml
{"type": "object", "properties": {"id": {"description": "The cloud provider identifier.", "type": "string", "example": "aws"}, "regions": {"description": "List of regions available for this provider.", "type": "array", "items": {"$ref": "#/components/schemas/workers_placement-region"}}}, "required": ["id", "regions"]}
```
