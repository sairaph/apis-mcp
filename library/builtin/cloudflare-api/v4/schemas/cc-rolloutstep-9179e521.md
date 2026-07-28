---
title: cc_RolloutStep
page_id: schema-cc-rolloutstep-9179e521
path: schemas
description: Steps within the rollout process.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_RolloutStep

Steps within the rollout process.

```yaml
{"description": "Steps within the rollout process.", "type": "object", "properties": {"completed_at": {"$ref": "#/components/schemas/cc_ISO8601Timestamp"}, "description": {"description": "Description of the rollout step.", "type": "string"}, "id": {"description": "The sequential order of the rollout step, automatically assigned starting from 1, based on the total number of steps in the rollout process.", "type": "integer"}, "reason": {"description": "Reason why the step has the current status", "type": "string"}, "started_at": {"$ref": "#/components/schemas/cc_ISO8601Timestamp"}, "status": {"description": "Status of the rollout step.", "type": "string", "enum": ["pending", "progressing", "reverting", "completed", "reverted"]}, "step_size": {"type": "object", "properties": {"percentage": {"description": "Percentage of instances affected in this step. Min 10% and Max 100%.", "type": "integer"}}, "required": ["percentage"]}}, "required": ["id", "step_size", "description", "status"]}
```
