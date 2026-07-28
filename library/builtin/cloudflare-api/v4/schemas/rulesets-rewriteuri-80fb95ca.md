---
title: rulesets_RewriteUri
page_id: schema-rulesets-rewriteuri-80fb95ca
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_RewriteUri

```yaml
{"allOf": [{"anyOf": [{"description": "A URI path rewrite.", "properties": {"path": {"$ref": "#/components/schemas/rulesets_RewriteUriPath"}}, "required": ["path"], "title": "URI Path"}, {"description": "A URI query rewrite.", "properties": {"query": {"$ref": "#/components/schemas/rulesets_RewriteUriQuery"}}, "required": ["query"], "title": "URI Query"}]}, {"properties": {"origin": {"description": "Whether to propagate the rewritten URI to origin.", "type": "boolean", "example": false, "readOnly": true, "title": "Origin"}}, "type": "object"}]}
```
