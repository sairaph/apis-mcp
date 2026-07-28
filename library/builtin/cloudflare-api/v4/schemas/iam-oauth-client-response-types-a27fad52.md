---
title: iam_oauth_client_response_types
page_id: schema-iam-oauth-client-response-types-a27fad52
path: schemas
description: Array of OAuth response types the client is allowed to use.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_oauth_client_response_types

Array of OAuth response types the client is allowed to use.

```yaml
{"description": "Array of OAuth response types the client is allowed to use.", "type": "array", "items": {"enum": ["token", "id_token", "code"], "type": "string"}, "example": ["code"], "x-auditable": true}
```
