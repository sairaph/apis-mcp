---
title: Retrieve the DCV Delegation unique identifier.
page_id: operation-get-zones-zone-id-dcv-delegation-uuid-fbc0a7fc
path: operations/dcv-delegation
description: Retrieve the account and zone specific unique identifier used as part of the CNAME target for DCV Delegation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/dcv_delegation/uuid
operation_ids:
    - dcv-delegation-uuid-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve the DCV Delegation unique identifier.

`GET /zones/{zone_id}/dcv_delegation/uuid`

Operation ID: `dcv-delegation-uuid-get`

Retrieve the account and zone specific unique identifier used as part of the CNAME target for DCV Delegation.

## Definition

```yaml
{"operationId": "dcv-delegation-uuid-get", "summary": "Retrieve the DCV Delegation unique identifier.", "description": "Retrieve the account and zone specific unique identifier used as part of the CNAME target for DCV Delegation.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Retrieve the DCV Delegation unique identifier response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_dcv_delegation_response"}}}}, "4XX": {"description": "Retrieve the DCV Delegation unique identifier response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_dcv_delegation_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DCV Delegation"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dcv-delegation", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
