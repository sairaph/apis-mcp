---
title: Create Regional Hostname
page_id: operation-post-zones-zone-id-addressing-regional-hostnames-63130b80
path: operations/dls-regional-services
description: Create a new Regional Hostname entry. Cloudflare will only use data centers that are physically located within the chosen region to decrypt and service HTTPS traffic. Learn more about [Regional Services](https://developers.cloudflare.com/data-localization/regional-services/get-started/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/addressing/regional_hostnames
operation_ids:
    - dls-zone-regional-hostnames-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Regional Hostname

`POST /zones/{zone_id}/addressing/regional_hostnames`

Operation ID: `dls-zone-regional-hostnames-create`

Create a new Regional Hostname entry. Cloudflare will only use data centers that are physically located within the chosen region to decrypt and service HTTPS traffic. Learn more about [Regional Services](https://developers.cloudflare.com/data-localization/regional-services/get-started/).

## Definition

```yaml
{"operationId": "dls-zone-regional-hostnames-create", "summary": "Create Regional Hostname", "description": "Create a new Regional Hostname entry. Cloudflare will only use data centers that are physically located within the chosen region to decrypt and service HTTPS traffic. Learn more about [Regional Services](https://developers.cloudflare.com/data-localization/regional-services/get-started/).", "parameters": [{"$ref": "#/components/parameters/dls_zone_id"}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"hostname": {"$ref": "#/components/schemas/dls_hostname"}, "region_key": {"$ref": "#/components/schemas/dls_region_key"}, "routing": {"$ref": "#/components/schemas/dls_routing"}}, "required": ["hostname", "region_key"]}}}}, "responses": {"200": {"description": "Create hostname response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dls_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/dls_regional_hostname_response"}}, "type": "object"}]}}}}, "4XX": {"description": "Failure to create hostname", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dls_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["DLS Regional Services"], "x-api-token-group": ["DNS Write"]}
```
