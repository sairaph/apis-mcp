---
title: mconn_connector_interrupts_result
page_id: schema-mconn-connector-interrupts-result-13f05e5b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_connector_interrupts_result

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/mconn_interrupt"}, {"properties": {"submitted_at": {"type": "string"}, "triggered_at": {"type": "string"}}, "required": ["submitted_at"], "type": "object"}]}
```
