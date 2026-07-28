---
title: r2-data-catalog_table-details
page_id: schema-r2-data-catalog-table-details-e7628271
path: schemas
description: Contains table with metadata.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_table-details

Contains table with metadata.

```yaml
{"description": "Contains table with metadata.", "type": "object", "properties": {"created_at": {"description": "Indicates the creation timestamp in ISO 8601 format.", "type": "string", "format": "date-time", "nullable": true}, "identifier": {"$ref": "#/components/schemas/r2-data-catalog_table-identifier"}, "location": {"description": "Specifies the base S3 URI for table storage location.", "type": "string", "nullable": true}, "metadata_location": {"description": "Contains the S3 URI to table metadata file. Null for staged tables.", "type": "string", "nullable": true}, "table_uuid": {"description": "Contains the UUID that persists across renames.", "type": "string", "format": "uuid"}, "updated_at": {"description": "Shows the last update timestamp in ISO 8601 format. Null if never updated.", "type": "string", "format": "date-time", "nullable": true}}, "required": ["identifier", "table_uuid"]}
```
