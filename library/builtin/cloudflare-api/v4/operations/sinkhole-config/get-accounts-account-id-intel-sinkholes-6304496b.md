---
title: List sinkholes owned by this account
page_id: operation-get-accounts-account-id-intel-sinkholes-b3536486
path: operations/sinkhole-config
description: Lists sinkholes owned by the account for redirecting malicious traffic.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/sinkholes
operation_ids:
    - sinkhole-config-list-sinkholes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List sinkholes owned by this account

`GET /accounts/{account_id}/intel/sinkholes`

Operation ID: `sinkhole-config-list-sinkholes`

Lists sinkholes owned by the account for redirecting malicious traffic.

## Definition

```yaml
{"operationId": "sinkhole-config-list-sinkholes", "summary": "List sinkholes owned by this account", "description": "Lists sinkholes owned by the account for redirecting malicious traffic.", "parameters": [{"$ref": "#/components/parameters/intel-sinkholes_account_id"}], "responses": {"200": {"description": "Successful Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_sinkhole_collection_response"}}}}, "4XX": {"description": "Error Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_api_response_common_failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Sinkhole Config"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
