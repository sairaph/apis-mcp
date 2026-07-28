---
title: Create Zero Trust list from CSV
page_id: operation-post-accounts-account-id-gateway-lists-upload-5e253b00
path: operations/zero-trust-lists
description: Create a new Zero Trust list by uploading a CSV file. The file must be `text/csv` or `text/plain` and cannot exceed 2 MB. The operation is processed asynchronously. Use the returned operation ID to poll for status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/gateway/lists/upload
operation_ids:
    - zero-trust-lists-create-zero-trust-list-from-csv
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Zero Trust list from CSV

`POST /accounts/{account_id}/gateway/lists/upload`

Operation ID: `zero-trust-lists-create-zero-trust-list-from-csv`

Create a new Zero Trust list by uploading a CSV file. The file must be `text/csv` or `text/plain` and cannot exceed 2 MB. The operation is processed asynchronously. Use the returned operation ID to poll for status.

## Definition

```yaml
{"operationId": "zero-trust-lists-create-zero-trust-list-from-csv", "summary": "Create Zero Trust list from CSV", "description": "Create a new Zero Trust list by uploading a CSV file. The file must be `text/csv` or `text/plain` and cannot exceed 2 MB. The operation is processed asynchronously. Use the returned operation ID to poll for status.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"multipart/form-data": {"schema": {"type": "object", "properties": {"file": {"description": "The CSV file containing list items. Must be `text/csv` or `text/plain` and cannot exceed 2 MB.", "type": "string", "format": "binary"}}, "required": ["file"]}}}}, "responses": {"200": {"description": "Create Zero Trust list from CSV response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_csv_operation_single_response"}}}}, "4XX": {"description": "Create Zero Trust list from CSV response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_csv_operation_single_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust lists"]}
```
