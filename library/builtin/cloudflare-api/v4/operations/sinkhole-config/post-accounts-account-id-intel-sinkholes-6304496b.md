---
title: Create a new sinkhole for your account
page_id: operation-post-accounts-account-id-intel-sinkholes-1467d8d2
path: operations/sinkhole-config
description: Create a new sinkhole. Logs of large request bodies will be truncated, but the full request body can be recorded in R2. If you wish to record large request bodies in R2, include the R2 key ID, key secret, and bucket name in the request body.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/intel/sinkholes
operation_ids:
    - sinkhole-config-create-sinkhole
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new sinkhole for your account

`POST /accounts/{account_id}/intel/sinkholes`

Operation ID: `sinkhole-config-create-sinkhole`

Create a new sinkhole. Logs of large request bodies will be truncated, but the full request body can be recorded in R2. If you wish to record large request bodies in R2, include the R2 key ID, key secret, and bucket name in the request body.

## Definition

```yaml
{"operationId": "sinkhole-config-create-sinkhole", "summary": "Create a new sinkhole for your account", "description": "Create a new sinkhole. Logs of large request bodies will be truncated, but the full request body can be recorded in R2. If you wish to record large request bodies in R2, include the R2 key ID, key secret, and bucket name in the request body.", "parameters": [{"$ref": "#/components/parameters/intel-sinkholes_account_id"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_sinkhole_create_params"}}}}, "responses": {"201": {"description": "Sinkhole Created", "headers": {"Location": {"description": "URI of the created sinkhole", "schema": {"type": "string"}}}, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_sinkhole_single_response"}}}}, "4XX": {"description": "Error Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_api_response_common_failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Sinkhole Config"], "x-api-token-group": ["Intel Write"]}
```
