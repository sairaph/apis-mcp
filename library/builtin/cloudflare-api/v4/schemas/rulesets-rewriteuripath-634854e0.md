---
title: rulesets_RewriteUriPath
page_id: schema-rulesets-rewriteuripath-634854e0
path: schemas
description: A URI path rewrite.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_RewriteUriPath

A URI path rewrite.

```yaml
{"description": "A URI path rewrite.", "type": "object", "properties": {"expression": {"description": "An expression that evaluates to a value to rewrite the URI path to.", "type": "string", "example": "regex_replace(http.request.uri.path, \"/foo$\", \"/bar\")", "minLength": 1, "title": "Path Expression"}, "value": {"description": "A value to rewrite the URI path to.", "type": "string", "example": "/foo", "minLength": 1, "title": "Path Value"}}, "maxProperties": 1, "minProperties": 1, "title": "URI Path"}
```
