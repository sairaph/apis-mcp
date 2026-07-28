---
title: Delete OAuth Client
page_id: operation-delete-accounts-account-id-oauth-clients-oauth-client-id-6c075c8d
path: operations/oauth-clients
description: Delete an OAuth client.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/oauth_clients/{oauth_client_id}
operation_ids:
    - oauth-clients-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete OAuth Client

`DELETE /accounts/{account_id}/oauth_clients/{oauth_client_id}`

Operation ID: `oauth-clients-delete`

Delete an OAuth client.

## Definition

```yaml
{"operationId": "oauth-clients-delete", "summary": "Delete OAuth Client", "description": "Delete an OAuth client.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "oauth_client_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_oauth_client_identifier"}}], "responses": {"200": {"description": "Delete OAuth Client response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-single-id"}}}}, "4XX": {"description": "Delete OAuth Client response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["OAuth Clients"], "x-api-token-group": ["OAuth Client Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.oauth-client.delete"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
