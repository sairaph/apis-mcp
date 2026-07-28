---
title: rulesets_RewriteUriQuery
page_id: schema-rulesets-rewriteuriquery-9151f3db
path: schemas
description: A URI query rewrite.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_RewriteUriQuery

A URI query rewrite.

```yaml
{"description": "A URI query rewrite.", "type": "object", "properties": {"expression": {"description": "An expression that evaluates to a value to rewrite the URI query to.", "type": "string", "example": "regex_replace(http.request.uri.query, \"foo=bar\", \"\")", "minLength": 1, "title": "Query Expression"}, "value": {"description": "A value to rewrite the URI query to.", "type": "string", "example": "foo=bar", "title": "Query Value"}}, "maxProperties": 1, "minProperties": 1, "title": "URI Query"}
```
