---
title: OutputFusionServerToolItem
page_id: schema-outputfusionservertoolitem-a8683dc0
path: schemas
description: An openrouter:fusion server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputFusionServerToolItem

An openrouter:fusion server tool output item

```yaml
{"description": "An openrouter:fusion server tool output item", "example": {"id": "st_tmp_abc123", "status": "completed", "type": "openrouter:fusion"}, "properties": {"analysis": {"$ref": "#/components/schemas/FusionAnalysisResult"}, "error": {"description": "Error message when the fusion run did not produce an analysis result.", "type": "string"}, "failed_models": {"description": "Models that were requested as part of the analysis panel but did not produce a response. Present when at least one requested analysis model failed. The fusion result is still usable but was produced from a degraded panel.", "items": {"properties": {"error": {"description": "Error message describing why the model failed.", "type": "string"}, "model": {"description": "Slug of the analysis model that failed.", "type": "string"}, "status_code": {"description": "HTTP status code from the upstream response, when available (e.g. 402, 429).", "type": "integer"}}, "required": ["model", "error"], "type": "object"}, "type": "array"}, "failure_reason": {"description": "Typed failure reason when the fusion run failed. Possible values include: all_panels_failed, insufficient_credits, rate_limited, judge_not_valid_json, judge_schema_mismatch, judge_upstream_error, judge_empty_completion.", "type": "string"}, "id": {"type": "string"}, "responses": {"description": "Analysis models that produced a response in this fusion run, with each model's full panel content.", "items": {"properties": {"content": {"type": "string"}, "model": {"type": "string"}}, "required": ["model"], "type": "object"}, "type": "array"}, "sources": {"description": "Web pages the analysis panels and judge retrieved via web search during this fusion run, deduplicated by URL across the whole run. Present when at least one model cited a source.", "items": {"$ref": "#/components/schemas/FusionSource"}, "type": "array"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["openrouter:fusion"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
