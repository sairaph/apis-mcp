---
title: Create new risk score integration.
page_id: operation-post-accounts-account-id-zt-risk-scoring-integrations-95d74bfe
path: operations/zero-trust-risk-scoring-integrations
description: Creates a new Zero Trust risk score integration, connecting external risk signals to Cloudflare's risk scoring system.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/zt_risk_scoring/integrations
operation_ids:
    - dlp-zt-risk-score-integration-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create new risk score integration.

`POST /accounts/{account_id}/zt_risk_scoring/integrations`

Operation ID: `dlp-zt-risk-score-integration-create`

Creates a new Zero Trust risk score integration, connecting external risk signals to Cloudflare's risk scoring system.

## Definition

```yaml
{"operationId": "dlp-zt-risk-score-integration-create", "summary": "Create new risk score integration.", "description": "Creates a new Zero Trust risk score integration, connecting external risk signals to Cloudflare's risk scoring system.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_CreateIntegrationBody"}}}}, "responses": {"200": {"description": "Create response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_RiskScoreIntegration"}}, "type": "object"}]}}}}, "4XX": {"description": "Create failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Risk Scoring Integrations"], "x-api-token-group": ["Zero Trust Write"]}
```
