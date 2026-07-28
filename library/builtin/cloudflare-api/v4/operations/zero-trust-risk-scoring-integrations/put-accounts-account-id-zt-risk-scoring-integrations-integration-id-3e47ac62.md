---
title: Update a risk score integration.
page_id: operation-put-accounts-account-id-zt-risk-scoring-integrations-integration-id-14a2c833
path: operations/zero-trust-risk-scoring-integrations
description: Overwrite the reference_id, tenant_url, and active values with the ones provided.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/zt_risk_scoring/integrations/{integration_id}
operation_ids:
    - dlp-zt-risk-score-integration-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a risk score integration.

`PUT /accounts/{account_id}/zt_risk_scoring/integrations/{integration_id}`

Operation ID: `dlp-zt-risk-score-integration-update`

Overwrite the reference_id, tenant_url, and active values with the ones provided.

## Definition

```yaml
{"operationId": "dlp-zt-risk-score-integration-update", "summary": "Update a risk score integration.", "description": "Overwrite the reference_id, tenant_url, and active values with the ones provided.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "integration_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_UpdateIntegrationBody"}}}}, "responses": {"200": {"description": "Update response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_RiskScoreIntegration"}}, "type": "object"}]}}}}, "4XX": {"description": "Update failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Risk Scoring Integrations"], "x-api-token-group": ["Zero Trust Write"]}
```
