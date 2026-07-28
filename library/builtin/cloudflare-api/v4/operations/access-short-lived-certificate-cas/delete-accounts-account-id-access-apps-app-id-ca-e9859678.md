---
title: Delete a short-lived certificate CA
page_id: operation-delete-accounts-account-id-access-apps-app-id-ca-7725c59f
path: operations/access-short-lived-certificate-cas
description: Deletes a short-lived certificate CA.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/access/apps/{app_id}/ca
operation_ids:
    - access-short-lived-certificate-c-as-delete-a-short-lived-certificate-ca
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a short-lived certificate CA

`DELETE /accounts/{account_id}/access/apps/{app_id}/ca`

Operation ID: `access-short-lived-certificate-c-as-delete-a-short-lived-certificate-ca`

Deletes a short-lived certificate CA.

## Definition

```yaml
{"operationId": "access-short-lived-certificate-c-as-delete-a-short-lived-certificate-ca", "summary": "Delete a short-lived certificate CA", "description": "Deletes a short-lived certificate CA.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"202": {"description": "Delete a short-lived certificate CA response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_id_response-2"}}}}, "4XX": {"description": "Delete a short-lived certificate CA response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access short-lived certificate CAs"], "x-api-token-group": ["Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications.cas", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
