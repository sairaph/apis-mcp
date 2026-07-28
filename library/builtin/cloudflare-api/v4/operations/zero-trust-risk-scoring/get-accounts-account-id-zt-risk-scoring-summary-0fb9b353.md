---
title: Get risk score info for all users in the account
page_id: operation-get-accounts-account-id-zt-risk-scoring-summary-22b09077
path: operations/zero-trust-risk-scoring
description: Gets an aggregate summary of risk scores across the account, including distribution and trends.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/zt_risk_scoring/summary
operation_ids:
    - dlp-risk-score-summary-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get risk score info for all users in the account

`GET /accounts/{account_id}/zt_risk_scoring/summary`

Operation ID: `dlp-risk-score-summary-get`

Gets an aggregate summary of risk scores across the account, including distribution and trends.

## Definition

```yaml
{"operationId": "dlp-risk-score-summary-get", "summary": "Get risk score info for all users in the account", "description": "Gets an aggregate summary of risk scores across the account, including distribution and trends.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Risk score for all users in the account.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_RiskSummary"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to get risk scores.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Risk Scoring"], "x-api-token-group": ["Zero Trust: PII Read"]}
```
