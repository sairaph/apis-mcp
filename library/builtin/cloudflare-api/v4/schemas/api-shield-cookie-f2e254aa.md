---
title: api-shield_cookie
page_id: schema-api-shield-cookie-f2e254aa
path: schemas
description: HTTP request cookie
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_cookie

HTTP request cookie

```yaml
{"description": "HTTP request cookie", "type": "string", "example": "http.request.cookies[\"Authorization\"][0]", "pattern": "^http.request.cookies\\[.*?\\]\\[\\d+\\]$", "x-auditable": true}
```
