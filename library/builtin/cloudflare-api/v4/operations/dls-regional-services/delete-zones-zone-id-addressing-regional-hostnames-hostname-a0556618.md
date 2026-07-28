---
title: Delete Regional Hostname
page_id: operation-delete-zones-zone-id-addressing-regional-hostnames-hostname-70cd091e
path: operations/dls-regional-services
description: Delete the region configuration for a specific Regional Hostname.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/addressing/regional_hostnames/{hostname}
operation_ids:
    - dls-zone-regional-hostnames-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Regional Hostname

`DELETE /zones/{zone_id}/addressing/regional_hostnames/{hostname}`

Operation ID: `dls-zone-regional-hostnames-delete`

Delete the region configuration for a specific Regional Hostname.

## Definition

```yaml
{"operationId": "dls-zone-regional-hostnames-delete", "summary": "Delete Regional Hostname", "description": "Delete the region configuration for a specific Regional Hostname.", "parameters": [{"$ref": "#/components/parameters/dls_zone_id"}, {"$ref": "#/components/parameters/dls_hostname"}], "responses": {"200": {"description": "Delete hostname response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dls_api-response-common"}]}}}}, "4XX": {"description": "Failure to delete hostname", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dls_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["DLS Regional Services"], "x-api-token-group": ["DNS Write"]}
```
