---
title: builds_CreateDeployHookRequest
page_id: schema-builds-createdeployhookrequest-dadb0e51
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_CreateDeployHookRequest

```yaml
{"type": "object", "properties": {"branch": {"$ref": "#/components/schemas/builds_branch"}, "deploy_hook_name": {"$ref": "#/components/schemas/builds_deploy_hook_name"}}, "required": ["deploy_hook_name", "branch"]}
```
