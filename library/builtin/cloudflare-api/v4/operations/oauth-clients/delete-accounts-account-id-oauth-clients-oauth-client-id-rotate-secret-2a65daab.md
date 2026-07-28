---
title: Delete Rotated OAuth Client Secret
page_id: operation-delete-accounts-account-id-oauth-clients-oauth-client-id-rotate-secret-d4f2316a
path: operations/oauth-clients
description: Removes the old client secret after a rotation, keeping only the new one. Use this after you have updated your client configuration to use the new secret. The `has_rotated_secret` field on the client indicates whether there is an old secret to delete.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/oauth_clients/{oauth_client_id}/rotate_secret
operation_ids:
    - oauth-clients-delete-rotated-secret
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Rotated OAuth Client Secret

`DELETE /accounts/{account_id}/oauth_clients/{oauth_client_id}/rotate_secret`

Operation ID: `oauth-clients-delete-rotated-secret`

Removes the old client secret after a rotation, keeping only the new one. Use this after you have updated your client configuration to use the new secret. The `has_rotated_secret` field on the client indicates whether there is an old secret to delete.

## Definition

```yaml
{"operationId": "oauth-clients-delete-rotated-secret", "summary": "Delete Rotated OAuth Client Secret", "description": "Removes the old client secret after a rotation, keeping only the new one. Use this after you have updated your client configuration to use the new secret. The `has_rotated_secret` field on the client indicates whether there is an old secret to delete.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "oauth_client_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_oauth_client_identifier"}}], "responses": {"200": {"description": "Delete Rotated OAuth Client Secret response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-single-id"}}}}, "4XX": {"description": "Delete Rotated OAuth Client Secret response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["OAuth Clients"], "x-api-token-group": ["OAuth Client Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.oauth-client.delete"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
