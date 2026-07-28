---
title: builds_WorkerGitRepository
page_id: schema-builds-workergitrepository-7f5eae1d
path: schemas
description: Git repository details linked to a Worker script build configuration
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_WorkerGitRepository

Git repository details linked to a Worker script build configuration

```yaml
{"description": "Git repository details linked to a Worker script build configuration", "type": "object", "properties": {"branch": {"description": "Git branch to watch for builds", "type": "string", "example": "main", "maxLength": 256, "minLength": 1}, "grant_id": {"description": "Internal grant ID for GitLab internal integrations", "type": "string", "nullable": true}, "provider_account_id": {"$ref": "#/components/schemas/builds_provider_account_id"}, "provider_account_name": {"$ref": "#/components/schemas/builds_provider_account_name"}, "provider_type": {"$ref": "#/components/schemas/builds_SCMProviderType"}, "repo_id": {"$ref": "#/components/schemas/builds_repo_id"}, "repo_name": {"$ref": "#/components/schemas/builds_repo_name"}}}
```
