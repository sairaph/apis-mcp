---
title: builds_CreateWorkerGitRepositoryGrantInput
page_id: schema-builds-createworkergitrepositorygrantinput-6e2fa306
path: schemas
description: Internal GitLab repository input for creating a Worker build configuration
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_CreateWorkerGitRepositoryGrantInput

Internal GitLab repository input for creating a Worker build configuration

```yaml
{"description": "Internal GitLab repository input for creating a Worker build configuration", "type": "object", "properties": {"branch": {"description": "Git branch to watch for builds", "type": "string", "example": "main", "maxLength": 256, "minLength": 1}, "grant_id": {"description": "Grant ID required when provider_type is gitlab_internal", "type": "string", "example": "grant-123", "minLength": 1}, "provider_account_id": {"$ref": "#/components/schemas/builds_provider_account_id"}, "provider_account_name": {"$ref": "#/components/schemas/builds_provider_account_name"}, "provider_type": {"type": "string", "example": "gitlab_internal", "enum": ["gitlab_internal"]}, "repo_id": {"$ref": "#/components/schemas/builds_repo_id"}, "repo_name": {"$ref": "#/components/schemas/builds_repo_name"}}, "required": ["repo_id", "repo_name", "provider_type", "provider_account_id", "provider_account_name", "branch", "grant_id"]}
```
