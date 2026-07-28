---
title: builds_WorkerBuildSettings
page_id: schema-builds-workerbuildsettings-9e2fa888
path: schemas
description: Build and deploy settings for a Worker script environment
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_WorkerBuildSettings

Build and deploy settings for a Worker script environment

```yaml
{"description": "Build and deploy settings for a Worker script environment", "type": "object", "properties": {"build_caching_enabled": {"$ref": "#/components/schemas/builds_build_caching_enabled"}, "build_command": {"$ref": "#/components/schemas/builds_build_command"}, "build_token_uuid": {"$ref": "#/components/schemas/builds_build_token_uuid"}, "deploy_command": {"$ref": "#/components/schemas/builds_deploy_command"}, "environment_variables": {"$ref": "#/components/schemas/builds_EnvironmentVariablesResponse"}, "path_excludes": {"$ref": "#/components/schemas/builds_path_excludes"}, "path_includes": {"$ref": "#/components/schemas/builds_path_includes"}, "root_directory": {"$ref": "#/components/schemas/builds_root_directory"}}}
```
