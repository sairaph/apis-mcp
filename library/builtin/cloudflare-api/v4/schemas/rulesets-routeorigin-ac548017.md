---
title: rulesets_RouteOrigin
page_id: schema-rulesets-routeorigin-ac548017
path: schemas
description: An origin to route to.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_RouteOrigin

An origin to route to.

```yaml
{"description": "An origin to route to.", "type": "object", "properties": {"host": {"description": "A resolved host to route to.", "type": "string", "example": "static.example.com", "minLength": 1, "title": "Host"}, "port": {"description": "A destination port to route to.", "type": "integer", "example": 80, "maximum": 65535, "minimum": 1, "title": "Port"}}, "minProperties": 1, "title": "Origin"}
```
