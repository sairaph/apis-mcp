---
title: dlp_EntryWithUploadStatus
page_id: schema-dlp-entrywithuploadstatus-93833823
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_EntryWithUploadStatus

```yaml
{"allOf": [{"$ref": "#/components/schemas/dlp_Entry"}, {"properties": {"upload_status": {"allOf": [{"$ref": "#/components/schemas/dlp_DatasetUploadStatus"}], "x-stainless-terraform-configurability": "computed_optional"}}, "type": "object"}]}
```
