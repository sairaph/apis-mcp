---
title: builds_CreateWorkerBuildSettingsInput
page_id: schema-builds-createworkerbuildsettingsinput-837fd07a
path: schemas
description: Build and deploy settings when creating a Worker build configuration
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_CreateWorkerBuildSettingsInput

Build and deploy settings when creating a Worker build configuration

```yaml
{"description": "Build and deploy settings when creating a Worker build configuration", "type": "object", "properties": {"build_caching_enabled": {"default": true, "allOf": [{"$ref": "#/components/schemas/builds_build_caching_enabled"}]}, "build_command": {"$ref": "#/components/schemas/builds_build_command"}, "build_token_uuid": {"$ref": "#/components/schemas/builds_build_token_uuid"}, "deploy_command": {"$ref": "#/components/schemas/builds_deploy_command"}, "environment_variables": {"$ref": "#/components/schemas/builds_EnvironmentVariablesRequest"}, "path_excludes": {"$ref": "#/components/schemas/builds_path_excludes"}, "path_includes": {"$ref": "#/components/schemas/builds_path_includes"}, "root_directory": {"default": "/", "allOf": [{"$ref": "#/components/schemas/builds_root_directory"}]}}, "required": ["build_command", "deploy_command", "build_token_uuid"]}
```
