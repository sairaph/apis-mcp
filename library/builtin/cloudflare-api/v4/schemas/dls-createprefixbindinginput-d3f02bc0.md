---
title: dls_CreatePrefixBindingInput
page_id: schema-dls-createprefixbindinginput-d3f02bc0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dls_CreatePrefixBindingInput

```yaml
{"type": "object", "properties": {"cidr": {"description": "IP prefix in CIDR notation to bind.", "type": "string", "example": "10.0.1.0/24"}, "prefix_id": {"description": "The ID of the parent IP prefix that contains the CIDR.", "type": "string", "example": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "maxLength": 64}, "region_key": {"description": "Region key from managed regions (e.g., \"us\", \"eu\").", "type": "string", "example": "eu", "maxLength": 128, "minLength": 1, "pattern": "^[a-z0-9_-]+$"}}, "required": ["prefix_id", "cidr", "region_key"]}
```
