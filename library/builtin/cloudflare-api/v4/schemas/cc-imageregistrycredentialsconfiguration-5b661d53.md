---
title: cc_ImageRegistryCredentialsConfiguration
page_id: schema-cc-imageregistrycredentialsconfiguration-5b661d53
path: schemas
description: Specifies the configuration for the image registry credential to create.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ImageRegistryCredentialsConfiguration

Specifies the configuration for the image registry credential to create.

```yaml
{"description": "Specifies the configuration for the image registry credential to create.", "type": "object", "properties": {"expiration_minutes": {"description": "The minimum number of minutes the token will be valid for. Must be positive. We make a best effort to respect this value, but some registry providers do not let us configure the token lifetime, so the token may be valid for longer.", "type": "integer", "minimum": 1}, "permissions": {"type": "array", "items": {"$ref": "#/components/schemas/cc_ImageRegistryPermissions"}}}, "required": ["permissions", "expiration_minutes"]}
```
