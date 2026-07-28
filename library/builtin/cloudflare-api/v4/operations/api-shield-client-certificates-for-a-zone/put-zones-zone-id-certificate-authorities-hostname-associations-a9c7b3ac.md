---
title: Replace Hostname Associations
page_id: operation-put-zones-zone-id-certificate-authorities-hostname-associations-711a9c3c
path: operations/api-shield-client-certificates-for-a-zone
description: Replace Hostname Associations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/certificate_authorities/hostname_associations
operation_ids:
    - client-certificate-for-a-zone-put-hostname-associations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Replace Hostname Associations

`PUT /zones/{zone_id}/certificate_authorities/hostname_associations`

Operation ID: `client-certificate-for-a-zone-put-hostname-associations`

Replace Hostname Associations.

## Definition

```yaml
{"operationId": "client-certificate-for-a-zone-put-hostname-associations", "summary": "Replace Hostname Associations", "description": "Replace Hostname Associations.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname_association"}}}}, "responses": {"200": {"description": "Replace Hostname Associations Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname_associations_response"}}}}, "4XX": {"description": "Replace Hostname Associations Response Failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["API Shield Client Certificates for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "certificate-authorities.hostname-associations", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
