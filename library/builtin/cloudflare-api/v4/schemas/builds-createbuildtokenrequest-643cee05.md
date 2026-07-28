---
title: builds_CreateBuildTokenRequest
page_id: schema-builds-createbuildtokenrequest-643cee05
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_CreateBuildTokenRequest

```yaml
{"type": "object", "properties": {"build_token_name": {"$ref": "#/components/schemas/builds_build_token_name"}, "build_token_secret": {"type": "string", "example": "super-secret-token"}, "cloudflare_token_id": {"$ref": "#/components/schemas/builds_cloudflare_token_id"}}, "required": ["build_token_name", "build_token_secret", "cloudflare_token_id"]}
```
