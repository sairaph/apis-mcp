---
title: Get Recent Account Scans
page_id: operation-get-accounts-account-id-security-center-insights-scans-b4e68f7b
path: operations/security-center-scans
description: Returns the most recent on-demand scans for the account, up to a maximum of 5. Each scan includes its ID, start time, and current status. This includes both account-wide and zone-scoped scans. Also returns quota information showing how many scans have been used and how many remain in the current 24-hour window.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/security-center/insights/scans
operation_ids:
    - get-security-center-account-scans
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Recent Account Scans

`GET /accounts/{account_id}/security-center/insights/scans`

Operation ID: `get-security-center-account-scans`

Returns the most recent on-demand scans for the account, up to a maximum of 5. Each scan includes its ID, start time, and current status. This includes both account-wide and zone-scoped scans. Also returns quota information showing how many scans have been used and how many remain in the current 24-hour window.

## Definition

```yaml
{"operationId": "get-security-center-account-scans", "summary": "Get Recent Account Scans", "description": "Returns the most recent on-demand scans for the account, up to a maximum of 5. Each scan includes its ID, start time, and current status. This includes both account-wide and zone-scoped scans. Also returns quota information showing how many scans have been used and how many remain in the current 24-hour window.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/security-center_identifier"}}], "responses": {"200": {"description": "The request was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/security-center_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/security-center_scansListResponse"}}, "type": "object"}]}}}}, "4XX": {"description": "A client error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/security-center_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["Security Center Scans"]}
```
