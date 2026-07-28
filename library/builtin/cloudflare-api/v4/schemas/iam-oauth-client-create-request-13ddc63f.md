---
title: iam_oauth_client_create_request
page_id: schema-iam-oauth-client-create-request-13ddc63f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_oauth_client_create_request

```yaml
{"allOf": [{"$ref": "#/components/schemas/iam_oauth_client_common"}, {"required": ["client_name", "grant_types", "redirect_uris", "response_types", "scopes", "token_endpoint_auth_method"], "type": "object"}], "title": "Create OAuth Client Request"}
```
