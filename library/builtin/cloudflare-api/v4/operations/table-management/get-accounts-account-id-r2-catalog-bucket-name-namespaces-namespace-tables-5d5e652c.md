---
title: List tables in namespace
page_id: operation-get-accounts-account-id-r2-catalog-bucket-name-namespaces-namespace-tabl-ca3a5b06
path: operations/table-management
description: |-
    Returns a list of tables in the specified namespace within an R2 catalog.
    Supports pagination for efficient traversal of large table collections.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2-catalog/{bucket_name}/namespaces/{namespace}/tables
operation_ids:
    - list-tables
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List tables in namespace

`GET /accounts/{account_id}/r2-catalog/{bucket_name}/namespaces/{namespace}/tables`

Operation ID: `list-tables`

Returns a list of tables in the specified namespace within an R2 catalog.
Supports pagination for efficient traversal of large table collections.

## Definition

```yaml
{"operationId": "list-tables", "summary": "List tables in namespace", "description": "Returns a list of tables in the specified namespace within an R2 catalog.\nSupports pagination for efficient traversal of large table collections.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Identifies the account.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_account-id"}}, {"name": "bucket_name", "in": "path", "description": "Specifies the R2 bucket name.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_bucket-name"}}, {"name": "namespace", "in": "path", "description": "The namespace identifier.\nFor nested namespaces, use %1F as separator (e.g., \"bronze%1Fanalytics\").\n", "required": true, "schema": {"type": "string"}, "example": "bronze"}, {"name": "page_token", "in": "query", "description": "Opaque pagination token from a previous response.\nUse this to fetch the next page of results.\n", "schema": {"type": "string"}, "example": "MSYxNzU5NzU2MTI4NTU2Njk2JjAxOTliOWEzLTkxMmUtN2ZhMS05YzllLTg5MTAxMGQzYTg0MQ"}, {"name": "page_size", "in": "query", "description": "Maximum number of tables to return per page.\nDefaults to 100, maximum 1000.\n", "schema": {"type": "integer", "default": 100, "maximum": 1000, "minimum": 1}, "example": 100}, {"name": "return_uuids", "in": "query", "description": "Whether to include table UUIDs in the response.\nSet to true to receive the table_uuids array.\n", "schema": {"type": "boolean", "default": false}, "example": true}, {"name": "return_details", "in": "query", "description": "Whether to include additional metadata (timestamps, locations).\nWhen true, response includes created_at, updated_at, metadata_locations, and locations arrays.\n", "schema": {"type": "boolean", "default": false}, "example": true}], "responses": {"200": {"description": "List of tables retrieved successfully.", "content": {"application/json": {"example": {"errors": [], "messages": [], "result": {"details": [{"created_at": "2025-10-07T10:00:00Z", "identifier": {"name": "events", "namespace": ["bronze"]}, "location": "s3://my-bucket/__r2_data_catalog/wh-id/table-id/", "metadata_location": "s3://my-bucket/__r2_data_catalog/wh-id/table-id/metadata/v1.metadata.json", "table_uuid": "0199b9a1-28a0-71e0-a73e-b0fc32c8468e", "updated_at": "2025-10-07T15:00:00Z"}, {"created_at": "2025-10-07T10:30:00Z", "identifier": {"name": "users", "namespace": ["bronze"]}, "location": "s3://my-bucket/__r2_data_catalog/wh-id/table-id-2/", "metadata_location": "s3://my-bucket/__r2_data_catalog/wh-id/table-id-2/metadata/v2.metadata.json", "table_uuid": "0199b9a1-3c74-7731-bf53-d8c67ead079d", "updated_at": "2025-10-07T16:00:00Z"}], "identifiers": [{"name": "events", "namespace": ["bronze"]}, {"name": "users", "namespace": ["bronze"]}], "next_page_token": null, "table_uuids": ["0199b9a1-28a0-71e0-a73e-b0fc32c8468e", "0199b9a1-3c74-7731-bf53-d8c67ead079d"]}, "success": true}, "schema": {"allOf": [{"$ref": "#/components/schemas/r2-data-catalog_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/r2-data-catalog_table-list-response"}}, "type": "object"}]}}}}, "400": {"description": "Bad request (e.g., invalid page_size, malformed namespace).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "401": {"description": "Authentication failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "404": {"description": "Catalog or namespace not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Table Management"], "x-api-token-group": ["Workers R2 Data Catalog Write", "Workers R2 Data Catalog Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2-data-catalog.namespaces.tables", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
