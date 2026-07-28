---
title: OAuth Client Details
page_id: operation-get-accounts-account-id-oauth-clients-oauth-client-id-fa75fb38
path: operations/oauth-clients
description: Get details of a specific OAuth client.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/oauth_clients/{oauth_client_id}
operation_ids:
    - oauth-clients-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# OAuth Client Details

`GET /accounts/{account_id}/oauth_clients/{oauth_client_id}`

Operation ID: `oauth-clients-get`

Get details of a specific OAuth client.

## Definition

```yaml
{"operationId": "oauth-clients-get", "summary": "OAuth Client Details", "description": "Get details of a specific OAuth client.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "oauth_client_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_oauth_client_identifier"}}], "responses": {"200": {"description": "OAuth Client Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_oauth_client_response"}}}}, "4XX": {"description": "OAuth Client Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["OAuth Clients"], "x-api-token-group": ["OAuth Client Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.oauth-client.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
