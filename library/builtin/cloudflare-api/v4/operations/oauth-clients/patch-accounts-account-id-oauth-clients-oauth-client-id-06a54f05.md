---
title: Update OAuth Client
page_id: operation-patch-accounts-account-id-oauth-clients-oauth-client-id-d74242f6
path: operations/oauth-clients
description: Update an existing OAuth client. Only include fields you want to update.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/oauth_clients/{oauth_client_id}
operation_ids:
    - oauth-clients-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update OAuth Client

`PATCH /accounts/{account_id}/oauth_clients/{oauth_client_id}`

Operation ID: `oauth-clients-update`

Update an existing OAuth client. Only include fields you want to update.

## Definition

```yaml
{"operationId": "oauth-clients-update", "summary": "Update OAuth Client", "description": "Update an existing OAuth client. Only include fields you want to update.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "oauth_client_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_oauth_client_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_oauth_client_update_request"}}}}, "responses": {"200": {"description": "Update OAuth Client response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_oauth_client_response"}}}}, "4XX": {"description": "Update OAuth Client response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["OAuth Clients"], "x-api-token-group": ["OAuth Client Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.oauth-client.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
