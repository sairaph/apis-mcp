---
title: List OAuth Clients
page_id: operation-get-accounts-account-id-oauth-clients-a91a146d
path: operations/oauth-clients
description: List all OAuth clients for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/oauth_clients
operation_ids:
    - oauth-clients-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List OAuth Clients

`GET /accounts/{account_id}/oauth_clients`

Operation ID: `oauth-clients-list`

List all OAuth clients for an account.

## Definition

```yaml
{"operationId": "oauth-clients-list", "summary": "List OAuth Clients", "description": "List all OAuth clients for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "responses": {"200": {"description": "List OAuth Clients response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_collection_oauth_client_response"}}}}, "4XX": {"description": "List OAuth Clients response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["OAuth Clients"], "x-api-token-group": ["OAuth Client Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.oauth-client.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
