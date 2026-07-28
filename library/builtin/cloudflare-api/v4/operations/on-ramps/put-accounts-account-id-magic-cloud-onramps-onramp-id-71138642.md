---
title: Update On-ramp
page_id: operation-put-accounts-account-id-magic-cloud-onramps-onramp-id-1310956b
path: operations/on-ramps
description: Update an On-ramp (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/magic/cloud/onramps/{onramp_id}
operation_ids:
    - onramps-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update On-ramp

`PUT /accounts/{account_id}/magic/cloud/onramps/{onramp_id}`

Operation ID: `onramps-update`

Update an On-ramp (Closed Beta).

## Definition

```yaml
{"operationId": "onramps-update", "summary": "Update On-ramp", "description": "Update an On-ramp (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "onramp_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_onramp_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_update_onramp_request"}}}}, "responses": {"200": {"description": "OK.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_update_onramp_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "404": {"description": "Not Found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "409": {"description": "Conflict.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "422": {"description": "Unprocessable Entity.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["On-ramps"], "x-api-token-group": ["Magic WAN Write"]}
```
