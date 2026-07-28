---
title: Get an ingress rule
page_id: operation-get-zones-zone-id-intel-sinkholes-sinkhole-id-ingresses-ingress-id-f6660cc2
path: operations/sinkhole-config
description: Get the specified ingress rule associated with a sinkhole. The sinkhole must belong to the same account as the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/intel/sinkholes/{sinkhole_id}/ingresses/{ingress_id}
operation_ids:
    - sinkhole-config-get-ingress
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an ingress rule

`GET /zones/{zone_id}/intel/sinkholes/{sinkhole_id}/ingresses/{ingress_id}`

Operation ID: `sinkhole-config-get-ingress`

Get the specified ingress rule associated with a sinkhole. The sinkhole must belong to the same account as the zone.

## Definition

```yaml
{"operationId": "sinkhole-config-get-ingress", "summary": "Get an ingress rule", "description": "Get the specified ingress rule associated with a sinkhole. The sinkhole must belong to the same account as the zone.", "parameters": [{"$ref": "#/components/parameters/intel-sinkholes_zone_id"}, {"$ref": "#/components/parameters/intel-sinkholes_sinkhole_id"}, {"$ref": "#/components/parameters/intel-sinkholes_ingress_id"}], "responses": {"200": {"description": "Successful Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_ingress_single_response"}}}}, "4XX": {"description": "Error Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_api_response_common_failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Sinkhole Config"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
