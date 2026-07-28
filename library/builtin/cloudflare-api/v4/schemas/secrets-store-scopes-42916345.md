---
title: secrets-store_scopes
page_id: schema-secrets-store-scopes-42916345
path: schemas
description: The list of services that can use this secret.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# secrets-store_scopes

The list of services that can use this secret.

```yaml
{"description": "The list of services that can use this secret.", "type": "array", "items": {"enum": ["workers", "ai_gateway", "dex", "access", "containers", "websearch"], "type": "string"}, "example": ["workers", "ai_gateway", "dex", "access", "websearch"]}
```
