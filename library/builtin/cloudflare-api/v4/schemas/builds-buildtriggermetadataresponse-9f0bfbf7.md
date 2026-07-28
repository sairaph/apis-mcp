---
title: builds_BuildTriggerMetadataResponse
page_id: schema-builds-buildtriggermetadataresponse-9f0bfbf7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_BuildTriggerMetadataResponse

```yaml
{"type": "object", "properties": {"author": {"type": "string", "example": "developer@cloudflare.com"}, "branch": {"$ref": "#/components/schemas/builds_branch"}, "build_command": {"$ref": "#/components/schemas/builds_build_command"}, "build_token_name": {"$ref": "#/components/schemas/builds_build_token_name"}, "build_token_uuid": {"$ref": "#/components/schemas/builds_build_token_uuid"}, "build_trigger_source": {"$ref": "#/components/schemas/builds_BuildTriggerSource"}, "commit_hash": {"$ref": "#/components/schemas/builds_commit_hash"}, "commit_message": {"type": "string", "example": "Add new feature"}, "deploy_command": {"$ref": "#/components/schemas/builds_deploy_command"}, "environment_variables": {"type": "object", "additionalProperties": {"type": "string"}}, "provider_account_name": {"$ref": "#/components/schemas/builds_provider_account_name"}, "provider_type": {"$ref": "#/components/schemas/builds_SCMProviderType"}, "repo_name": {"$ref": "#/components/schemas/builds_repo_name"}, "root_directory": {"$ref": "#/components/schemas/builds_root_directory"}}}
```
