---
title: pages_pages_assets_check_missing_response
page_id: schema-pages-pages-assets-check-missing-response-fe559d4a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# pages_pages_assets_check_missing_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"description": "List of file content hashes that are missing from the asset store and need to be uploaded.", "type": "array", "items": {"type": "string"}, "example": ["b026324c6904b2a9cb4b88d6d61c81d1"]}}, "required": ["result"], "type": "object"}]}
```
