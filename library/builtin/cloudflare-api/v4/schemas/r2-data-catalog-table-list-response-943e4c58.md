---
title: r2-data-catalog_table-list-response
page_id: schema-r2-data-catalog-table-list-response-943e4c58
path: schemas
description: Contains the list of tables with optional pagination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_table-list-response

Contains the list of tables with optional pagination.

```yaml
{"description": "Contains the list of tables with optional pagination.", "type": "object", "properties": {"details": {"description": "Contains detailed metadata for each table when return_details is true.\nEach object includes identifier, UUID, timestamps, and locations.\n", "type": "array", "items": {"$ref": "#/components/schemas/r2-data-catalog_table-details"}, "nullable": true}, "identifiers": {"description": "Lists tables in the namespace.", "type": "array", "items": {"$ref": "#/components/schemas/r2-data-catalog_table-identifier"}}, "next_page_token": {"description": "Use this opaque token to fetch the next page of results.\nA null or absent value indicates the last page.\n", "type": "string", "example": "MSYxNzU5NzU2MTI4NTU2Njk2JjAxOTliOWEzLTkxMmUtN2ZhMS05YzllLTg5MTAxMGQzYTg0MQ", "nullable": true}, "table_uuids": {"description": "Contains UUIDs for each table when return_uuids is true.\nThe order corresponds to the identifiers array.\n", "type": "array", "items": {"format": "uuid", "type": "string"}, "example": ["0199b9a1-28a0-71e0-a73e-b0fc32c8468e", "0199b9a1-3c74-7731-bf53-d8c67ead079d"], "nullable": true}}, "required": ["identifiers"]}
```
