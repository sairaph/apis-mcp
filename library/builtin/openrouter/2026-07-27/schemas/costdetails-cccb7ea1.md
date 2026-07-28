---
title: CostDetails
page_id: schema-costdetails-cccb7ea1
path: schemas
description: Breakdown of upstream inference costs
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# CostDetails

Breakdown of upstream inference costs

```yaml
{"description": "Breakdown of upstream inference costs", "example": {"upstream_inference_completions_cost": 0.0004, "upstream_inference_cost": null, "upstream_inference_prompt_cost": 0.0008}, "properties": {"upstream_inference_completions_cost": {"format": "double", "type": "number"}, "upstream_inference_cost": {"format": "double", "type": ["number", "null"]}, "upstream_inference_prompt_cost": {"format": "double", "type": "number"}}, "required": ["upstream_inference_prompt_cost", "upstream_inference_completions_cost"], "type": ["object", "null"]}
```
