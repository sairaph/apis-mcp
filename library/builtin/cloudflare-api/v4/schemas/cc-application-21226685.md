---
title: cc_Application
page_id: schema-cc-application-21226685
path: schemas
description: Describes multiple deployments with parameters that describe how they should be placed
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_Application

Describes multiple deployments with parameters that describe how they should be placed

```yaml
{"description": "Describes multiple deployments with parameters that describe how they should be placed", "type": "object", "properties": {"account_id": {"$ref": "#/components/schemas/cc_AccountID"}, "active_rollout_id": {"$ref": "#/components/schemas/cc_RolloutID"}, "configuration": {"$ref": "#/components/schemas/cc_UserDeploymentConfiguration"}, "constraints": {"$ref": "#/components/schemas/cc_ApplicationConstraints"}, "created_at": {"$ref": "#/components/schemas/cc_ISO8601Timestamp"}, "durable_objects": {"$ref": "#/components/schemas/cc_ApplicationDurableObjectsConfiguration"}, "health": {"$ref": "#/components/schemas/cc_ApplicationHealth"}, "id": {"$ref": "#/components/schemas/cc_ApplicationID"}, "instances": {"description": "Number of deployments to create", "type": "integer"}, "max_instances": {"description": "Maximum number of instances that the application will allow. This is relevant for applications that auto-scale.", "type": "integer"}, "name": {"$ref": "#/components/schemas/cc_ApplicationName"}, "observability": {"$ref": "#/components/schemas/cc_ApplicationObservability"}, "rollout_active_grace_period": {"$ref": "#/components/schemas/cc_ApplicationRolloutActiveGracePeriod"}, "scheduling_policy": {"$ref": "#/components/schemas/cc_SchedulingPolicy"}, "updated_at": {"$ref": "#/components/schemas/cc_ISO8601Timestamp"}, "version": {"type": "integer"}}, "required": ["id", "created_at", "updated_at", "account_id", "name", "version", "scheduling_policy", "instances", "configuration"]}
```
