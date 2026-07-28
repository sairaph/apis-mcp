---
title: Delete a sinkhole
page_id: operation-delete-accounts-account-id-intel-sinkholes-sinkhole-id-4c594531
path: operations/sinkhole-config
description: Delete the specified sinkhole. The sinkhole must not have any active ingress rules defined. A 409 response code indicates that this condition is not met.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/intel/sinkholes/{sinkhole_id}
operation_ids:
    - sinkhole-config-delete-sinkhole
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a sinkhole

`DELETE /accounts/{account_id}/intel/sinkholes/{sinkhole_id}`

Operation ID: `sinkhole-config-delete-sinkhole`

Delete the specified sinkhole. The sinkhole must not have any active ingress rules defined. A 409 response code indicates that this condition is not met.

## Definition

```yaml
{"operationId": "sinkhole-config-delete-sinkhole", "summary": "Delete a sinkhole", "description": "Delete the specified sinkhole. The sinkhole must not have any active ingress rules defined. A 409 response code indicates that this condition is not met.", "parameters": [{"$ref": "#/components/parameters/intel-sinkholes_account_id"}, {"$ref": "#/components/parameters/intel-sinkholes_sinkhole_id"}], "responses": {"200": {"description": "Successful Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_api_response_single_empty"}}}}, "409": {"description": "Conflict -- the sinkhole has active ingress rules that must be deleted first.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_api_response_common_failure"}}}}, "4XX": {"description": "Error Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_api_response_common_failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Sinkhole Config"], "x-api-token-group": ["Intel Write"]}
```
