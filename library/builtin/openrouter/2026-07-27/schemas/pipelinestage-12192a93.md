---
title: PipelineStage
page_id: schema-pipelinestage-12192a93
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PipelineStage

```yaml
{"example": {"data": {"action": "redacted", "engines": ["presidio"], "flagged": true, "matched_entity_types": ["EMAIL", "PHONE"]}, "name": "content-filter", "summary": "PII redacted via Presidio (EMAIL, PHONE)", "type": "guardrail"}, "properties": {"cost_usd": {"format": "double", "type": ["number", "null"]}, "data": {"additionalProperties": {}, "type": "object"}, "guardrail_id": {"type": "string"}, "guardrail_scope": {"type": "string"}, "name": {"type": "string"}, "summary": {"type": "string"}, "type": {"$ref": "#/components/schemas/PipelineStageType"}}, "required": ["type", "name"], "type": "object"}
```
