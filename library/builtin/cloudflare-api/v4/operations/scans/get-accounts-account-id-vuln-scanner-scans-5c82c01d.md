---
title: List Scans
page_id: operation-get-accounts-account-id-vuln-scanner-scans-0f856236
path: operations/scans
description: Returns all scans for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/scans
operation_ids:
    - list-scans
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Scans

`GET /accounts/{account_id}/vuln_scanner/scans`

Operation ID: `list-scans`

Returns all scans for the account.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}]
```

## Definition

```yaml
{"operationId": "list-scans", "summary": "List Scans", "description": "Returns all scans for the account.", "parameters": [{"$ref": "#/components/parameters/vuln_scanner_page"}, {"$ref": "#/components/parameters/vuln_scanner_per_page"}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/vuln_scanner_scan"}}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Scans"]}
```
