---
title: cc_ApplicationRollout
page_id: schema-cc-applicationrollout-4ebd8b41
path: schemas
description: |-
    Represents the status and metadata of a rollout process for an application.
    For "rolling" strategy: includes steps and progress with instance counts.
    For "new_instances" strategy: steps and progress are omitted. Use percentage,
    version_distribution, and health.summary for status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ApplicationRollout

Represents the status and metadata of a rollout process for an application.
For "rolling" strategy: includes steps and progress with instance counts.
For "new_instances" strategy: steps and progress are omitted. Use percentage,
version_distribution, and health.summary for status.

```yaml
{"description": "Represents the status and metadata of a rollout process for an application.\nFor \"rolling\" strategy: includes steps and progress with instance counts.\nFor \"new_instances\" strategy: steps and progress are omitted. Use percentage,\nversion_distribution, and health.summary for status.\n", "type": "object", "properties": {"created_at": {"$ref": "#/components/schemas/cc_ISO8601Timestamp"}, "current_configuration": {"$ref": "#/components/schemas/cc_ModifyUserDeploymentConfiguration"}, "current_version": {"description": "Current application version before the rollout.", "type": "integer"}, "description": {"type": "string"}, "health": {"$ref": "#/components/schemas/cc_ApplicationHealth"}, "id": {"$ref": "#/components/schemas/cc_RolloutID"}, "kind": {"description": "Kind of the rollout process.\n - \"full_auto\": For rolling rollouts, starts progressing steps upon rollout creation. For new_instances rollouts, advances percentage targets automatically after target-version health is observed.\n - \"full_manual\": Requires manually progressing each step in the rollout using the UpdateRollout's action paramater.\n - \"durable_objects_auto\": Default when the application is a DO application.\n", "type": "string", "enum": ["full_auto", "full_manual", "durable_objects_auto"]}, "last_updated_at": {"$ref": "#/components/schemas/cc_ISO8601Timestamp"}, "percentage": {"description": "Current target version percentage (0-100). Only present for \"new_instances\" strategy.\n", "type": "integer"}, "progress": {"$ref": "#/components/schemas/cc_ApplicationRolloutProgress"}, "started_at": {"description": "Timestamp when the rollout started.", "type": "string", "format": "date-time"}, "status": {"description": "Current status of the rollout.", "type": "string", "enum": ["pending", "progressing", "completed", "reverted", "replaced"]}, "steps": {"type": "array", "items": {"$ref": "#/components/schemas/cc_RolloutStep"}}, "strategy": {"description": "The rollout strategy.\n- \"rolling\": Step-based rollout with health gates. Actively replaces instances to reach each step's target percentage. Response includes steps and progress.\n- \"new_instances\": Percentage control over version distribution. Version sync actively replaces instances to match the configured percentage. \"full_auto\" ramps through fixed percentage targets after target-version health is observed. Response includes percentage, version_distribution, and health.summary.\n", "type": "string", "enum": ["rolling", "new_instances"]}, "target_configuration": {"$ref": "#/components/schemas/cc_ModifyUserDeploymentConfiguration"}, "target_version": {"description": "Target application version after the rollout is complete and applied to all current instances.", "type": "integer"}, "version_distribution": {"description": "Version percentage distribution. Only present for \"new_instances\" strategy.\nFor \"rolling\" strategy, see progress.version_distribution instead.\n", "type": "object", "properties": {"current_version_percentage": {"description": "Percentage of instances on the current (old) version.", "type": "integer"}, "target_version_percentage": {"description": "Percentage of instances on the target (new) version.", "type": "integer"}}, "required": ["target_version_percentage", "current_version_percentage"]}}, "required": ["id", "created_at", "last_updated_at", "description", "kind", "strategy", "current_version", "target_version", "current_configuration", "target_configuration", "status", "health"]}
```
