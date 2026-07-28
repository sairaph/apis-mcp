---
title: AABenchmarkEntry
page_id: schema-aabenchmarkentry-e16b4662
path: schemas
description: Artificial Analysis benchmark index scores.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AABenchmarkEntry

Artificial Analysis benchmark index scores.

```yaml
{"description": "Artificial Analysis benchmark index scores.", "example": {"agentic_index": 55.8, "coding_index": 63.2, "intelligence_index": 71.4}, "properties": {"agentic_index": {"description": "Artificial Analysis Agentic Index score", "example": 55.8, "format": "double", "type": ["number", "null"]}, "coding_index": {"description": "Artificial Analysis Coding Index score", "example": 63.2, "format": "double", "type": ["number", "null"]}, "intelligence_index": {"description": "Artificial Analysis Intelligence Index score", "example": 71.4, "format": "double", "type": ["number", "null"]}}, "required": ["intelligence_index", "coding_index", "agentic_index"], "type": "object"}
```
