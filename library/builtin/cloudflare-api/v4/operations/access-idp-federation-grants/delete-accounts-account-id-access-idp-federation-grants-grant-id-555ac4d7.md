---
title: Delete an IdP federation grant
page_id: operation-delete-accounts-account-id-access-idp-federation-grants-grant-id-aeea9031
path: operations/access-idp-federation-grants
description: |-
    Deletes an IdP federation grant. The identity provider remains in the account,
    but it is no longer available for federation to other accounts.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/access/idp_federation_grants/{grant_id}
operation_ids:
    - access-idp-federation-grants-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an IdP federation grant

`DELETE /accounts/{account_id}/access/idp_federation_grants/{grant_id}`

Operation ID: `access-idp-federation-grants-delete`

Deletes an IdP federation grant. The identity provider remains in the account,
but it is no longer available for federation to other accounts.

## Definition

```yaml
{"operationId": "access-idp-federation-grants-delete", "summary": "Delete an IdP federation grant", "description": "Deletes an IdP federation grant. The identity provider remains in the account,\nbut it is no longer available for federation to other accounts.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "grant_id", "in": "path", "description": "UID of the IdP federation grant.", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"202": {"description": "Delete IdP federation grant response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_idp_federation_grant_id_response"}}}}, "404": {"description": "IdP federation grant does not exist", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}, "4XX": {"description": "Delete IdP federation grant response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access IdP federation grants"]}
```
