---
title: dos_NewExpressionFilter
page_id: schema-dos-newexpressionfilter-6bf3d604
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dos_NewExpressionFilter

```yaml
{"type": "object", "properties": {"expression": {"description": "The filter expression.", "type": "string", "example": "ip.dst in { 192.0.2.0/24 198.51.100.0/24 } and tcp.srcport in { 80 443 10000..65535 }", "x-auditable": true}, "mode": {"description": "The filter's mode. Must be one of 'enabled', 'disabled', 'monitoring'.", "type": "string", "x-auditable": true}}, "additionalProperties": false, "required": ["expression", "mode"]}
```
