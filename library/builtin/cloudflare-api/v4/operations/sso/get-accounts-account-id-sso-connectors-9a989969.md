---
title: Get all SSO connectors
page_id: operation-get-accounts-account-id-sso-connectors-dae6a9b4
path: operations/sso
description: Lists all SSO connectors configured for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/sso_connectors
operation_ids:
    - get-all-sso-connectors
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get all SSO connectors

`GET /accounts/{account_id}/sso_connectors`

Operation ID: `get-all-sso-connectors`

Lists all SSO connectors configured for the account.

## Definition

```yaml
{"operationId": "get-all-sso-connectors", "summary": "Get all SSO connectors", "description": "Lists all SSO connectors configured for the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "responses": {"200": {"description": "Get all SSO connectors response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_sso_connector_collection_response"}}}}, "4XX": {"description": "Get all SSO connectors response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["SSO"], "x-api-token-group": ["SSO Connector Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
