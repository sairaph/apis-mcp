---
title: builds_CreateTriggerRequest
page_id: schema-builds-createtriggerrequest-9952c039
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_CreateTriggerRequest

```yaml
{"type": "object", "properties": {"branch_excludes": {"$ref": "#/components/schemas/builds_branch_excludes"}, "branch_includes": {"$ref": "#/components/schemas/builds_branch_includes"}, "build_caching_enabled": {"$ref": "#/components/schemas/builds_build_caching_enabled"}, "build_command": {"$ref": "#/components/schemas/builds_build_command"}, "build_token_uuid": {"$ref": "#/components/schemas/builds_build_token_uuid"}, "deploy_command": {"$ref": "#/components/schemas/builds_deploy_command"}, "external_script_id": {"$ref": "#/components/schemas/builds_external_script_id"}, "path_excludes": {"$ref": "#/components/schemas/builds_path_excludes"}, "path_includes": {"$ref": "#/components/schemas/builds_path_includes"}, "repo_connection_uuid": {"$ref": "#/components/schemas/builds_repo_connection_uuid"}, "root_directory": {"$ref": "#/components/schemas/builds_root_directory"}, "trigger_name": {"$ref": "#/components/schemas/builds_trigger_name"}}, "required": ["external_script_id", "build_token_uuid", "trigger_name", "build_command", "deploy_command", "root_directory", "branch_includes", "branch_excludes", "path_includes", "path_excludes", "repo_connection_uuid"]}
```
