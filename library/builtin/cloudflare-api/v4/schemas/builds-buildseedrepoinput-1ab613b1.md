---
title: builds_BuildSeedRepoInput
page_id: schema-builds-buildseedrepoinput-1ab613b1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_BuildSeedRepoInput

```yaml
{"type": "object", "properties": {"branch": {"$ref": "#/components/schemas/builds_branch"}, "files": {"type": "array", "items": {"$ref": "#/components/schemas/builds_BuildSeedRepoInputFile"}, "maxItems": 2}, "owner": {"type": "string", "example": "cloudflare"}, "path": {"type": "string", "example": "/"}, "provider": {"$ref": "#/components/schemas/builds_SCMProviderType"}, "repository": {"type": "string", "example": "workers-sdk"}}, "required": ["provider", "owner", "repository", "branch", "path"]}
```
