---
title: rulesets_RewriteHeaders
page_id: schema-rulesets-rewriteheaders-1a126d19
path: schemas
description: A map of headers to rewrite.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_RewriteHeaders

A map of headers to rewrite.

```yaml
{"description": "A map of headers to rewrite.", "type": "object", "example": {"client-http-version": {"expression": "http.request.version", "operation": "set"}}, "additionalProperties": {"oneOf": [{"description": "A header with a static value to add.", "properties": {"operation": {"allOf": [{"$ref": "#/components/schemas/rulesets_RewriteHeaderOperation"}, {"enum": ["add"]}]}, "value": {"$ref": "#/components/schemas/rulesets_RewriteHeaderValue"}}, "required": ["operation", "value"], "title": "Add Static Header", "type": "object"}, {"description": "A header with a dynamic value to add.", "properties": {"expression": {"$ref": "#/components/schemas/rulesets_RewriteHeaderExpression"}, "operation": {"allOf": [{"$ref": "#/components/schemas/rulesets_RewriteHeaderOperation"}, {"enum": ["add"]}]}}, "required": ["operation", "expression"], "title": "Add Dynamic Header", "type": "object"}, {"description": "A header with a static value to set.", "properties": {"operation": {"allOf": [{"$ref": "#/components/schemas/rulesets_RewriteHeaderOperation"}, {"enum": ["set"]}]}, "value": {"$ref": "#/components/schemas/rulesets_RewriteHeaderValue"}}, "required": ["operation", "value"], "title": "Set Static Header", "type": "object"}, {"description": "A header with a dynamic value to set.", "properties": {"expression": {"$ref": "#/components/schemas/rulesets_RewriteHeaderExpression"}, "operation": {"allOf": [{"$ref": "#/components/schemas/rulesets_RewriteHeaderOperation"}, {"enum": ["set"]}]}}, "required": ["operation", "expression"], "title": "Set Dynamic Header", "type": "object"}, {"description": "A header to remove.", "properties": {"operation": {"allOf": [{"$ref": "#/components/schemas/rulesets_RewriteHeaderOperation"}, {"enum": ["remove"]}]}}, "required": ["operation"], "title": "Remove Header", "type": "object"}]}, "minProperties": 1, "title": "Headers"}
```
