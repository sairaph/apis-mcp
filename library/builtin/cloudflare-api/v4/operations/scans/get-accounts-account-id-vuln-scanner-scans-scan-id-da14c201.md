---
title: Get Scan
page_id: operation-get-accounts-account-id-vuln-scanner-scans-scan-id-05682c44
path: operations/scans
description: Returns a single scan by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/scans/{scan_id}
operation_ids:
    - get-scan
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Scan

`GET /accounts/{account_id}/vuln_scanner/scans/{scan_id}`

Operation ID: `get-scan`

Returns a single scan by ID.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}, {"$ref": "#/components/parameters/vuln_scanner_scan_id"}]
```

## Definition

```yaml
{"operationId": "get-scan", "summary": "Get Scan", "description": "Returns a single scan by ID.", "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/vuln_scanner_scan"}, "result_info": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Scans"]}
```
