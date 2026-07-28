---
title: r2-data-catalog_table-metadata
page_id: schema-r2-data-catalog-table-metadata-ea634178
path: schemas
description: |-
    Contains standard Apache Iceberg table metadata (format version 1 or 2).
    See https://iceberg.apache.org/spec/#table-metadata-fields for the full
    field reference. Uses the `s3://` URI scheme for S3 compatibility,
    regardless of the underlying R2 backend.

    Limits the `snapshots`, `snapshot-log`, and `metadata-log` arrays to the
    most recent 10 entries (by `timestamp-ms`). Always keeps the current
    snapshot (via `current-snapshot-id`) even when it falls outside that
    window.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_table-metadata

Contains standard Apache Iceberg table metadata (format version 1 or 2).
See https://iceberg.apache.org/spec/#table-metadata-fields for the full
field reference. Uses the `s3://` URI scheme for S3 compatibility,
regardless of the underlying R2 backend.

Limits the `snapshots`, `snapshot-log`, and `metadata-log` arrays to the
most recent 10 entries (by `timestamp-ms`). Always keeps the current
snapshot (via `current-snapshot-id`) even when it falls outside that
window.

```yaml
{"description": "Contains standard Apache Iceberg table metadata (format version 1 or 2).\nSee https://iceberg.apache.org/spec/#table-metadata-fields for the full\nfield reference. Uses the `s3://` URI scheme for S3 compatibility,\nregardless of the underlying R2 backend.\n\nLimits the `snapshots`, `snapshot-log`, and `metadata-log` arrays to the\nmost recent 10 entries (by `timestamp-ms`). Always keeps the current\nsnapshot (via `current-snapshot-id`) even when it falls outside that\nwindow.\n", "type": "object", "example": {"current-schema-id": 1, "current-snapshot-id": 3055729675574597000, "default-sort-order-id": 0, "default-spec-id": 0, "format-version": 2, "last-column-id": 3, "last-sequence-number": 34, "last-updated-ms": 1602638573590, "location": "s3://my-bucket/__r2_data_catalog/wh-id/table-id", "metadata-log": [], "partition-specs": [{"fields": [{"field-id": 1000, "name": "user_id_bucket", "source-id": 1, "transform": "bucket[16]"}], "spec-id": 0}], "properties": {"write.parquet.compression-codec": "zstd"}, "refs": {"main": {"snapshot-id": 3055729675574597000, "type": "branch"}}, "schemas": [{"fields": [{"id": 1, "name": "user_id", "required": true, "type": "long"}, {"id": 2, "name": "email", "required": true, "type": "string"}], "identifier-field-ids": [1], "schema-id": 1, "type": "struct"}], "snapshot-log": [{"snapshot-id": 3055729675574597000, "timestamp-ms": 1555100955770}], "snapshots": [{"manifest-list": "s3://my-bucket/__r2_data_catalog/wh-id/table-id/metadata/snap-3055729675574597004.avro", "schema-id": 1, "sequence-number": 1, "snapshot-id": 3055729675574597000, "summary": {"operation": "append"}, "timestamp-ms": 1555100955770}], "sort-orders": [{"fields": [], "order-id": 0}], "table-uuid": "550e8400-e29b-41d4-a716-446655440000"}, "additionalProperties": true}
```
