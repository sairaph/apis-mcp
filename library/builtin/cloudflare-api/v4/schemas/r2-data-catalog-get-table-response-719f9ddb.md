---
title: r2-data-catalog_get-table-response
page_id: schema-r2-data-catalog-get-table-response-719f9ddb
path: schemas
description: |-
    Single-table detail response wrapping standard Iceberg `TableMetadata`
    with snapshot-pruning counters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_get-table-response

Single-table detail response wrapping standard Iceberg `TableMetadata`
with snapshot-pruning counters.

```yaml
{"description": "Single-table detail response wrapping standard Iceberg `TableMetadata`\nwith snapshot-pruning counters.\n", "type": "object", "properties": {"identifier": {"$ref": "#/components/schemas/r2-data-catalog_table-identifier"}, "metadata": {"$ref": "#/components/schemas/r2-data-catalog_table-metadata"}, "metadata_location": {"description": "Specifies the S3-compatible URI to the current Iceberg metadata file\nin R2. Omitted for staged tables that have not yet been committed.\n", "type": "string", "example": "s3://my-bucket/__r2_data_catalog/wh-id/table-id/metadata/00001-abc.metadata.json"}, "returned_snapshots": {"description": "Describes the number of snapshots that appear in `metadata.snapshots`. Caps the\nlist at 10 (the most recent by `timestamp-ms`), but if the current\nsnapshot falls outside the recent window, the result includes it in the\nset, leading to a possible total of 11.\n", "type": "integer", "example": 10, "minimum": 0}, "table_uuid": {"description": "Contains the Iceberg table UUID, stable across renames.", "type": "string", "format": "uuid", "example": "550e8400-e29b-41d4-a716-446655440000"}, "total_snapshots": {"description": "Indicates the total number of snapshots stored for the table, before pruning.", "type": "integer", "example": 122113, "minimum": 0}}, "required": ["identifier", "table_uuid", "total_snapshots", "returned_snapshots", "metadata"]}
```
