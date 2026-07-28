---
title: builds_DetailedDeployHookResponse
page_id: schema-builds-detaileddeployhookresponse-a93328aa
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_DetailedDeployHookResponse

```yaml
{"type": "object", "properties": {"branch": {"$ref": "#/components/schemas/builds_branch"}, "created_on": {"$ref": "#/components/schemas/builds_created_on"}, "deploy_hook_name": {"$ref": "#/components/schemas/builds_deploy_hook_name"}, "deploy_hook_uuid": {"$ref": "#/components/schemas/builds_deploy_hook_uuid"}, "external_script_id": {"$ref": "#/components/schemas/builds_external_script_id"}, "latest_build": {"type": "object", "nullable": true, "properties": {"created_on": {"$ref": "#/components/schemas/builds_created_on"}}}, "modified_on": {"$ref": "#/components/schemas/builds_modified_on"}}}
```
