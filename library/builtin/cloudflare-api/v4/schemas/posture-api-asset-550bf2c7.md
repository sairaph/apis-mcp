---
title: posture-api_Asset
page_id: schema-posture-api-asset-550bf2c7
path: schemas
description: Asset information including metadata and categorization.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_Asset

Asset information including metadata and categorization.

```yaml
{"description": "Asset information including metadata and categorization.", "type": "object", "properties": {"category": {"$ref": "#/components/schemas/posture-api_AssetCategory"}, "external_id": {"description": "External identifier from the source system.", "type": "string", "example": "external-file-id-123", "maxLength": 512}, "fields": {"description": "The fields associated with the asset.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_AssetField"}, "readOnly": true}, "id": {"description": "Unique identifier for the asset.", "type": "string", "format": "uuid", "example": "8a043daf-def4-403e-9d28-da2e93d9b824"}, "link": {"description": "Direct link to the asset.", "type": "string", "format": "uri", "example": "https://slack-files.com/TYJH37DCK-E0238GG6B8-92fd5y5674", "maxLength": 2048, "nullable": true}, "name": {"description": "Human-readable name of the asset.", "type": "string", "example": "Public.svg"}}, "example": {"external_id": "external-file-id-123", "id": "8a043daf-def4-403e-9d28-da2e93d9b824", "link": "https://slack-files.com/TYJH37DCK-E0238GG6B8-92fd5y5674", "name": "Public.svg"}, "required": ["category", "external_id", "fields", "name"]}
```
