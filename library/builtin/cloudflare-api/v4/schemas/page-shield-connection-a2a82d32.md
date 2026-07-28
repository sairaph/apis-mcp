---
title: page-shield_connection
page_id: schema-page-shield-connection-a2a82d32
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# page-shield_connection

```yaml
{"properties": {"added_at": {"type": "string", "format": "date-time", "example": "2021-08-18T10:51:10.09615Z"}, "domain_reported_malicious": {"type": "boolean", "example": false}, "first_page_url": {"type": "string", "example": "blog.cloudflare.com/page"}, "first_seen_at": {"type": "string", "format": "date-time", "example": "2021-08-18T10:51:08Z"}, "host": {"type": "string", "example": "blog.cloudflare.com"}, "id": {"$ref": "#/components/schemas/page-shield_id"}, "last_seen_at": {"type": "string", "format": "date-time", "example": "2021-09-02T09:57:54Z"}, "malicious_domain_categories": {"type": "array", "items": {"type": "string"}, "example": ["Malware"]}, "malicious_url_categories": {"type": "array", "items": {"type": "string"}, "example": ["Malware"]}, "page_urls": {"type": "array", "items": {"type": "string"}, "example": ["blog.cloudflare.com/page1", "blog.cloudflare.com/page2"]}, "url": {"type": "string", "example": "https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.6.0/js/bootstrap.min.js"}, "url_contains_cdn_cgi_path": {"type": "boolean", "example": false}, "url_reported_malicious": {"type": "boolean", "example": false}}, "required": ["id", "url", "added_at", "first_seen_at", "last_seen_at", "host", "url_contains_cdn_cgi_path"]}
```
