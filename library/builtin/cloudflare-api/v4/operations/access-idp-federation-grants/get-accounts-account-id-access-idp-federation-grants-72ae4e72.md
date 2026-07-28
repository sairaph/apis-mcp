---
title: List IdP federation grants
page_id: operation-get-accounts-account-id-access-idp-federation-grants-33009abb
path: operations/access-idp-federation-grants
description: Lists the IdP federation grants owned by the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/idp_federation_grants
operation_ids:
    - access-idp-federation-grants-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List IdP federation grants

`GET /accounts/{account_id}/access/idp_federation_grants`

Operation ID: `access-idp-federation-grants-list`

Lists the IdP federation grants owned by the account.

## Definition

```yaml
{"operationId": "access-idp-federation-grants-list", "summary": "List IdP federation grants", "description": "Lists the IdP federation grants owned by the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "List IdP federation grants response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_idp_federation_grant_list_response"}}}}, "4XX": {"description": "List IdP federation grants response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access IdP federation grants"], "x-stability": "beta"}
```
