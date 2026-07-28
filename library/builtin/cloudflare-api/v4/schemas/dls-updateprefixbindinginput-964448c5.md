---
title: dls_UpdatePrefixBindingInput
page_id: schema-dls-updateprefixbindinginput-964448c5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dls_UpdatePrefixBindingInput

```yaml
{"type": "object", "properties": {"region_key": {"description": "New region key to assign (e.g., \"us\", \"eu\", \"cfcanary\").", "type": "string", "example": "eu", "maxLength": 128, "minLength": 1, "pattern": "^[a-z0-9_-]+$"}}, "required": ["region_key"]}
```
