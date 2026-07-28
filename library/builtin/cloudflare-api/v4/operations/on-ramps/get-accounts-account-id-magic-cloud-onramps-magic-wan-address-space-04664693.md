---
title: Read Magic WAN Address Space
page_id: operation-get-accounts-account-id-magic-cloud-onramps-magic-wan-address-space-43706cc6
path: operations/on-ramps
description: Read the Magic WAN Address Space (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/cloud/onramps/magic_wan_address_space
operation_ids:
    - onramps-mwan-addr-space-read
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read Magic WAN Address Space

`GET /accounts/{account_id}/magic/cloud/onramps/magic_wan_address_space`

Operation ID: `onramps-mwan-addr-space-read`

Read the Magic WAN Address Space (Closed Beta).

## Definition

```yaml
{"operationId": "onramps-mwan-addr-space-read", "summary": "Read Magic WAN Address Space", "description": "Read the Magic WAN Address Space (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}], "responses": {"200": {"description": "OK.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_get_magic_wan_address_space_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "404": {"description": "Not Found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["On-ramps"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read"]}
```
