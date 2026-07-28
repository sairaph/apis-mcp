---
title: page-shield_cookie
page_id: schema-page-shield-cookie-7c2d6334
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# page-shield_cookie

```yaml
{"properties": {"domain_attribute": {"type": "string", "example": "cloudflare.com"}, "expires_attribute": {"type": "string", "format": "date-time", "example": "2021-10-02T09:57:54Z"}, "first_seen_at": {"type": "string", "format": "date-time", "example": "2021-08-18T10:51:08Z"}, "host": {"type": "string", "example": "blog.cloudflare.com"}, "http_only_attribute": {"type": "boolean", "example": true}, "id": {"$ref": "#/components/schemas/page-shield_id"}, "last_seen_at": {"type": "string", "format": "date-time", "example": "2021-09-02T09:57:54Z"}, "max_age_attribute": {"type": "integer", "example": 3600}, "name": {"type": "string", "example": "session_id"}, "page_urls": {"type": "array", "items": {"type": "string"}, "example": ["blog.cloudflare.com/page1", "blog.cloudflare.com/page2"]}, "path_attribute": {"type": "string", "example": "/"}, "same_site_attribute": {"type": "string", "example": "strict", "enum": ["lax", "strict", "none"]}, "secure_attribute": {"type": "boolean", "example": true}, "type": {"type": "string", "example": "first_party", "enum": ["first_party", "unknown"]}}, "required": ["id", "type", "name", "host", "first_seen_at", "last_seen_at"]}
```
