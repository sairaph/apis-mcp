---
title: builds_UpsertRepoConnectionRequest
page_id: schema-builds-upsertrepoconnectionrequest-e96e46a7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_UpsertRepoConnectionRequest

```yaml
{"type": "object", "properties": {"provider_account_id": {"$ref": "#/components/schemas/builds_provider_account_id"}, "provider_account_name": {"$ref": "#/components/schemas/builds_provider_account_name"}, "provider_type": {"$ref": "#/components/schemas/builds_SCMProviderType"}, "repo_id": {"$ref": "#/components/schemas/builds_repo_id"}, "repo_name": {"$ref": "#/components/schemas/builds_repo_name"}}, "required": ["repo_id", "repo_name", "provider_type", "provider_account_id", "provider_account_name"]}
```
