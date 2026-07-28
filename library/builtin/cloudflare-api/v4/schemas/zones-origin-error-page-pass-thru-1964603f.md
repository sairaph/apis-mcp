---
title: zones_origin_error_page_pass_thru
page_id: schema-zones-origin-error-page-pass-thru-1964603f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_origin_error_page_pass_thru

```yaml
{"type": "object", "properties": {"id": {"description": "Turn on or off Cloudflare error pages generated from issues sent from the origin server. If enabled, this setting triggers error pages issued by the origin.\n", "type": "string", "example": "origin_error_page_pass_thru", "enum": ["origin_error_page_pass_thru"], "x-auditable": true}, "value": {"description": "The status of Origin Error Page Passthru.\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "title": "Origin Error Page Passthru", "x-stainless-skip": ["terraform"]}
```
