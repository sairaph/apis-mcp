---
title: api-shield_header
page_id: schema-api-shield-header-1e4f8a74
path: schemas
description: HTTP request header (must be lowercase)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_header

HTTP request header (must be lowercase)

```yaml
{"description": "HTTP request header (must be lowercase)", "type": "string", "example": "http.request.headers[\"x-auth\"][0]", "pattern": "^http.request.headers\\[.*?\\]\\[\\d+\\]$", "x-auditable": true}
```
