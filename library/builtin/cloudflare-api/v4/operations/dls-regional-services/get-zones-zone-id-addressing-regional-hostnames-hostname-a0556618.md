---
title: Fetch Regional Hostname
page_id: operation-get-zones-zone-id-addressing-regional-hostnames-hostname-9f76655b
path: operations/dls-regional-services
description: Fetch the configuration for a specific Regional Hostname, within a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/addressing/regional_hostnames/{hostname}
operation_ids:
    - dls-zone-regional-hostnames-fetch
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch Regional Hostname

`GET /zones/{zone_id}/addressing/regional_hostnames/{hostname}`

Operation ID: `dls-zone-regional-hostnames-fetch`

Fetch the configuration for a specific Regional Hostname, within a zone.

## Definition

```yaml
{"operationId": "dls-zone-regional-hostnames-fetch", "summary": "Fetch Regional Hostname", "description": "Fetch the configuration for a specific Regional Hostname, within a zone.", "parameters": [{"$ref": "#/components/parameters/dls_zone_id"}, {"$ref": "#/components/parameters/dls_hostname"}], "responses": {"200": {"description": "Fetch hostname response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dls_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/dls_regional_hostname_response"}}, "type": "object"}]}}}}, "4XX": {"description": "Failure to fetch hostname", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dls_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["DLS Regional Services"], "x-api-token-group": ["DNS Read", "DNS Write"]}
```
