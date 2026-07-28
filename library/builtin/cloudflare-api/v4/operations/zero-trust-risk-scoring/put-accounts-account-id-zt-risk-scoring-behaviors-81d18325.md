---
title: Update configuration for risk behaviors
page_id: operation-put-accounts-account-id-zt-risk-scoring-behaviors-f7a4a29e
path: operations/zero-trust-risk-scoring
description: Updates risk score behavior configurations, defining weights and thresholds for risk calculation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/zt_risk_scoring/behaviors
operation_ids:
    - dlp-risk-score-behaviors-put
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update configuration for risk behaviors

`PUT /accounts/{account_id}/zt_risk_scoring/behaviors`

Operation ID: `dlp-risk-score-behaviors-put`

Updates risk score behavior configurations, defining weights and thresholds for risk calculation.

## Definition

```yaml
{"operationId": "dlp-risk-score-behaviors-put", "summary": "Update configuration for risk behaviors", "description": "Updates risk score behavior configurations, defining weights and thresholds for risk calculation.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "Behaviors.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_UpdateBehaviors"}}}}, "responses": {"200": {"description": "Dataset created successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_UpdateBehaviors"}}, "type": "object"}]}}}}, "4XX": {"description": "Dataset creation failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Risk Scoring"], "x-api-token-group": ["Zero Trust Write"]}
```
