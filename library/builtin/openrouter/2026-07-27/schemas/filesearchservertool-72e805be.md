---
title: FileSearchServerTool
page_id: schema-filesearchservertool-72e805be
path: schemas
description: File search tool configuration
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FileSearchServerTool

File search tool configuration

```yaml
{"description": "File search tool configuration", "example": {"type": "file_search", "vector_store_ids": ["vs_abc123"]}, "properties": {"filters": {"anyOf": [{"properties": {"key": {"type": "string"}, "type": {"enum": ["eq", "ne", "gt", "gte", "lt", "lte"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "value": {"anyOf": [{"type": "string"}, {"format": "double", "type": "number"}, {"type": "boolean"}, {"items": {"anyOf": [{"type": "string"}, {"format": "double", "type": "number"}]}, "type": "array"}]}}, "required": ["key", "type", "value"], "type": "object"}, {"$ref": "#/components/schemas/CompoundFilter"}, {"type": "null"}]}, "max_num_results": {"type": "integer"}, "ranking_options": {"properties": {"ranker": {"enum": ["auto", "default-2024-11-15"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "score_threshold": {"format": "double", "type": "number"}}, "type": "object"}, "type": {"enum": ["file_search"], "type": "string"}, "vector_store_ids": {"items": {"type": "string"}, "type": "array"}}, "required": ["type", "vector_store_ids"], "type": "object"}
```
