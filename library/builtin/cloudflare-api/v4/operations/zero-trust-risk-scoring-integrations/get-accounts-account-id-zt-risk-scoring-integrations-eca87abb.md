---
title: List all risk score integrations for the account.
page_id: operation-get-accounts-account-id-zt-risk-scoring-integrations-92d33e35
path: operations/zero-trust-risk-scoring-integrations
description: Lists all configured Zero Trust risk score integrations for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/zt_risk_scoring/integrations
operation_ids:
    - dlp-zt-risk-score-integration-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all risk score integrations for the account.

`GET /accounts/{account_id}/zt_risk_scoring/integrations`

Operation ID: `dlp-zt-risk-score-integration-list`

Lists all configured Zero Trust risk score integrations for the account.

## Definition

```yaml
{"operationId": "dlp-zt-risk-score-integration-list", "summary": "List all risk score integrations for the account.", "description": "Lists all configured Zero Trust risk score integrations for the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "List response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_RiskScoreIntegrationArray"}}, "type": "object"}]}}}}, "4XX": {"description": "List failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Risk Scoring Integrations"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
