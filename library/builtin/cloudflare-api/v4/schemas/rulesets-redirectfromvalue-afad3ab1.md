---
title: rulesets_RedirectFromValue
page_id: schema-rulesets-redirectfromvalue-afad3ab1
path: schemas
description: A redirect based on the request properties.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_RedirectFromValue

A redirect based on the request properties.

```yaml
{"description": "A redirect based on the request properties.", "type": "object", "properties": {"preserve_query_string": {"description": "Whether to keep the query string of the original request.", "type": "boolean", "example": true, "default": false, "title": "Preserve Query String"}, "status_code": {"description": "The status code to use for the redirect.", "type": "integer", "example": 302, "enum": [301, 302, 303, 307, 308], "title": "Status Code"}, "target_url": {"description": "A URL to redirect the request to.", "type": "object", "maxProperties": 1, "minProperties": 1, "properties": {"expression": {"description": "An expression that evaluates to a URL to redirect the request to.", "type": "string", "example": "concat(\"https://example.com\", http.request.uri.path)", "minLength": 1, "title": "Redirect Expression"}, "value": {"description": "A URL to redirect the request to.", "type": "string", "example": "https://example.com", "minLength": 1, "title": "Redirect Value"}}, "title": "Target URL"}}, "required": ["target_url"], "title": "Single Redirect"}
```
