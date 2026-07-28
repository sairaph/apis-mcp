---
title: nsc_InterconnectCreate_Physical_Body
page_id: schema-nsc-interconnectcreate-physical-body-3e357200
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# nsc_InterconnectCreate_Physical_Body

```yaml
{"type": "object", "allOf": [{"properties": {"account": {"type": "string"}, "type": {"type": "string"}}, "required": ["type", "account"], "type": "object"}, {"properties": {"slot_id": {"type": "string", "format": "uuid"}, "speed": {"type": "string", "nullable": true}}, "required": ["slot_id"], "type": "object"}], "title": "Physical"}
```
