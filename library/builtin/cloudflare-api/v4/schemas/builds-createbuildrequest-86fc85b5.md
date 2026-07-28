---
title: builds_CreateBuildRequest
page_id: schema-builds-createbuildrequest-86fc85b5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_CreateBuildRequest

```yaml
{"type": "object", "properties": {"branch": {"description": "Git branch name (required if commit_hash not provided)", "allOf": [{"$ref": "#/components/schemas/builds_branch"}]}, "commit_hash": {"description": "Git commit hash (required if branch not provided)", "allOf": [{"$ref": "#/components/schemas/builds_commit_hash"}]}, "seed_repo": {"$ref": "#/components/schemas/builds_BuildSeedRepoInput"}}, "anyOf": [{"required": ["commit_hash"]}, {"required": ["branch"]}]}
```
