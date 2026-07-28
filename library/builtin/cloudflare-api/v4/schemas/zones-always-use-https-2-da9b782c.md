---
title: zones_always_use_https-2
page_id: schema-zones-always-use-https-2-da9b782c
path: schemas
description: Reply to all requests for URLs that use "http" with a 301 redirect to the equivalent "https" URL. If you only want to redirect for a subset of requests, consider creating an "Always use HTTPS" page rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_always_use_https-2

Reply to all requests for URLs that use "http" with a 301 redirect to the equivalent "https" URL. If you only want to redirect for a subset of requests, consider creating an "Always use HTTPS" page rule.

```yaml
{"description": "Reply to all requests for URLs that use \"http\" with a 301 redirect to the equivalent \"https\" URL. If you only want to redirect for a subset of requests, consider creating an \"Always use HTTPS\" page rule.", "default": "off", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "always_use_https", "enum": ["always_use_https"]}, "value": {"$ref": "#/components/schemas/zones_always_use_https_value"}}}], "title": "Zone Enable Always Use HTTPS"}
```
