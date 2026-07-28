---
title: cc_CustomerImageRegistry
page_id: schema-cc-customerimageregistry-fee54c07
path: schemas
description: An image registry added in a customer account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_CustomerImageRegistry

An image registry added in a customer account

```yaml
{"description": "An image registry added in a customer account", "type": "object", "properties": {"created_at": {"$ref": "#/components/schemas/cc_ISO8601Timestamp"}, "domain": {"$ref": "#/components/schemas/cc_Domain"}, "kind": {"description": "The type of registry that is being configured.", "oneOf": [{"$ref": "#/components/schemas/cc_ExternalRegistryKind"}, {"$ref": "#/components/schemas/cc_DefaultImageRegistryKind"}]}, "private_credential": {"$ref": "#/components/schemas/cc_SecretsStoreRef"}, "public_key": {"description": "A base64 representation of the public key that you can set to configure the registry. If null, the registry is public and doesn't have authentication setup with Cloudchamber", "type": "string"}}, "required": ["domain", "created_at"]}
```
