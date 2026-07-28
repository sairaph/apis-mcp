---
title: builds_CreateWorkerGitRepositoryFunfettiInput
page_id: schema-builds-createworkergitrepositoryfunfettiinput-21038a11
path: schemas
description: GitHub or GitLab repository input for creating a Worker build configuration
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_CreateWorkerGitRepositoryFunfettiInput

GitHub or GitLab repository input for creating a Worker build configuration

```yaml
{"description": "GitHub or GitLab repository input for creating a Worker build configuration", "type": "object", "properties": {"branch": {"description": "Git branch to watch for builds", "type": "string", "example": "main", "maxLength": 256, "minLength": 1}, "provider_account_id": {"$ref": "#/components/schemas/builds_provider_account_id"}, "provider_account_name": {"$ref": "#/components/schemas/builds_provider_account_name"}, "provider_type": {"type": "string", "example": "github", "enum": ["github", "gitlab"]}, "repo_id": {"$ref": "#/components/schemas/builds_repo_id"}, "repo_name": {"$ref": "#/components/schemas/builds_repo_name"}}, "required": ["repo_id", "repo_name", "provider_type", "provider_account_id", "provider_account_name", "branch"]}
```
