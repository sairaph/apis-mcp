---
title: custom-pages_custom_asset
page_id: schema-custom-pages-custom-asset-8a4443f7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# custom-pages_custom_asset

```yaml
{"type": "object", "properties": {"description": {"$ref": "#/components/schemas/custom-pages_asset_description"}, "last_updated": {"$ref": "#/components/schemas/custom-pages_timestamp"}, "name": {"$ref": "#/components/schemas/custom-pages_asset_name"}, "size_bytes": {"description": "The size of the asset content in bytes.", "type": "integer", "example": 1024, "readOnly": true}, "url": {"$ref": "#/components/schemas/custom-pages_asset_url"}}}
```
