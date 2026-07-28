---
title: ReasoningDetailSummary
page_id: schema-reasoningdetailsummary-0ba780cc
path: schemas
description: Reasoning detail summary schema
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ReasoningDetailSummary

Reasoning detail summary schema

```yaml
{"description": "Reasoning detail summary schema", "example": {"summary": "The model analyzed the problem by first identifying key constraints, then evaluating possible solutions...", "type": "reasoning.summary"}, "properties": {"format": {"$ref": "#/components/schemas/ReasoningFormat"}, "id": {"type": ["string", "null"]}, "index": {"type": "integer"}, "summary": {"type": "string"}, "type": {"enum": ["reasoning.summary"], "type": "string"}}, "required": ["type", "summary"], "type": "object"}
```
