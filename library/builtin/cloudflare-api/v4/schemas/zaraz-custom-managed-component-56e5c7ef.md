---
title: zaraz_custom-managed-component
page_id: schema-zaraz-custom-managed-component-56e5c7ef
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_custom-managed-component

```yaml
{"allOf": [{"$ref": "#/components/schemas/zaraz_base-mc"}, {"properties": {"type": {"type": "string", "enum": ["custom-mc"]}, "worker": {"description": "Cloudflare worker that acts as a managed component.", "type": "object", "properties": {"escapedWorkerName": {"type": "string", "x-auditable": true}, "workerTag": {"type": "string", "x-auditable": true}}, "required": ["workerTag", "escapedWorkerName"]}}, "required": ["worker", "type"], "type": "object"}]}
```
