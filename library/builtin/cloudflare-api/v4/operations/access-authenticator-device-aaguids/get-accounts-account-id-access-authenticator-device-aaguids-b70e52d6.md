---
title: List authenticator device AAGUIDs
page_id: operation-get-accounts-account-id-access-authenticator-device-aaguids-d0a5cd0c
path: operations/access-authenticator-device-aaguids
description: Returns a list of Authenticator Device AAGUIDs for MFA configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/authenticator_device_aaguids
operation_ids:
    - access-authenticator-device-aaguids-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List authenticator device AAGUIDs

`GET /accounts/{account_id}/access/authenticator_device_aaguids`

Operation ID: `access-authenticator-device-aaguids-list`

Returns a list of Authenticator Device AAGUIDs for MFA configuration.

## Definition

```yaml
{"operationId": "access-authenticator-device-aaguids-list", "summary": "List authenticator device AAGUIDs", "description": "Returns a list of Authenticator Device AAGUIDs for MFA configuration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "List authenticator device AAGUIDs response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-16"}}}}, "4XX": {"description": "List authenticator device AAGUIDs response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access Authenticator Device AAGUIDs"]}
```
