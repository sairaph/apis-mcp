---
title: builds_UpdateWorkerBuildSettingsInput
page_id: schema-builds-updateworkerbuildsettingsinput-3eb33945
path: schemas
description: Partial build settings for updating a Worker build configuration
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_UpdateWorkerBuildSettingsInput

Partial build settings for updating a Worker build configuration

```yaml
{"description": "Partial build settings for updating a Worker build configuration", "type": "object", "properties": {"build_caching_enabled": {"$ref": "#/components/schemas/builds_build_caching_enabled"}, "build_command": {"$ref": "#/components/schemas/builds_build_command"}, "build_token_uuid": {"$ref": "#/components/schemas/builds_build_token_uuid"}, "deploy_command": {"$ref": "#/components/schemas/builds_deploy_command"}, "environment_variables": {"description": "Environment variable updates. Set the variable entry to null to delete it.", "type": "object", "additionalProperties": {"nullable": true, "properties": {"is_secret": {"$ref": "#/components/schemas/builds_is_secret"}, "value": {"type": "string"}}, "required": ["is_secret", "value"], "type": "object"}}, "path_excludes": {"$ref": "#/components/schemas/builds_path_excludes"}, "path_includes": {"$ref": "#/components/schemas/builds_path_includes"}, "root_directory": {"$ref": "#/components/schemas/builds_root_directory"}}, "minProperties": 1}
```
