---
title: zones_url_target
page_id: schema-zones-url-target-77555cc1
path: schemas
description: URL target.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_url_target

URL target.

```yaml
{"description": "URL target.", "type": "object", "properties": {"constraint": {"description": "The constraint of a target.", "type": "object", "allOf": [{"$ref": "#/components/schemas/zones_string_constraint"}, {"properties": {"value": {"description": "The URL pattern to match against the current request. The pattern may contain up to four asterisks ('*') as placeholders.", "type": "string", "example": "*example.com/images/*", "pattern": "^(https?://)?(([-a-zA-Z0-9*]*\\.)+[-a-zA-Z0-9]{2,20})(:(8080|8443|443|80))?(/[\\S]+)?$", "x-auditable": true}}}]}, "target": {"description": "A target based on the URL of the request.", "example": "url", "enum": ["url"], "x-auditable": true}}}
```
