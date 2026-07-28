---
title: builds_CreateWorkerGitRepositoryInput
page_id: schema-builds-createworkergitrepositoryinput-002bdfb0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_CreateWorkerGitRepositoryInput

```yaml
{"discriminator": {"mapping": {"github": "#/components/schemas/builds_CreateWorkerGitRepositoryFunfettiInput", "gitlab": "#/components/schemas/builds_CreateWorkerGitRepositoryFunfettiInput", "gitlab_internal": "#/components/schemas/builds_CreateWorkerGitRepositoryGrantInput"}, "propertyName": "provider_type"}, "oneOf": [{"$ref": "#/components/schemas/builds_CreateWorkerGitRepositoryFunfettiInput"}, {"$ref": "#/components/schemas/builds_CreateWorkerGitRepositoryGrantInput"}]}
```
