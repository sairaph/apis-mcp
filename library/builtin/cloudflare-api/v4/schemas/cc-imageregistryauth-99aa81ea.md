---
title: cc_ImageRegistryAuth
page_id: schema-cc-imageregistryauth-99aa81ea
path: schemas
description: Credentials needed to authenticate with an external image registry.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ImageRegistryAuth

Credentials needed to authenticate with an external image registry.

```yaml
{"description": "Credentials needed to authenticate with an external image registry.", "type": "object", "properties": {"private_credential": {"oneOf": [{"type": "string"}, {"$ref": "#/components/schemas/cc_SecretsStoreRef"}]}, "public_credential": {"description": "The format of this value is determined by the registry being configured.", "type": "string", "example": "AWS_ACCESS_KEY_ID"}}, "required": ["public_credential", "private_credential"]}
```
