---
title: zones_cache_on_cookie
page_id: schema-zones-cache-on-cookie-f9621d13
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_cache_on_cookie

```yaml
{"type": "object", "properties": {"id": {"description": "Apply the Cache Everything option (Cache Level setting) based on a\nregular expression match against a cookie name.\n", "type": "string", "enum": ["cache_on_cookie"], "x-auditable": true}, "value": {"description": "The regular expression to use for matching cookie names in the\nrequest.\n", "type": "string", "example": "bypass=.*|PHPSESSID=.*", "maxLength": 150, "minLength": 1, "x-auditable": true}}, "title": "Cache on Cookie", "x-stainless-skip": ["terraform"]}
```
