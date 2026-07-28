---
title: Get all behaviors and associated configuration
page_id: operation-get-accounts-account-id-zt-risk-scoring-behaviors-6cd4d7b6
path: operations/zero-trust-risk-scoring
description: Retrieves configured risk score behaviors that define how user actions affect their overall risk score.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/zt_risk_scoring/behaviors
operation_ids:
    - dlp-risk-score-behaviors-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get all behaviors and associated configuration

`GET /accounts/{account_id}/zt_risk_scoring/behaviors`

Operation ID: `dlp-risk-score-behaviors-get`

Retrieves configured risk score behaviors that define how user actions affect their overall risk score.

## Definition

```yaml
{"operationId": "dlp-risk-score-behaviors-get", "summary": "Get all behaviors and associated configuration", "description": "Retrieves configured risk score behaviors that define how user actions affect their overall risk score.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Risk scoring behaviors.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Behaviors"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to get risk scoring behaviors.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Risk Scoring"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
