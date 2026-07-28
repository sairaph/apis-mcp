---
title: FusionCallAnalysisInProgressEvent
page_id: schema-fusioncallanalysisinprogressevent-11b5161b
path: schemas
description: Emitted when the fusion judge starts producing the structured analysis.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FusionCallAnalysisInProgressEvent

Emitted when the fusion judge starts producing the structured analysis.

```yaml
{"description": "Emitted when the fusion judge starts producing the structured analysis.", "example": {"item_id": "st_fusion_abc", "judge_model": "openai/gpt-5", "output_index": 0, "sequence_number": 25, "type": "response.fusion_call.analysis.in_progress"}, "properties": {"item_id": {"type": "string"}, "judge_model": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.fusion_call.analysis.in_progress"], "type": "string"}}, "required": ["type", "judge_model", "output_index", "item_id", "sequence_number"], "type": "object"}
```
