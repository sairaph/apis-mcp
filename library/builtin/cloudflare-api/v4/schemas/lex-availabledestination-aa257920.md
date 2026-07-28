---
title: lex_AvailableDestination
page_id: schema-lex-availabledestination-aa257920
path: schemas
description: A dataset type that the account or zone can create.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lex_AvailableDestination

A dataset type that the account or zone can create.

```yaml
{"description": "A dataset type that the account or zone can create.", "type": "object", "properties": {"dataset": {"description": "Dataset type name (e.g. `http_requests`).", "type": "string"}, "object_type": {"description": "Whether this dataset type is account-scoped or zone-scoped.", "type": "string", "enum": ["account", "zone"]}, "schema": {"description": "JSON Schema that describes the fields this dataset exposes.", "type": "object", "properties": {"properties": {"type": "object", "additionalProperties": true}, "required": {"type": "array", "items": {"type": "string"}}, "type": {"type": "string", "enum": ["object"]}}}, "timestamp_field": {"description": "The primary timestamp field name for this dataset.", "type": "string"}}, "required": ["dataset", "timestamp_field", "schema", "object_type"]}
```
