---
title: secrets-store_createSecretObject
page_id: schema-secrets-store-createsecretobject-08d13a6d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# secrets-store_createSecretObject

```yaml
{"properties": {"comment": {"$ref": "#/components/schemas/secrets-store_comment"}, "name": {"$ref": "#/components/schemas/secrets-store_secret_name"}, "scopes": {"$ref": "#/components/schemas/secrets-store_scopes"}, "value": {"$ref": "#/components/schemas/secrets-store_value"}}, "required": ["name", "value", "scopes"]}
```
