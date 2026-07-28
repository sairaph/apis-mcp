---
title: Update SSO connector state
page_id: operation-patch-accounts-account-id-sso-connectors-sso-connector-id-02594ee3
path: operations/sso
description: Updates the state or configuration of an SSO connector.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/sso_connectors/{sso_connector_id}
operation_ids:
    - update-sso-connector-state
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update SSO connector state

`PATCH /accounts/{account_id}/sso_connectors/{sso_connector_id}`

Operation ID: `update-sso-connector-state`

Updates the state or configuration of an SSO connector.

## Definition

```yaml
{"operationId": "update-sso-connector-state", "summary": "Update SSO connector state", "description": "Updates the state or configuration of an SSO connector.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "sso_connector_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_sso_connector_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"enabled": {"description": "SSO Connector enabled state", "type": "boolean", "example": true}, "use_fedramp_language": {"$ref": "#/components/schemas/iam_use_fedramp_language"}}}}}}, "responses": {"200": {"description": "Update SSO connector state response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_sso_connector_response"}}}}, "4XX": {"description": "Update SSO connector state response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["SSO"], "x-api-token-group": ["SSO Connector Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
