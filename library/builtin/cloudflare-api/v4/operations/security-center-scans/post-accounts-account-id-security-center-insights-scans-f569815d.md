---
title: Start On-Demand Account Scan
page_id: operation-post-accounts-account-id-security-center-insights-scans-631bb82a
path: operations/security-center-scans
description: Initiates an on-demand security scan for the entire account, scanning all zones associated with the account. Rate limited to 5 scans per account per 24-hour window.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/security-center/insights/scans
operation_ids:
    - start-security-center-account-scan
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Start On-Demand Account Scan

`POST /accounts/{account_id}/security-center/insights/scans`

Operation ID: `start-security-center-account-scan`

Initiates an on-demand security scan for the entire account, scanning all zones associated with the account. Rate limited to 5 scans per account per 24-hour window.

## Definition

```yaml
{"operationId": "start-security-center-account-scan", "summary": "Start On-Demand Account Scan", "description": "Initiates an on-demand security scan for the entire account, scanning all zones associated with the account. Rate limited to 5 scans per account per 24-hour window.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_newScanRequest"}}}}, "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/security-center_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/security-center_newScanResponse"}}, "type": "object"}]}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Scans"]}
```
