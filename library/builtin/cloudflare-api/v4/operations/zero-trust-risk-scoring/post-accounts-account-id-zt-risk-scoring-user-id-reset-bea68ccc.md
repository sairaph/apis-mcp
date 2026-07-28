---
title: Clear the risk score for a particular user
page_id: operation-post-accounts-account-id-zt-risk-scoring-user-id-reset-5ab7f232
path: operations/zero-trust-risk-scoring
description: Resets risk scores for specified users, clearing their accumulated risk history.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/zt_risk_scoring/{user_id}/reset
operation_ids:
    - dlp-risk-score-reset-post
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Clear the risk score for a particular user

`POST /accounts/{account_id}/zt_risk_scoring/{user_id}/reset`

Operation ID: `dlp-risk-score-reset-post`

Resets risk scores for specified users, clearing their accumulated risk history.

## Definition

```yaml
{"operationId": "dlp-risk-score-reset-post", "summary": "Clear the risk score for a particular user", "description": "Resets risk scores for specified users, clearing their accumulated risk history.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "user_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Dataset created successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Empty"}}, "type": "object"}]}}}}, "4XX": {"description": "Dataset creation failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Risk Scoring"], "x-api-token-group": ["Zero Trust Write"]}
```
