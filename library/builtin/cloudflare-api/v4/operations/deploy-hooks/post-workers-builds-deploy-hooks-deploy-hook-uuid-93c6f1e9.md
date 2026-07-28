---
title: Trigger deploy hook
page_id: operation-post-workers-builds-deploy-hooks-deploy-hook-uuid-93539086
path: operations/deploy-hooks
description: Trigger a build using a deploy hook. This endpoint does not require authentication - the deploy_hook_uuid acts as a secret token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /workers/builds/deploy_hooks/{deploy_hook_uuid}
operation_ids:
    - triggerDeployHook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Trigger deploy hook

`POST /workers/builds/deploy_hooks/{deploy_hook_uuid}`

Operation ID: `triggerDeployHook`

Trigger a build using a deploy hook. This endpoint does not require authentication - the deploy_hook_uuid acts as a secret token.

## Definition

```yaml
{"operationId": "triggerDeployHook", "summary": "Trigger deploy hook", "description": "Trigger a build using a deploy hook. This endpoint does not require authentication - the deploy_hook_uuid acts as a secret token.", "parameters": [{"$ref": "#/components/parameters/builds_DeployHookUuid"}], "responses": {"200": {"description": "Build triggered successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_TriggerDeployHookResponse"}}, "type": "object"}]}}}}, "404": {"$ref": "#/components/responses/builds_NotFound"}, "429": {"$ref": "#/components/responses/builds_TooManyRequests"}}, "security": [], "tags": ["Deploy Hooks"]}
```
