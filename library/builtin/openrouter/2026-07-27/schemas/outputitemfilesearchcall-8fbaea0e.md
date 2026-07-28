---
title: OutputItemFileSearchCall
page_id: schema-outputitemfilesearchcall-8fbaea0e
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputItemFileSearchCall

```yaml
{"example": {"id": "filesearch-abc123", "queries": ["machine learning algorithms", "neural networks"], "status": "completed", "type": "file_search_call"}, "properties": {"id": {"type": "string"}, "queries": {"items": {"type": "string"}, "type": "array"}, "status": {"$ref": "#/components/schemas/WebSearchStatus"}, "type": {"enum": ["file_search_call"], "type": "string"}}, "required": ["type", "id", "queries", "status"], "type": "object"}
```
