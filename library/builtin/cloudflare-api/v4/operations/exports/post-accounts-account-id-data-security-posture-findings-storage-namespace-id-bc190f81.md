---
title: Create a finding instances export
page_id: operation-post-accounts-account-id-data-security-posture-findings-storage-namespac-67fe8129
path: operations/exports
description: |-
    Creates a CSV export for Finding instances and accepts optional filters in the payload.

    The `storage_namespace_id` path parameter is derived from the finding ID by base64-decoding it
    (which yields `integration_id:finding_type_id`) and replacing the colon with a hyphen.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/data-security/posture/findings/{storage_namespace_id}/instances/export
operation_ids:
    - CreateFindingInstancesExportCSV
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a finding instances export

`POST /accounts/{account_id}/data-security/posture/findings/{storage_namespace_id}/instances/export`

Operation ID: `CreateFindingInstancesExportCSV`

Creates a CSV export for Finding instances and accepts optional filters in the payload.

The `storage_namespace_id` path parameter is derived from the finding ID by base64-decoding it
(which yields `integration_id:finding_type_id`) and replacing the colon with a hyphen.

## Definition

```yaml
{"operationId": "CreateFindingInstancesExportCSV", "summary": "Create a finding instances export", "description": "Creates a CSV export for Finding instances and accepts optional filters in the payload.\n\nThe `storage_namespace_id` path parameter is derived from the finding ID by base64-decoding it\n(which yields `integration_id:finding_type_id`) and replacing the colon with a hyphen.\n", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_StorageNamespaceId"}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_FindingInstanceExportFilterRequest"}}}}, "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_export-job-response"}}}}, "400": {"description": "Bad Request: Invalid request parameters"}, "404": {"description": "Not Found: Finding not found"}}, "security": [{"api_token": []}], "tags": ["exports"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
