---
title: Create OAuth Client
page_id: operation-post-accounts-account-id-oauth-clients-760ff85b
path: operations/oauth-clients
description: Create a new OAuth client for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/oauth_clients
operation_ids:
    - oauth-clients-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create OAuth Client

`POST /accounts/{account_id}/oauth_clients`

Operation ID: `oauth-clients-create`

Create a new OAuth client for an account.

## Definition

```yaml
{"operationId": "oauth-clients-create", "summary": "Create OAuth Client", "description": "Create a new OAuth client for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_oauth_client_create_request"}}}}, "responses": {"200": {"description": "Create OAuth Client response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_oauth_client_create_response"}}}}, "4XX": {"description": "Create OAuth Client response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["OAuth Clients"], "x-api-token-group": ["OAuth Client Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.oauth-client.create"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
