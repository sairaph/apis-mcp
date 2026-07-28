---
title: dls_PrefixBinding
page_id: schema-dls-prefixbinding-8fe001b6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dls_PrefixBinding

```yaml
{"type": "object", "properties": {"cidr": {"description": "The CIDR that is bound.", "type": "string"}, "id": {"description": "The ID of the binding.", "type": "string"}, "prefix_id": {"description": "The ID of the parent prefix.", "type": "string"}, "region_key": {"description": "The region key used for the binding.", "type": "string", "maxLength": 128, "minLength": 1, "pattern": "^[a-z0-9_-]+$"}}, "required": ["id", "prefix_id", "cidr", "region_key"]}
```
