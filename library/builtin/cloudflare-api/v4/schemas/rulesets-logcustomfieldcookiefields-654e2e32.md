---
title: rulesets_LogCustomFieldCookieFields
page_id: schema-rulesets-logcustomfieldcookiefields-654e2e32
path: schemas
description: The cookie fields to log.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_LogCustomFieldCookieFields

The cookie fields to log.

```yaml
{"description": "The cookie fields to log.", "type": "array", "items": {"description": "The cookie field to log.", "properties": {"name": {"description": "The name of the cookie.", "type": "string", "example": "myCookie", "minLength": 1, "title": "Cookie Name"}}, "required": ["name"], "title": "Cookie Field", "type": "object"}, "minItems": 1, "title": "Cookie Fields"}
```
