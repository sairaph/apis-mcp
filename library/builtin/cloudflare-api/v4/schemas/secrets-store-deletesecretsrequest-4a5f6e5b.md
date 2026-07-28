---
title: secrets-store_deleteSecretsRequest
page_id: schema-secrets-store-deletesecretsrequest-4a5f6e5b
path: schemas
description: Request body for bulk deleting secrets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# secrets-store_deleteSecretsRequest

Request body for bulk deleting secrets.

```yaml
{"description": "Request body for bulk deleting secrets.", "type": "object", "properties": {"ids": {"description": "List of secret identifier tags to delete.", "type": "array", "items": {"$ref": "#/components/schemas/secrets-store_identifier"}, "example": ["3fd85f74b32742f1bff64a85009dda07"]}}, "required": ["ids"]}
```
