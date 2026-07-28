---
title: zones_tls_1_2_only
page_id: schema-zones-tls-1-2-only-18365927
path: schemas
description: Only allows TLS1.2.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_tls_1_2_only

Only allows TLS1.2.

```yaml
{"description": "Only allows TLS1.2.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "Zone setting identifier.", "example": "tls_1_2_only", "enum": ["tls_1_2_only"]}, "value": {"$ref": "#/components/schemas/zones_tls_1_2_only_value"}}}], "title": "TLS1.2 only"}
```
