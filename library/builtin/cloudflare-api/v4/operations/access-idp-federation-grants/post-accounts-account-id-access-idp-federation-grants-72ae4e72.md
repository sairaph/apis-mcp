---
title: Create an IdP federation grant
page_id: operation-post-accounts-account-id-access-idp-federation-grants-b3ac53ca
path: operations/access-idp-federation-grants
description: |-
    Creates an IdP federation grant for the specified identity provider, making it
    available for federation to other accounts in the same Cloudflare organization.

    The account must belong to a Cloudflare organization. One-time pin and
    Cloudflare-managed identity providers cannot be federated. An account
    can federate at most five identity providers at a time.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/idp_federation_grants
operation_ids:
    - access-idp-federation-grants-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an IdP federation grant

`POST /accounts/{account_id}/access/idp_federation_grants`

Operation ID: `access-idp-federation-grants-create`

Creates an IdP federation grant for the specified identity provider, making it
available for federation to other accounts in the same Cloudflare organization.

The account must belong to a Cloudflare organization. One-time pin and
Cloudflare-managed identity providers cannot be federated. An account
can federate at most five identity providers at a time.

## Definition

```yaml
{"operationId": "access-idp-federation-grants-create", "summary": "Create an IdP federation grant", "description": "Creates an IdP federation grant for the specified identity provider, making it\navailable for federation to other accounts in the same Cloudflare organization.\n\nThe account must belong to a Cloudflare organization. One-time pin and\nCloudflare-managed identity providers cannot be federated. An account\ncan federate at most five identity providers at a time.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_idp_federation_grant_create_request"}}}}, "responses": {"201": {"description": "Create IdP federation grant response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_idp_federation_grant_response"}}}}, "4XX": {"description": "Create IdP federation grant response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access IdP federation grants"], "x-stability": "beta"}
```
