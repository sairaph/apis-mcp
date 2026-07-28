---
title: zones_bypass_cache_on_cookie
page_id: schema-zones-bypass-cache-on-cookie-b503ccba
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_bypass_cache_on_cookie

```yaml
{"type": "object", "properties": {"id": {"description": "Bypass cache and fetch resources from the origin server if a regular\nexpression matches against a cookie name present in the request.\n", "type": "string", "enum": ["bypass_cache_on_cookie"], "x-auditable": true}, "value": {"description": "The regular expression to use for matching cookie names in the\nrequest. Refer to [Bypass Cache on Cookie\nsetting](https://developers.cloudflare.com/rules/page-rules/reference/additional-reference/#bypass-cache-on-cookie-setting)\nto learn about limited regular expression support.\n", "type": "string", "example": "bypass=.*|PHPSESSID=.*", "maxLength": 150, "minLength": 1, "x-auditable": true}}, "title": "Bypass Cache on Cookie", "x-stainless-skip": ["terraform"]}
```
