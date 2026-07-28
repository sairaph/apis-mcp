---
title: Get an IdP federation grant
page_id: operation-get-accounts-account-id-access-idp-federation-grants-grant-id-d8cfcbf8
path: operations/access-idp-federation-grants
description: Retrieves a single IdP federation grant by its UID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/idp_federation_grants/{grant_id}
operation_ids:
    - access-idp-federation-grants-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an IdP federation grant

`GET /accounts/{account_id}/access/idp_federation_grants/{grant_id}`

Operation ID: `access-idp-federation-grants-get`

Retrieves a single IdP federation grant by its UID.

## Definition

```yaml
{"operationId": "access-idp-federation-grants-get", "summary": "Get an IdP federation grant", "description": "Retrieves a single IdP federation grant by its UID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "grant_id", "in": "path", "description": "UID of the IdP federation grant.", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get IdP federation grant response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_idp_federation_grant_response"}}}}, "404": {"description": "IdP federation grant does not exist", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}, "4XX": {"description": "Get IdP federation grant response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access IdP federation grants"], "x-stability": "beta"}
```
