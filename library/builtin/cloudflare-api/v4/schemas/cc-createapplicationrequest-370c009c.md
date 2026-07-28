---
title: cc_CreateApplicationRequest
page_id: schema-cc-createapplicationrequest-370c009c
path: schemas
description: Create a new application object for dynamic scheduling
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_CreateApplicationRequest

Create a new application object for dynamic scheduling

```yaml
{"description": "Create a new application object for dynamic scheduling", "type": "object", "properties": {"configuration": {"$ref": "#/components/schemas/cc_UserDeploymentConfiguration"}, "constraints": {"$ref": "#/components/schemas/cc_ApplicationConstraints"}, "durable_objects": {"$ref": "#/components/schemas/cc_DurableObjectsConfiguration"}, "instances": {"description": "Number of deployments to create", "type": "integer"}, "max_instances": {"description": "Maximum number of instances that the application will allow. This is relevant for applications that auto-scale.", "type": "integer"}, "name": {"description": "The name for this application", "type": "string"}, "observability": {"$ref": "#/components/schemas/cc_ApplicationObservability"}, "rollout_active_grace_period": {"$ref": "#/components/schemas/cc_ApplicationRolloutActiveGracePeriod"}, "scheduling_policy": {"$ref": "#/components/schemas/cc_SchedulingPolicy"}}, "required": ["name", "scheduling_policy", "instances", "configuration"]}
```
