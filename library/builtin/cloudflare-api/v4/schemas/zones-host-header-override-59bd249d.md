---
title: zones_host_header_override
page_id: schema-zones-host-header-override-59bd249d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_host_header_override

```yaml
{"type": "object", "properties": {"id": {"description": "Apply a specific host header.", "type": "string", "enum": ["host_header_override"], "x-auditable": true}, "value": {"description": "The hostname to use in the `Host` header", "type": "string", "example": "example.com", "minLength": 1, "x-auditable": true}}, "title": "Host Header Override", "x-stainless-skip": ["terraform"]}
```
