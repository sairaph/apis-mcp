---
title: cc_AccountRegistryToken
page_id: schema-cc-accountregistrytoken-b1412653
path: schemas
description: Credentials that can be used to interact with the requested image registry.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_AccountRegistryToken

Credentials that can be used to interact with the requested image registry.

```yaml
{"description": "Credentials that can be used to interact with the requested image registry.", "type": "object", "properties": {"account_id": {"$ref": "#/components/schemas/cc_AccountID"}, "password": {"description": "The password to use when authenticating to the image registry.", "type": "string"}, "registry_host": {"description": "The domain of the image registry the credentials are for.", "type": "string"}, "username": {"description": "The username to use when authenticating to the image registry.", "type": "string"}}, "required": ["account_id", "registry_host", "username"]}
```
