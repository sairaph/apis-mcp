---
title: Delete SSO connector
page_id: operation-delete-accounts-account-id-sso-connectors-sso-connector-id-0786fee5
path: operations/sso
description: Deletes an SSO connector from the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/sso_connectors/{sso_connector_id}
operation_ids:
    - delete-sso-connector
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete SSO connector

`DELETE /accounts/{account_id}/sso_connectors/{sso_connector_id}`

Operation ID: `delete-sso-connector`

Deletes an SSO connector from the account.

## Definition

```yaml
{"operationId": "delete-sso-connector", "summary": "Delete SSO connector", "description": "Deletes an SSO connector from the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "sso_connector_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_sso_connector_identifier"}}], "responses": {"200": {"description": "Delete SSO connector response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-single-id"}}}}, "4XX": {"description": "Delete SSO connector response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["SSO"], "x-api-token-group": ["SSO Connector Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
