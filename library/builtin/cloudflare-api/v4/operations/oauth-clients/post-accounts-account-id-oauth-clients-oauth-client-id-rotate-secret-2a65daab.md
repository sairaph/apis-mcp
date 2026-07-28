---
title: Rotate OAuth Client Secret
page_id: operation-post-accounts-account-id-oauth-clients-oauth-client-id-rotate-secret-316fc3c4
path: operations/oauth-clients
description: Creates a second client secret so you can update your client configuration before deleting the old one. The `has_rotated_secret` field on the client will be set to `true`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/oauth_clients/{oauth_client_id}/rotate_secret
operation_ids:
    - oauth-clients-rotate-secret
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Rotate OAuth Client Secret

`POST /accounts/{account_id}/oauth_clients/{oauth_client_id}/rotate_secret`

Operation ID: `oauth-clients-rotate-secret`

Creates a second client secret so you can update your client configuration before deleting the old one. The `has_rotated_secret` field on the client will be set to `true`.

## Definition

```yaml
{"operationId": "oauth-clients-rotate-secret", "summary": "Rotate OAuth Client Secret", "description": "Creates a second client secret so you can update your client configuration before deleting the old one. The `has_rotated_secret` field on the client will be set to `true`.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "oauth_client_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_oauth_client_identifier"}}], "responses": {"200": {"description": "Rotate OAuth Client Secret response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_oauth_client_rotate_secret_response"}}}}, "4XX": {"description": "Rotate OAuth Client Secret response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["OAuth Clients"], "x-api-token-group": ["OAuth Client Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.oauth-client.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
