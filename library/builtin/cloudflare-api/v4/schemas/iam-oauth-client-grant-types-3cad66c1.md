---
title: iam_oauth_client_grant_types
page_id: schema-iam-oauth-client-grant-types-3cad66c1
path: schemas
description: Array of OAuth grant types the client is allowed to use. `authorization_code` is required; `refresh_token` may be included optionally.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_oauth_client_grant_types

Array of OAuth grant types the client is allowed to use. `authorization_code` is required; `refresh_token` may be included optionally.

```yaml
{"description": "Array of OAuth grant types the client is allowed to use. `authorization_code` is required; `refresh_token` may be included optionally.", "type": "array", "items": {"enum": ["authorization_code", "refresh_token"], "type": "string"}, "example": ["authorization_code", "refresh_token"], "x-auditable": true}
```
