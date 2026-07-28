---
title: workers-kv_namespace
page_id: schema-workers-kv-namespace-a0fd42ec
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers-kv_namespace

```yaml
{"type": "object", "properties": {"id": {"$ref": "#/components/schemas/workers-kv_namespace_identifier"}, "supports_url_encoding": {"description": "True if keys written on the URL will be URL-decoded before storing. For example, if set to \"true\", a key written on the URL as \"%3F\" will be stored as \"?\".", "type": "boolean", "example": true, "readOnly": true, "x-auditable": true}, "title": {"$ref": "#/components/schemas/workers-kv_namespace_title"}}, "required": ["id", "title"]}
```
