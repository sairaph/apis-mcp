---
title: TaskClassificationModel
page_id: schema-taskclassificationmodel-033a257a
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# TaskClassificationModel

```yaml
{"example": {"id": "openai/gpt-4.1-mini", "tag_token_share": 0.75, "tag_usage_share": 0.55}, "properties": {"id": {"description": "Model identifier (permaslug).", "example": "openai/gpt-4.1-mini", "type": "string"}, "tag_token_share": {"description": "Fraction of this classification's sampled token volume attributed to this model (0–1). Sums to ≤1 across the returned models (only top-N are included and unattributed requests are excluded).", "example": 0.75, "format": "double", "type": "number"}, "tag_usage_share": {"description": "Fraction of this classification's sampled requests attributed to this model (0–1). Sums to ≤1 across the returned models (only top-N are included and unattributed requests are excluded).", "example": 0.55, "format": "double", "type": "number"}}, "required": ["id", "tag_usage_share", "tag_token_share"], "type": "object"}
```
