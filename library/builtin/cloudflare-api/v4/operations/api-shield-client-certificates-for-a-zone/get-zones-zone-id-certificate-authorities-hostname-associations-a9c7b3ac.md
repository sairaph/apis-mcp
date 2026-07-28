---
title: List Hostname Associations
page_id: operation-get-zones-zone-id-certificate-authorities-hostname-associations-6b4c2491
path: operations/api-shield-client-certificates-for-a-zone
description: List Hostname Associations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/certificate_authorities/hostname_associations
operation_ids:
    - client-certificate-for-a-zone-list-hostname-associations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Hostname Associations

`GET /zones/{zone_id}/certificate_authorities/hostname_associations`

Operation ID: `client-certificate-for-a-zone-list-hostname-associations`

List Hostname Associations.

## Definition

```yaml
{"operationId": "client-certificate-for-a-zone-list-hostname-associations", "summary": "List Hostname Associations", "description": "List Hostname Associations.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "mtls_certificate_id", "in": "query", "schema": {"description": "The UUID to match against for a certificate that was uploaded to the mTLS Certificate Management endpoint. If no mtls_certificate_id is given, the results will be the hostnames associated to your active Cloudflare Managed CA.", "type": "string", "example": "b2134436-2555-4acf-be5b-26c48136575e", "maxLength": 36, "minLength": 36}}], "responses": {"200": {"description": "List Hostname Associations Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname_associations_response"}}}}, "4XX": {"description": "List Hostname Associations Response Failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["API Shield Client Certificates for a Zone"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "certificate-authorities.hostname-associations", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
