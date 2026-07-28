---
title: workers-kv_key
page_id: schema-workers-kv-key-8089fe59
path: schemas
description: A name for a value. A value stored under a given key may be retrieved via the same key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers-kv_key

A name for a value. A value stored under a given key may be retrieved via the same key.

```yaml
{"description": "A name for a value. A value stored under a given key may be retrieved via the same key.", "type": "object", "properties": {"expiration": {"description": "The time, measured in number of seconds since the UNIX epoch, at which the key will expire. This property is omitted for keys that will not expire.", "type": "number", "example": 1577836800}, "metadata": {"$ref": "#/components/schemas/workers-kv_list_metadata"}, "name": {"$ref": "#/components/schemas/workers-kv_key_name"}}, "required": ["name"]}
```
