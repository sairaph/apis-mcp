---
title: builds_UpdateWorkerRequest
page_id: schema-builds-updateworkerrequest-022c30f8
path: schemas
description: Request body for updating a Worker build configuration. At least one field must be provided.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_UpdateWorkerRequest

Request body for updating a Worker build configuration. At least one field must be provided.

```yaml
{"description": "Request body for updating a Worker build configuration. At least one field must be provided.", "type": "object", "properties": {"git_repository": {"description": "Git repository settings to update", "type": "object", "minProperties": 1, "properties": {"branch": {"description": "New git branch to watch for builds", "type": "string", "example": "main", "maxLength": 256, "minLength": 1}}}, "production_settings": {"$ref": "#/components/schemas/builds_UpdateWorkerBuildSettingsInput"}}, "minProperties": 1}
```
