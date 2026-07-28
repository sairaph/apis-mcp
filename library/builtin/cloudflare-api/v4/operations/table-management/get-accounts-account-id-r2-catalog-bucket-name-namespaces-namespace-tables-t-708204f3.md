---
title: Get table details
page_id: operation-get-accounts-account-id-r2-catalog-bucket-name-namespaces-namespace-tabl-f6b8b26e
path: operations/table-management
description: |-
    Returns full Apache Iceberg metadata for a single table: schema,
    partition specs, sort orders, properties, and recent snapshot history.
    Designed for catalog introspection UIs that need per-table details
    without holding R2 credentials.

    The `metadata.snapshots`, `metadata.snapshot-log`, and
    `metadata.metadata-log` arrays are pruned to the most recent 10
    entries by `timestamp-ms`. Use `total_snapshots` and
    `returned_snapshots` to surface the truncation to end users.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2-catalog/{bucket_name}/namespaces/{namespace}/tables/{table_name}
operation_ids:
    - get-table
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get table details

`GET /accounts/{account_id}/r2-catalog/{bucket_name}/namespaces/{namespace}/tables/{table_name}`

Operation ID: `get-table`

Returns full Apache Iceberg metadata for a single table: schema,
partition specs, sort orders, properties, and recent snapshot history.
Designed for catalog introspection UIs that need per-table details
without holding R2 credentials.

The `metadata.snapshots`, `metadata.snapshot-log`, and
`metadata.metadata-log` arrays are pruned to the most recent 10
entries by `timestamp-ms`. Use `total_snapshots` and
`returned_snapshots` to surface the truncation to end users.

## Definition

```yaml
{"operationId": "get-table", "summary": "Get table details", "description": "Returns full Apache Iceberg metadata for a single table: schema,\npartition specs, sort orders, properties, and recent snapshot history.\nDesigned for catalog introspection UIs that need per-table details\nwithout holding R2 credentials.\n\nThe `metadata.snapshots`, `metadata.snapshot-log`, and\n`metadata.metadata-log` arrays are pruned to the most recent 10\nentries by `timestamp-ms`. Use `total_snapshots` and\n`returned_snapshots` to surface the truncation to end users.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Identifies the account.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_account-id"}}, {"name": "bucket_name", "in": "path", "description": "Specifies the R2 bucket name.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_bucket-name"}}, {"name": "namespace", "in": "path", "description": "The namespace identifier.\nFor nested namespaces, use %1F as separator (e.g., \"bronze%1Fanalytics\").\n", "required": true, "schema": {"type": "string"}, "example": "bronze"}, {"name": "table_name", "in": "path", "description": "The table name within the given namespace.", "required": true, "schema": {"type": "string"}, "example": "events"}], "responses": {"200": {"description": "Table details retrieved successfully.", "content": {"application/json": {"example": {"errors": [], "messages": [], "result": {"identifier": {"name": "events", "namespace": ["bronze"]}, "metadata": {"current-schema-id": 1, "current-snapshot-id": 3055729675574597000, "default-spec-id": 0, "format-version": 2, "location": "s3://my-bucket/__r2_data_catalog/wh-id/table-id", "partition-specs": [{"fields": [{"field-id": 1000, "name": "user_id_bucket", "source-id": 1, "transform": "bucket[16]"}], "spec-id": 0}], "properties": {"write.parquet.compression-codec": "zstd"}, "schemas": [{"type": "struct", "fields": [{"id": 1, "name": "user_id", "required": true, "type": "long"}], "schema-id": 1}], "snapshots": [{"sequence-number": 1, "snapshot-id": 3055729675574597000, "summary": {"operation": "append"}, "timestamp-ms": 1555100955770}], "table-uuid": "550e8400-e29b-41d4-a716-446655440000"}, "metadata_location": "s3://my-bucket/__r2_data_catalog/wh-id/table-id/metadata/00001-abc.metadata.json", "returned_snapshots": 10, "table_uuid": "550e8400-e29b-41d4-a716-446655440000", "total_snapshots": 122113}, "success": true}, "schema": {"allOf": [{"$ref": "#/components/schemas/r2-data-catalog_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/r2-data-catalog_get-table-response"}}, "type": "object"}]}}}}, "400": {"description": "Bad request (e.g., malformed namespace).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "401": {"description": "Authentication failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "403": {"description": "Forbidden — JWT claims do not match account.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "404": {"description": "Catalog, namespace, or table not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Table Management"], "x-api-token-group": ["Workers R2 Data Catalog Write", "Workers R2 Data Catalog Read"]}
```
