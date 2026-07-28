---
title: cc_ModifyApplicationRequestBody
page_id: schema-cc-modifyapplicationrequestbody-20c77411
path: schemas
description: Request body for modifying an application
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ModifyApplicationRequestBody

Request body for modifying an application

```yaml
{"description": "Request body for modifying an application", "type": "object", "properties": {"configuration": {"$ref": "#/components/schemas/cc_ModifyUserDeploymentConfiguration"}, "constraints": {"$ref": "#/components/schemas/cc_ApplicationConstraints"}, "instances": {"description": "Number of deployments to maintain within this application. This can be used to scale the application up/down.", "type": "integer"}, "max_instances": {"description": "Maximum number of instances that the application will allow. This is relevant for applications that auto-scale.\nIt will reduce the number of running instances if there are more than `max_instances`.\n", "type": "integer"}, "name": {"description": "The name for this application", "type": "string"}, "observability": {"$ref": "#/components/schemas/cc_ApplicationObservability"}, "rollout_active_grace_period": {"$ref": "#/components/schemas/cc_ApplicationRolloutActiveGracePeriod"}, "scheduling_policy": {"$ref": "#/components/schemas/cc_SchedulingPolicy"}}}
```
