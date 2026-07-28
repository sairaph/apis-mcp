---
title: Apply On-ramp
page_id: operation-post-accounts-account-id-magic-cloud-onramps-onramp-id-apply-e9ae9cea
path: operations/on-ramps
description: Apply an On-ramp (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/cloud/onramps/{onramp_id}/apply
operation_ids:
    - onramps-apply
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Apply On-ramp

`POST /accounts/{account_id}/magic/cloud/onramps/{onramp_id}/apply`

Operation ID: `onramps-apply`

Apply an On-ramp (Closed Beta).

## Definition

```yaml
{"operationId": "onramps-apply", "summary": "Apply On-ramp", "description": "Apply an On-ramp (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "onramp_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_onramp_id"}}], "responses": {"202": {"description": "Accepted.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_good_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "404": {"description": "Not Found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "409": {"description": "Conflict.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["On-ramps"], "x-api-token-group": ["Magic WAN Write"]}
```
