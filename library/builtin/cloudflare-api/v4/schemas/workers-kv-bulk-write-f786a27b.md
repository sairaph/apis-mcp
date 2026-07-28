---
title: workers-kv_bulk_write
page_id: schema-workers-kv-bulk-write-f786a27b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers-kv_bulk_write

```yaml
{"type": "array", "items": {"properties": {"base64": {"description": "Indicates whether or not the server should base64 decode the value before storing it. Useful for writing values that wouldn't otherwise be valid JSON strings, such as images.", "type": "boolean", "default": false}, "expiration": {"$ref": "#/components/schemas/workers-kv_expiration"}, "expiration_ttl": {"$ref": "#/components/schemas/workers-kv_expiration_ttl"}, "key": {"$ref": "#/components/schemas/workers-kv_key_name_bulk"}, "metadata": {"$ref": "#/components/schemas/workers-kv_list_metadata"}, "value": {"description": "A UTF-8 encoded string to be stored, up to 25 MiB in length.", "type": "string", "example": "Some string", "maxLength": 26214400}}, "required": ["key", "value"], "type": "object"}}
```
