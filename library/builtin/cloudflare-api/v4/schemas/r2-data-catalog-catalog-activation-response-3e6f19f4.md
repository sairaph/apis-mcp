---
title: r2-data-catalog_catalog-activation-response
page_id: schema-r2-data-catalog-catalog-activation-response-3e6f19f4
path: schemas
description: Contains response from activating an R2 bucket as a catalog.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_catalog-activation-response

Contains response from activating an R2 bucket as a catalog.

```yaml
{"description": "Contains response from activating an R2 bucket as a catalog.", "type": "object", "properties": {"id": {"description": "Use this to uniquely identify the activated catalog.", "type": "string", "format": "uuid", "example": "550e8400-e29b-41d4-a716-446655440000"}, "name": {"description": "Specifies the name of the activated catalog.", "type": "string", "example": "account123_my-bucket"}}, "required": ["id", "name"]}
```
