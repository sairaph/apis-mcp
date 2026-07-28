---
title: dlp_NewDataset
page_id: schema-dlp-newdataset-7a1f4780
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_NewDataset

```yaml
{"type": "object", "properties": {"case_sensitive": {"description": "Only applies to custom word lists.\nDetermines if the words should be matched in a case-sensitive manner\nCannot be set to false if `secret` is true or undefined", "type": "boolean"}, "description": {"description": "The description of the dataset.", "type": "string", "nullable": true}, "encoding_version": {"description": "Dataset encoding version\n\nNon-secret custom word lists with no header are always version 1.\nSecret EDM lists with no header are version 1.\nMulticolumn CSV with headers are version 2.\nOmitting this field provides the default value 0, which is interpreted\nthe same as 1.", "type": "integer", "format": "int32", "minimum": 0}, "name": {"type": "string"}, "secret": {"description": "Generate a secret dataset.\n\nIf true, the response will include a secret to use with the EDM encoder.\nIf false, the response has no secret and the dataset is uploaded in plaintext.", "type": "boolean"}}, "required": ["name"]}
```
