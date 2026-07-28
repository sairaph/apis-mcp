---
title: posture-api_AssetCategory
page_id: schema-posture-api-assetcategory-d08f3d6c
path: schemas
description: Category information for an asset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_AssetCategory

Category information for an asset.

```yaml
{"description": "Category information for an asset.", "type": "object", "properties": {"id": {"description": "Unique identifier for the asset category.", "type": "string", "format": "uuid", "example": "1a78cbf3-b98f-4289-b1f2-22db64130f4f"}, "service": {"description": "The specific service within the vendor the asset is part of (often none). Example - AWS is the vendor, S3 is the service.", "type": "string", "example": "OneDrive", "nullable": true}, "type": {"description": "The type of asset.", "type": "string", "example": "file"}, "vendor": {"description": "The vendor the asset is part of.", "type": "string", "example": "Slack", "readOnly": true}}, "example": {"id": "1a78cbf3-b98f-4289-b1f2-22db64130f4f", "service": null, "type": "file", "vendor": "Slack"}, "required": ["service", "type", "vendor"]}
```
