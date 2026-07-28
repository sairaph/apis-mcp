---
title: iam_oauth_client_rotate_secret_response
page_id: schema-iam-oauth-client-rotate-secret-response-2313497f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_oauth_client_rotate_secret_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/iam_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"client_secret": {"description": "The new client secret.", "type": "string", "example": "cf-oauth-secret-new-example", "readOnly": true, "x-sensitive": true}}}}, "type": "object"}], "title": "Rotate OAuth Client Secret Response"}
```
