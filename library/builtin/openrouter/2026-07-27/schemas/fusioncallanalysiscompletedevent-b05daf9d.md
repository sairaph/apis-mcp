---
title: FusionCallAnalysisCompletedEvent
page_id: schema-fusioncallanalysiscompletedevent-b05daf9d
path: schemas
description: Emitted when the fusion judge completes with the structured analysis.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FusionCallAnalysisCompletedEvent

Emitted when the fusion judge completes with the structured analysis.

```yaml
{"description": "Emitted when the fusion judge completes with the structured analysis.", "example": {"analysis": {"blind_spots": [], "consensus": [], "contradictions": [], "partial_coverage": [], "unique_insights": []}, "item_id": "st_fusion_abc", "output_index": 0, "sequence_number": 40, "type": "response.fusion_call.analysis.completed"}, "properties": {"analysis": {"$ref": "#/components/schemas/FusionAnalysisResult"}, "item_id": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.fusion_call.analysis.completed"], "type": "string"}}, "required": ["type", "analysis", "output_index", "item_id", "sequence_number"], "type": "object"}
```
