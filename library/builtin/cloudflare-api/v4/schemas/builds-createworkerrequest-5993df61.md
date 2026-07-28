---
title: builds_CreateWorkerRequest
page_id: schema-builds-createworkerrequest-5993df61
path: schemas
description: Request body for creating a Worker build configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_CreateWorkerRequest

Request body for creating a Worker build configuration.

```yaml
{"description": "Request body for creating a Worker build configuration.", "type": "object", "properties": {"git_repository": {"$ref": "#/components/schemas/builds_CreateWorkerGitRepositoryInput"}, "production_settings": {"$ref": "#/components/schemas/builds_CreateWorkerBuildSettingsInput"}, "script_tag": {"$ref": "#/components/schemas/builds_external_script_id"}}, "required": ["script_tag", "git_repository", "production_settings"]}
```
