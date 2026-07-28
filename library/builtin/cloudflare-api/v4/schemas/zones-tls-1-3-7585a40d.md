---
title: zones_tls_1_3
page_id: schema-zones-tls-1-3-7585a40d
path: schemas
description: Enables Crypto TLS 1.3 feature for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_tls_1_3

Enables Crypto TLS 1.3 feature for a zone.

```yaml
{"description": "Enables Crypto TLS 1.3 feature for a zone.", "default": "off", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "tls_1_3", "enum": ["tls_1_3"]}, "value": {"$ref": "#/components/schemas/zones_tls_1_3_value"}}}], "title": "Enable TLS 1.3 value for a zone"}
```
