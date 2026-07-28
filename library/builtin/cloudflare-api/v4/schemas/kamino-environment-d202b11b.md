---
title: kamino_environment
page_id: schema-kamino-environment-d202b11b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# kamino_environment

```yaml
{"type": "object", "properties": {"expression": {"type": "string"}, "http_application_id": {"type": "string", "nullable": true}, "locked_on_deployment": {"type": "boolean", "nullable": true}, "name": {"type": "string"}, "position": {"$ref": "#/components/schemas/kamino_environment_position"}, "ref": {"type": "string"}, "version": {"type": "integer", "format": "int64", "nullable": true}}, "required": ["name", "ref", "version", "expression", "locked_on_deployment", "position"]}
```
