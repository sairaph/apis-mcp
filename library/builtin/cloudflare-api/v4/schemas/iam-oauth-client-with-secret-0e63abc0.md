---
title: iam_oauth_client_with_secret
page_id: schema-iam-oauth-client-with-secret-0e63abc0
path: schemas
description: An OAuth client response that includes the client secret. Only returned on client creation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_oauth_client_with_secret

An OAuth client response that includes the client secret. Only returned on client creation.

```yaml
{"description": "An OAuth client response that includes the client secret. Only returned on client creation.", "allOf": [{"$ref": "#/components/schemas/iam_oauth_client"}, {"properties": {"client_secret": {"description": "The client secret. This is the only time the secret is returned in a response.", "type": "string", "example": "cf-oauth-secret-example", "readOnly": true, "x-sensitive": true}}, "type": "object"}], "title": "OAuth Client with Secret"}
```
