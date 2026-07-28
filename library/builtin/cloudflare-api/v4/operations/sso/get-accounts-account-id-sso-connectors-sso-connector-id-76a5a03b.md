---
title: Get single SSO connector
page_id: operation-get-accounts-account-id-sso-connectors-sso-connector-id-7f3fbb63
path: operations/sso
description: Retrieves details for a specific SSO connector.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/sso_connectors/{sso_connector_id}
operation_ids:
    - get-sso-connector
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get single SSO connector

`GET /accounts/{account_id}/sso_connectors/{sso_connector_id}`

Operation ID: `get-sso-connector`

Retrieves details for a specific SSO connector.

## Definition

```yaml
{"operationId": "get-sso-connector", "summary": "Get single SSO connector", "description": "Retrieves details for a specific SSO connector.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "sso_connector_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_sso_connector_identifier"}}], "responses": {"200": {"description": "Get SSO connector response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_sso_connector_response"}}}}, "4XX": {"description": "Get SSO connector response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["SSO"], "x-api-token-group": ["SSO Connector Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
