---
title: cc_RolloutStepRequest
page_id: schema-cc-rolloutsteprequest-a403722c
path: schemas
description: Steps defining the rollout process.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_RolloutStepRequest

Steps defining the rollout process.

```yaml
{"description": "Steps defining the rollout process.", "type": "object", "properties": {"description": {"description": "Description of the rollout step.", "type": "string"}, "step_size": {"type": "object", "properties": {"percentage": {"description": "Percentage of instances affected in this step. Min 10% and Max 100%.", "type": "integer"}}, "required": ["percentage"]}}, "required": ["step_size", "description"]}
```
