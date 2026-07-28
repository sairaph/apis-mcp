---
title: zones_security_header
page_id: schema-zones-security-header-7a17d3ac
path: schemas
description: Cloudflare security header for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_security_header

Cloudflare security header for a zone.

```yaml
{"description": "Cloudflare security header for a zone.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone's security header.", "example": "security_header", "enum": ["security_header"]}, "value": {"$ref": "#/components/schemas/zones_security_header_value"}}}], "title": "Security Header"}
```
