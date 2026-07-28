---
title: Update a sinkhole
page_id: operation-put-accounts-account-id-intel-sinkholes-sinkhole-id-f0bddfc4
path: operations/sinkhole-config
description: Update the name or R2 configuration of the specified sinkhole.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/intel/sinkholes/{sinkhole_id}
operation_ids:
    - sinkhole-config-update-sinkhole
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a sinkhole

`PUT /accounts/{account_id}/intel/sinkholes/{sinkhole_id}`

Operation ID: `sinkhole-config-update-sinkhole`

Update the name or R2 configuration of the specified sinkhole.

## Definition

```yaml
{"operationId": "sinkhole-config-update-sinkhole", "summary": "Update a sinkhole", "description": "Update the name or R2 configuration of the specified sinkhole.", "parameters": [{"$ref": "#/components/parameters/intel-sinkholes_account_id"}, {"$ref": "#/components/parameters/intel-sinkholes_sinkhole_id"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_sinkhole_create_params"}}}}, "responses": {"200": {"description": "Successful Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_api_response_single_empty"}}}}, "4XX": {"description": "Error Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_api_response_common_failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Sinkhole Config"], "x-api-token-group": ["Intel Write"]}
```
