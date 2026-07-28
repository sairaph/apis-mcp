---
title: posture-api_AssetField
page_id: schema-posture-api-assetfield-4848616f
path: schemas
description: Additional field information for an asset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_AssetField

Additional field information for an asset.

```yaml
{"description": "Additional field information for an asset.", "type": "object", "properties": {"link": {"description": "Optional link associated with the field.", "type": "string", "format": "uri", "example": "https://example.com", "nullable": true}, "name": {"description": "The name of the field.", "type": "string", "example": "File Name"}, "value": {"description": "The value of the field.", "type": "string", "example": "sensitive-document.xlsx"}}, "example": {"link": "https://example.com", "name": "Credential name", "value": "Test asset 2"}, "required": ["name", "value"]}
```
