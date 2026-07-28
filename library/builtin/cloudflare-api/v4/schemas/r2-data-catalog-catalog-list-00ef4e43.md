---
title: r2-data-catalog_catalog-list
page_id: schema-r2-data-catalog-catalog-list-00ef4e43
path: schemas
description: Contains the list of catalogs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_catalog-list

Contains the list of catalogs.

```yaml
{"description": "Contains the list of catalogs.", "type": "object", "properties": {"warehouses": {"description": "Lists catalogs in the account.", "type": "array", "items": {"$ref": "#/components/schemas/r2-data-catalog_catalog"}}}, "required": ["warehouses"]}
```
