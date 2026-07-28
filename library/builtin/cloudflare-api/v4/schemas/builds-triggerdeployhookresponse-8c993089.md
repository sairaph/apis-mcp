---
title: builds_TriggerDeployHookResponse
page_id: schema-builds-triggerdeployhookresponse-8c993089
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_TriggerDeployHookResponse

```yaml
{"type": "object", "properties": {"already_exists": {"description": "True if a pending build already exists for this branch", "type": "boolean", "example": false}, "build_uuid": {"$ref": "#/components/schemas/builds_build_uuid"}, "created_on": {"$ref": "#/components/schemas/builds_created_on"}, "status": {"$ref": "#/components/schemas/builds_BuildStatus"}}}
```
