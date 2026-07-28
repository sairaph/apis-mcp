---
title: Delete a risk score integration.
page_id: operation-delete-accounts-account-id-zt-risk-scoring-integrations-integration-id-2a658cf1
path: operations/zero-trust-risk-scoring-integrations
description: Removes a Zero Trust risk score integration, disconnecting the external risk signal source.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/zt_risk_scoring/integrations/{integration_id}
operation_ids:
    - dlp-zt-risk-score-integration-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a risk score integration.

`DELETE /accounts/{account_id}/zt_risk_scoring/integrations/{integration_id}`

Operation ID: `dlp-zt-risk-score-integration-delete`

Removes a Zero Trust risk score integration, disconnecting the external risk signal source.

## Definition

```yaml
{"operationId": "dlp-zt-risk-score-integration-delete", "summary": "Delete a risk score integration.", "description": "Removes a Zero Trust risk score integration, disconnecting the external risk signal source.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "integration_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Delete response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Empty"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Risk Scoring Integrations"], "x-api-token-group": ["Zero Trust Write"]}
```
