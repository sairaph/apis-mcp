---
title: Run Discovery for All Integrations
page_id: operation-post-accounts-account-id-magic-cloud-providers-discover-642cb045
path: operations/cloud-integrations
description: Run discovery for all Cloud Integrations in an account (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/cloud/providers/discover
operation_ids:
    - providers-discover-all
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Run Discovery for All Integrations

`POST /accounts/{account_id}/magic/cloud/providers/discover`

Operation ID: `providers-discover-all`

Run discovery for all Cloud Integrations in an account (Closed Beta).

## Definition

```yaml
{"operationId": "providers-discover-all", "summary": "Run Discovery for All Integrations", "description": "Run discovery for all Cloud Integrations in an account (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}], "responses": {"202": {"description": "Accepted.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_good_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "409": {"description": "Conflict.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Cloud Integrations"], "x-api-token-group": ["Magic WAN Write"]}
```
