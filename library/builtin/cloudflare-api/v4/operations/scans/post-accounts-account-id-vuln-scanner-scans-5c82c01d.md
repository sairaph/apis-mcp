---
title: Create Scan
page_id: operation-post-accounts-account-id-vuln-scanner-scans-c165cf79
path: operations/scans
description: |-
    Creates and starts a new vulnerability scan. The response may include
    non-fatal warnings in the `messages` array.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/scans
operation_ids:
    - create-scan
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Scan

`POST /accounts/{account_id}/vuln_scanner/scans`

Operation ID: `create-scan`

Creates and starts a new vulnerability scan. The response may include
non-fatal warnings in the `messages` array.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}]
```

## Definition

```yaml
{"operationId": "create-scan", "summary": "Create Scan", "description": "Creates and starts a new vulnerability scan. The response may include\nnon-fatal warnings in the `messages` array.\n", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vuln_scanner_create-scan-request"}}}}, "responses": {"200": {"description": "Successful response. Check the `messages` array for non-fatal\nwarnings that arose during scan creation.\n", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/vuln_scanner_scan"}, "result_info": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Scans"]}
```
