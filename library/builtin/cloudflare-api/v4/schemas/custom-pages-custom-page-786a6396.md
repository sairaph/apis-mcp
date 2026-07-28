---
title: custom-pages_custom_page
page_id: schema-custom-pages-custom-page-786a6396
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# custom-pages_custom_page

```yaml
{"type": "object", "properties": {"created_on": {"$ref": "#/components/schemas/custom-pages_timestamp"}, "description": {"type": "string", "example": "Basic Challenge", "x-auditable": true}, "id": {"type": "string", "example": "basic_challenge", "x-auditable": true}, "modified_on": {"$ref": "#/components/schemas/custom-pages_timestamp"}, "preview_target": {"type": "string", "example": "block:basic-sec-captcha", "x-auditable": true}, "required_tokens": {"type": "array", "items": {"type": "string"}, "example": ["::CAPTCHA_BOX::"], "x-auditable": true}, "state": {"$ref": "#/components/schemas/custom-pages_state"}, "url": {"$ref": "#/components/schemas/custom-pages_url"}}}
```
