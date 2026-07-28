---
title: zones_forwarding_url
page_id: schema-zones-forwarding-url-f575a4f1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_forwarding_url

```yaml
{"type": "object", "properties": {"id": {"description": "Redirects one URL to another using an `HTTP 301/302` redirect. Refer\nto [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).\n", "type": "string", "example": "forwarding_url", "enum": ["forwarding_url"], "x-auditable": true}, "value": {"properties": {"status_code": {"description": "The status code to use for the URL redirect. 301 is a permanent\nredirect. 302 is a temporary redirect.\n", "enum": [301, 302], "example": "temporary", "type": "integer", "x-auditable": true}, "url": {"description": "The URL to redirect the request to.\nNotes: ${num} refers to the position of '*' in the constraint value.", "example": "http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3", "maxLength": 1500, "type": "string", "x-auditable": true}}}}, "title": "Forwarding URL", "x-stainless-skip": ["terraform"]}
```
