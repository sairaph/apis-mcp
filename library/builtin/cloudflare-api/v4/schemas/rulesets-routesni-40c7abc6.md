---
title: rulesets_RouteSNI
page_id: schema-rulesets-routesni-40c7abc6
path: schemas
description: A Server Name Indication (SNI) override.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_RouteSNI

A Server Name Indication (SNI) override.

```yaml
{"description": "A Server Name Indication (SNI) override.", "type": "object", "properties": {"value": {"description": "A value to override the SNI to.", "type": "string", "example": "static.example.com", "minLength": 1, "title": "Value"}}, "required": ["value"], "title": "Server Name Indication"}
```
