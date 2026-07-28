---
title: List ingresses for a sinkhole
page_id: operation-get-accounts-account-id-intel-sinkholes-sinkhole-id-ingresses-457b3d2f
path: operations/sinkhole-config
description: List all ingress rules associated with the specified sinkhole.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/sinkholes/{sinkhole_id}/ingresses
operation_ids:
    - sinkhole-config-list-sinkhole-ingresses
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List ingresses for a sinkhole

`GET /accounts/{account_id}/intel/sinkholes/{sinkhole_id}/ingresses`

Operation ID: `sinkhole-config-list-sinkhole-ingresses`

List all ingress rules associated with the specified sinkhole.

## Definition

```yaml
{"operationId": "sinkhole-config-list-sinkhole-ingresses", "summary": "List ingresses for a sinkhole", "description": "List all ingress rules associated with the specified sinkhole.", "parameters": [{"$ref": "#/components/parameters/intel-sinkholes_account_id"}, {"$ref": "#/components/parameters/intel-sinkholes_sinkhole_id"}], "responses": {"200": {"description": "Successful Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_ingress_collection_response"}}}}, "4XX": {"description": "Error Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_api_response_common_failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Sinkhole Config"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
