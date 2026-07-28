---
title: Read On-ramp
page_id: operation-get-accounts-account-id-magic-cloud-onramps-onramp-id-b741b774
path: operations/on-ramps
description: Read an On-ramp (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/cloud/onramps/{onramp_id}
operation_ids:
    - onramps-read
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read On-ramp

`GET /accounts/{account_id}/magic/cloud/onramps/{onramp_id}`

Operation ID: `onramps-read`

Read an On-ramp (Closed Beta).

## Definition

```yaml
{"operationId": "onramps-read", "summary": "Read On-ramp", "description": "Read an On-ramp (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "onramp_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_onramp_id"}}, {"name": "status", "in": "query", "schema": {"type": "boolean"}}, {"name": "vpcs", "in": "query", "schema": {"type": "boolean"}}, {"name": "post_apply_resources", "in": "query", "schema": {"type": "boolean"}}, {"name": "planned_resources", "in": "query", "schema": {"type": "boolean"}}], "responses": {"200": {"description": "OK.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_get_onramp_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "404": {"description": "Not Found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["On-ramps"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read"]}
```
