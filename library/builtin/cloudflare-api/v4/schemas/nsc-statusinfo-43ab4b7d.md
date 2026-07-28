---
title: nsc_StatusInfo
page_id: schema-nsc-statusinfo-43ab4b7d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# nsc_StatusInfo

```yaml
{"type": "object", "discriminator": {"propertyName": "state"}, "oneOf": [{"properties": {"state": {"type": "string", "enum": ["Pending"]}}, "required": ["state"], "type": "object"}, {"properties": {"reason": {"description": "Diagnostic information, if available", "type": "string", "nullable": true}, "state": {"type": "string", "enum": ["Down"]}}, "required": ["state"], "type": "object"}, {"properties": {"reason": {"description": "Diagnostic information, if available", "type": "string", "nullable": true}, "state": {"type": "string", "enum": ["Unhealthy"]}}, "required": ["state"], "type": "object"}, {"properties": {"state": {"type": "string", "enum": ["Healthy"]}}, "required": ["state"], "type": "object"}]}
```
