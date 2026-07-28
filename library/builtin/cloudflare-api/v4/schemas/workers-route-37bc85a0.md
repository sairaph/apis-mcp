---
title: workers_route
page_id: schema-workers-route-37bc85a0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_route

```yaml
{"type": "object", "properties": {"id": {"allOf": [{"$ref": "#/components/schemas/workers_identifier"}], "readOnly": true}, "pattern": {"description": "Pattern to match incoming requests against. [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).", "type": "string", "example": "example.com/*", "x-auditable": true}, "script": {"description": "Name of the script to run if the route matches.", "type": "string", "example": "my-workers-script", "x-auditable": true}}, "required": ["id", "pattern"]}
```
