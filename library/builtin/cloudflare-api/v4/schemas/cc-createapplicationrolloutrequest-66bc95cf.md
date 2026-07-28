---
title: cc_CreateApplicationRolloutRequest
page_id: schema-cc-createapplicationrolloutrequest-66bc95cf
path: schemas
description: Request body to create a new rollout for an application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_CreateApplicationRolloutRequest

Request body to create a new rollout for an application.

```yaml
{"description": "Request body to create a new rollout for an application.", "type": "object", "properties": {"description": {"description": "Description of the rollout process.", "type": "string"}, "kind": {"description": "Kind of the rollout process.\n - \"full_auto\": For rolling rollouts, starts progressing steps upon rollout creation. For new_instances rollouts, advances percentage targets automatically after target-version health is observed.\n - \"full_manual\": Requires manually progressing each step in the rollout using the UpdateRollout's action paramater.\nFor \"new_instances\" strategy, omit \"kind\" to preserve manual percentage behavior, or set \"full_auto\" to start an automatic 10% health-gated ramp.\n", "type": "string", "enum": ["full_auto", "full_manual"]}, "percentage": {"description": "Initial target version percentage (0-100). Version sync will actively replace instances to match.\nRequired when strategy is \"new_instances\" and kind is \"full_manual\". When strategy is \"new_instances\" and kind is \"full_auto\", omitted percentage starts at 10% or the smallest percentage that targets at least one instance. Not used for \"rolling\".\n", "type": "integer", "maximum": 100, "minimum": 0}, "step_percentage": {"description": "Percentage of rollout to increase in each step when \"steps\" is not specificed. Applicable values are 5, 10, 20, 25, 50, 100.\nThese create rollouts with 20, 10, 5, 4, 2, 1 steps respectively.\nOnly valid for \"rolling\" strategy.\n", "type": "integer", "enum": [5, 10, 20, 25, 50, 100]}, "steps": {"description": "Steps defining the rollout process, when \"step_percentage\" is not defined.\nOnly one of \"step_percentage\" or \"steps\" can be defined when creating a rollout.\n\"steps\" allow granular control over each step.\nOnly valid for \"rolling\" strategy.\n", "type": "array", "items": {"$ref": "#/components/schemas/cc_RolloutStepRequest"}}, "strategy": {"description": "Strategy used for the rollout.\n- \"rolling\": Step-based rollout with health gates. Actively replaces instances to reach each step's target percentage.\n- \"new_instances\": Percentage control over version distribution. Version sync actively replaces instances to match the configured percentage. The \"full_auto\" kind advances through fixed percentage targets after target-version health is observed.\n", "type": "string", "enum": ["rolling", "new_instances"]}, "target_configuration": {"$ref": "#/components/schemas/cc_ModifyUserDeploymentConfiguration"}}, "required": ["target_configuration", "strategy", "description"]}
```
