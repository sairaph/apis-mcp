---
title: dos_ExpressionFilter
page_id: schema-dos-expressionfilter-6191f39e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dos_ExpressionFilter

```yaml
{"type": "object", "properties": {"created_on": {"description": "The creation timestamp of the expression filter.", "type": "string", "format": "date-time", "x-auditable": true}, "expression": {"description": "The filter expression.", "type": "string", "example": "ip.dst in { 192.0.2.0/24 198.51.100.0/24 } and tcp.srcport in { 80 443 10000..65535 }", "x-auditable": true}, "id": {"description": "The unique ID of the expression filter.", "type": "string", "x-auditable": true}, "mode": {"description": "The filter's mode. Must be one of 'enabled', 'disabled', 'monitoring'.", "type": "string", "x-auditable": true}, "modified_on": {"description": "The last modification timestamp of the expression filter.", "type": "string", "format": "date-time", "x-auditable": true}}, "additionalProperties": false, "required": ["id", "expression", "mode", "created_on", "modified_on"]}
```
