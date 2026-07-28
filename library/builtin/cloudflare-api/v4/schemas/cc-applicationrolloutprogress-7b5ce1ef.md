---
title: cc_ApplicationRolloutProgress
page_id: schema-cc-applicationrolloutprogress-7b5ce1ef
path: schemas
description: Progress details of an application rollout.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ApplicationRolloutProgress

Progress details of an application rollout.

```yaml
{"description": "Progress details of an application rollout.", "type": "object", "properties": {"current_step": {"description": "Current step being executed in the rollout process. Initialized to 0.", "type": "integer"}, "total_instances": {"description": "Total number of instances affected by the rollout.", "type": "integer"}, "total_steps": {"description": "Total number of steps in the rollout.", "type": "integer"}, "updated_instances": {"description": "Number of instances updated in the rollout process.", "type": "integer"}, "version_distribution": {"description": "Expected distribution of instances by version based on the current percentage split.\nPopulated during active rollouts. Values are computed from the version percentage weights,\nnot actual running instance counts.\n", "type": "object", "properties": {"current_version_instances": {"description": "Expected number of instances remaining on the current (old) version based on the current percentage split. Only populated for \"rolling\" strategy.", "type": "integer"}, "current_version_percentage": {"description": "The percentage of new instances being scheduled on the current version (100 - target_version_percentage).\nOnly populated for \"new_instances\" strategy.\n", "type": "integer"}, "target_version_instances": {"description": "Expected number of instances scheduled for the target (new) version based on the current percentage split. Only populated for \"rolling\" strategy.", "type": "integer"}, "target_version_percentage": {"description": "The active percentage of new instances being scheduled on the target version.\nFor \"rolling\", this reflects the step_size.percentage of the current active step.\nFor \"new_instances\", this reflects the user-set percentage.\n", "type": "integer"}}}}, "required": ["total_steps", "current_step", "updated_instances", "total_instances"]}
```
