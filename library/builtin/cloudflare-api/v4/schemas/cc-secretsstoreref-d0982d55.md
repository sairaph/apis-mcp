---
title: cc_SecretsStoreRef
page_id: schema-cc-secretsstoreref-d0982d55
path: schemas
description: A reference to a secret stored in Secrets Store
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_SecretsStoreRef

A reference to a secret stored in Secrets Store

```yaml
{"description": "A reference to a secret stored in Secrets Store", "type": "object", "properties": {"secret_name": {"description": "Name of the secret being referenced", "type": "string", "example": "API_KEY", "maxLength": 255, "minLength": 1, "pattern": "^[A-z0-9-_]+$"}, "store_id": {"description": "Store ID where the secret is stored", "type": "string", "example": "14758f1afd44c09b7992073ccf00b43d", "maxLength": 32, "minLength": 32, "pattern": "^[a-f0-9]{32}$"}}, "required": ["store_id", "secret_name"]}
```
