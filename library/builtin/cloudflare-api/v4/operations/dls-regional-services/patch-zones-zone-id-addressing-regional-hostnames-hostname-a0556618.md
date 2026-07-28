---
title: Update Regional Hostname
page_id: operation-patch-zones-zone-id-addressing-regional-hostnames-hostname-0a1bbc9a
path: operations/dls-regional-services
description: Update the configuration for a specific Regional Hostname. Only the region_key of a hostname is mutable.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/addressing/regional_hostnames/{hostname}
operation_ids:
    - dls-zone-regional-hostnames-patch
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Regional Hostname

`PATCH /zones/{zone_id}/addressing/regional_hostnames/{hostname}`

Operation ID: `dls-zone-regional-hostnames-patch`

Update the configuration for a specific Regional Hostname. Only the region_key of a hostname is mutable.

## Definition

```yaml
{"operationId": "dls-zone-regional-hostnames-patch", "summary": "Update Regional Hostname", "description": "Update the configuration for a specific Regional Hostname. Only the region_key of a hostname is mutable.", "parameters": [{"$ref": "#/components/parameters/dls_zone_id"}, {"$ref": "#/components/parameters/dls_hostname"}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"region_key": {"$ref": "#/components/schemas/dls_region_key"}}, "required": ["region_key"]}}}}, "responses": {"200": {"description": "Update hostname response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dls_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/dls_regional_hostname_response"}}, "type": "object"}]}}}}, "4XX": {"description": "Failure to update hostname", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dls_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["DLS Regional Services"], "x-api-token-group": ["DNS Write"]}
```
