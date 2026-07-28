---
title: secrets-store_value
page_id: schema-secrets-store-value-852e7965
path: schemas
description: The value of the secret. Maximum 64 KiB (65,536 bytes). Note that this is 'write only' - the API never returns this value; it exists only to create or modify secrets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# secrets-store_value

The value of the secret. Maximum 64 KiB (65,536 bytes). Note that this is 'write only' - the API never returns this value; it exists only to create or modify secrets.

```yaml
{"description": "The value of the secret. Maximum 64 KiB (65,536 bytes). Note that this is 'write only' - the API never returns this value; it exists only to create or modify secrets.", "type": "string", "example": "api-token-secret-123", "maxLength": 65536, "writeOnly": true, "x-sensitive": true}
```
