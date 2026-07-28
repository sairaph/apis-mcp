---
title: workers_tail
page_id: schema-workers-tail-2650cdee
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_tail

```yaml
{"type": "object", "properties": {"expires_at": {"type": "string", "readOnly": true, "x-auditable": true}, "id": {"allOf": [{"$ref": "#/components/schemas/workers_identifier"}], "readOnly": true}, "url": {"type": "string", "readOnly": true, "x-auditable": true}}, "required": ["id", "url", "expires_at"]}
```
