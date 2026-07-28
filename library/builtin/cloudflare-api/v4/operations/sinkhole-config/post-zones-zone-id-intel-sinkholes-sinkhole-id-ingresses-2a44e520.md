---
title: Create an ingress rule
page_id: operation-post-zones-zone-id-intel-sinkholes-sinkhole-id-ingresses-0beee51b
path: operations/sinkhole-config
description: Create a new ingress rule for the specified sinkhole. The CIDR block must be a Cloudflare BYOIP associated with your account. The zone_id must be a zone with the ability to create Spectrum Apps. The sinkhole must belong to the same account as the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/intel/sinkholes/{sinkhole_id}/ingresses
operation_ids:
    - sinkhole-config-create-ingress
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an ingress rule

`POST /zones/{zone_id}/intel/sinkholes/{sinkhole_id}/ingresses`

Operation ID: `sinkhole-config-create-ingress`

Create a new ingress rule for the specified sinkhole. The CIDR block must be a Cloudflare BYOIP associated with your account. The zone_id must be a zone with the ability to create Spectrum Apps. The sinkhole must belong to the same account as the zone.

## Definition

```yaml
{"operationId": "sinkhole-config-create-ingress", "summary": "Create an ingress rule", "description": "Create a new ingress rule for the specified sinkhole. The CIDR block must be a Cloudflare BYOIP associated with your account. The zone_id must be a zone with the ability to create Spectrum Apps. The sinkhole must belong to the same account as the zone.", "parameters": [{"$ref": "#/components/parameters/intel-sinkholes_zone_id"}, {"$ref": "#/components/parameters/intel-sinkholes_sinkhole_id"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_ingress_create_params"}}}}, "responses": {"201": {"description": "Ingress Created", "headers": {"Location": {"description": "URI of the created ingress", "schema": {"type": "string"}}}, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_ingress_single_response"}}}}, "4XX": {"description": "Error Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel-sinkholes_api_response_common_failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Sinkhole Config"], "x-api-token-group": ["Intel Write"]}
```
