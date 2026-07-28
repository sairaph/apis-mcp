---
title: rulesets_RedirectFromList
page_id: schema-rulesets-redirectfromlist-9ff53cdc
path: schemas
description: A redirect based on a bulk list lookup.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_RedirectFromList

A redirect based on a bulk list lookup.

```yaml
{"description": "A redirect based on a bulk list lookup.", "type": "object", "properties": {"key": {"description": "An expression that evaluates to the list lookup key.", "type": "string", "example": "http.request.full_uri", "minLength": 1, "title": "Lookup Key"}, "name": {"description": "The name of the list to match against.", "type": "string", "example": "my_list", "pattern": "^[a-zA-Z0-9_]+$", "title": "List Name"}}, "required": ["name", "key"], "title": "Bulk Redirect"}
```
