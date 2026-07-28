---
title: dlp_EntryWithSharedProfiles
page_id: schema-dlp-entrywithsharedprofiles-e50f5170
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_EntryWithSharedProfiles

```yaml
{"allOf": [{"$ref": "#/components/schemas/dlp_EntryWithUploadStatus"}, {"properties": {"profiles": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_EntryProfile"}, "x-stainless-terraform-configurability": "computed_optional"}}, "required": ["profiles"], "type": "object"}]}
```
