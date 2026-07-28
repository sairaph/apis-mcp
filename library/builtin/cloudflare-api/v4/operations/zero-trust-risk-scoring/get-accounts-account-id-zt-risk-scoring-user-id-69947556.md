---
title: Get risk event/score information for a specific user
page_id: operation-get-accounts-account-id-zt-risk-scoring-user-id-7c47f45e
path: operations/zero-trust-risk-scoring
description: Retrieves the detailed risk score breakdown for a specific user, including contributing factors.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/zt_risk_scoring/{user_id}
operation_ids:
    - dlp-risk-score-summary-get-for-user
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get risk event/score information for a specific user

`GET /accounts/{account_id}/zt_risk_scoring/{user_id}`

Operation ID: `dlp-risk-score-summary-get-for-user`

Retrieves the detailed risk score breakdown for a specific user, including contributing factors.

## Definition

```yaml
{"operationId": "dlp-risk-score-summary-get-for-user", "summary": "Get risk event/score information for a specific user", "description": "Retrieves the detailed risk score breakdown for a specific user, including contributing factors.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "user_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Risk events.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_RiskEvents"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to get risk events.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Risk Scoring"], "x-api-token-group": ["Zero Trust: PII Read"]}
```
