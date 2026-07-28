---
title: rulesets_CompressResponseRule
page_id: schema-rulesets-compressresponserule-5f2b6a9d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_CompressResponseRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["compress_response"]}, "action_parameters": {"properties": {"algorithms": {"description": "Custom order for compression algorithms.", "type": "array", "items": {"description": "Compression algorithm to enable.", "properties": {"name": {"description": "Name of the compression algorithm to enable.", "type": "string", "example": "none", "enum": ["none", "auto", "default", "gzip", "brotli", "zstd"], "title": "Algorithm Name"}}, "title": "Algorithm", "type": "object"}, "minItems": 1, "title": "Algorithms", "uniqueItems": true}}, "required": ["algorithms"]}, "description": {"example": "Modify the compression algorithm used in the response."}}, "title": "Response Compression Rule"}]}
```
