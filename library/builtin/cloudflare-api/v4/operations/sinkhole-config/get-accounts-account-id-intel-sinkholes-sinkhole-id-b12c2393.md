---
title: Get a sinkhole
page_id: operation-get-accounts-account-id-intel-sinkholes-sinkhole-id-600a3cba
path: operations/sinkhole-config
description: Get the specified sinkhole by its unique identifier.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/sinkholes/{sinkhole_id}
operation_ids:
    - sinkhole-config-get-sinkhole
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a sinkhole

`GET /accounts/{account_id}/intel/sinkholes/{sinkhole_id}`

Operation ID: `sinkhole-config-get-sinkhole`

Get the specified sinkhole by its unique identifier.

## Definition

```yaml
{"operationId": "sinkhole-config-get-sinkhole", "summary": "Get a sinkhole", "description": "Get the specified sinkhole by its unique identifier.", "parameters": [{"$ref": "#/components/parameters/intel-sinkholes_account_id"}, {"$ref": "#/components/parameters/intel-sinkholes_sinkhole_id"}], "responses": {"200": {"description": "Successful Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_sinkhole_single_response"}}}}, "4XX": {"description": "Error Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_api_response_common_failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Sinkhole Config"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
