---
title: secrets-store_secretObject
page_id: schema-secrets-store-secretobject-f940e39b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# secrets-store_secretObject

```yaml
{"properties": {"comment": {"$ref": "#/components/schemas/secrets-store_comment"}, "created": {"$ref": "#/components/schemas/secrets-store_created"}, "id": {"$ref": "#/components/schemas/secrets-store_identifier"}, "modified": {"$ref": "#/components/schemas/secrets-store_modified"}, "name": {"$ref": "#/components/schemas/secrets-store_secret_name"}, "scopes": {"$ref": "#/components/schemas/secrets-store_scopes"}, "status": {"$ref": "#/components/schemas/secrets-store_SecretStatus"}, "store_id": {"$ref": "#/components/schemas/secrets-store_store_identifier"}}, "required": ["id", "name", "store_id", "created", "modified", "status"]}
```
