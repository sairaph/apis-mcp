---
title: Get risk score integration by reference id.
page_id: operation-get-accounts-account-id-zt-risk-scoring-integrations-reference-id-refere-68a1ed7b
path: operations/zero-trust-risk-scoring-integrations
description: Retrieves a Zero Trust risk score integration using its external reference ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/zt_risk_scoring/integrations/reference_id/{reference_id}
operation_ids:
    - dlp-zt-risk-score-integration-get-by-reference-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get risk score integration by reference id.

`GET /accounts/{account_id}/zt_risk_scoring/integrations/reference_id/{reference_id}`

Operation ID: `dlp-zt-risk-score-integration-get-by-reference-id`

Retrieves a Zero Trust risk score integration using its external reference ID.

## Definition

```yaml
{"operationId": "dlp-zt-risk-score-integration-get-by-reference-id", "summary": "Get risk score integration by reference id.", "description": "Retrieves a Zero Trust risk score integration using its external reference ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "reference_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Get response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_RiskScoreIntegration"}}, "type": "object"}]}}}}, "4XX": {"description": "Get failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Risk Scoring Integrations"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
