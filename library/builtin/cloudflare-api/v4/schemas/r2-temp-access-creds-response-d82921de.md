---
title: r2_temp_access_creds_response
page_id: schema-r2-temp-access-creds-response-d82921de
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_temp_access_creds_response

```yaml
{"type": "object", "properties": {"accessKeyId": {"description": "ID for new access key.", "type": "string"}, "secretAccessKey": {"description": "Secret access key.", "type": "string", "x-sensitive": true}, "sessionToken": {"description": "Security token.", "type": "string", "x-sensitive": true}}, "example": {"accessKeyId": "example-access-key-id", "secretAccessKey": "example-secret-key", "sessionToken": "example-session-token"}}
```
