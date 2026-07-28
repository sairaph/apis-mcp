---
title: builds_TriggerResponse
page_id: schema-builds-triggerresponse-b3b383fb
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_TriggerResponse

```yaml
{"type": "object", "properties": {"branch_excludes": {"$ref": "#/components/schemas/builds_branch_excludes"}, "branch_includes": {"$ref": "#/components/schemas/builds_branch_includes"}, "build_caching_enabled": {"$ref": "#/components/schemas/builds_build_caching_enabled"}, "build_command": {"$ref": "#/components/schemas/builds_build_command"}, "build_token_name": {"$ref": "#/components/schemas/builds_build_token_name"}, "build_token_uuid": {"$ref": "#/components/schemas/builds_build_token_uuid"}, "created_on": {"$ref": "#/components/schemas/builds_created_on"}, "deleted_on": {"$ref": "#/components/schemas/builds_deleted_on"}, "deploy_command": {"$ref": "#/components/schemas/builds_deploy_command"}, "external_script_id": {"$ref": "#/components/schemas/builds_external_script_id"}, "modified_on": {"$ref": "#/components/schemas/builds_modified_on"}, "path_excludes": {"$ref": "#/components/schemas/builds_path_excludes"}, "path_includes": {"$ref": "#/components/schemas/builds_path_includes"}, "repo_connection": {"$ref": "#/components/schemas/builds_UpsertRepoConnectionResponse"}, "root_directory": {"$ref": "#/components/schemas/builds_root_directory"}, "trigger_name": {"$ref": "#/components/schemas/builds_trigger_name"}, "trigger_uuid": {"$ref": "#/components/schemas/builds_trigger_uuid"}}}
```
