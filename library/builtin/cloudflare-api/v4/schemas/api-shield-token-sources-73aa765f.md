---
title: api-shield_token_sources
page_id: schema-api-shield-token-sources-73aa765f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_token_sources

```yaml
{"type": "array", "items": {"anyOf": [{"$ref": "#/components/schemas/api-shield_header"}, {"$ref": "#/components/schemas/api-shield_cookie"}]}, "example": ["http.request.headers[\"x-auth\"][0]", "http.request.cookies[\"Authorization\"][0]"], "maxItems": 4, "minItems": 1}
```
