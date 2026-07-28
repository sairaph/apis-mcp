---
title: List Regional Hostnames
page_id: operation-get-zones-zone-id-addressing-regional-hostnames-311d9108
path: operations/dls-regional-services
description: List all Regional Hostnames within a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/addressing/regional_hostnames
operation_ids:
    - dls-zone-regional-hostnames-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Regional Hostnames

`GET /zones/{zone_id}/addressing/regional_hostnames`

Operation ID: `dls-zone-regional-hostnames-list`

List all Regional Hostnames within a zone.

## Definition

```yaml
{"operationId": "dls-zone-regional-hostnames-list", "summary": "List Regional Hostnames", "description": "List all Regional Hostnames within a zone.", "parameters": [{"$ref": "#/components/parameters/dls_zone_id"}], "responses": {"200": {"description": "List hostnames response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dls_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/dls_regional_hostname_response"}}}}]}}}}, "4XX": {"description": "Failure to list hostnames", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dls_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["DLS Regional Services"], "x-api-token-group": ["DNS Read", "DNS Write"]}
```
