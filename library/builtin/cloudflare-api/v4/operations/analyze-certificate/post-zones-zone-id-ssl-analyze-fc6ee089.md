---
title: Analyze Certificate
page_id: operation-post-zones-zone-id-ssl-analyze-98c29807
path: operations/analyze-certificate
description: Returns the set of hostnames, the signature algorithm, and the expiration date of the certificate.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/ssl/analyze
operation_ids:
    - analyze-certificate-analyze-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Analyze Certificate

`POST /zones/{zone_id}/ssl/analyze`

Operation ID: `analyze-certificate-analyze-certificate`

Returns the set of hostnames, the signature algorithm, and the expiration date of the certificate.

## Definition

```yaml
{"operationId": "analyze-certificate-analyze-certificate", "summary": "Analyze Certificate", "description": "Returns the set of hostnames, the signature algorithm, and the expiration date of the certificate.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"bundle_method": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_bundle_method"}, "certificate": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate"}}}}}}, "responses": {"200": {"description": "Analyze Certificate response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_analyze_response"}}}}, "4XX": {"description": "Analyze Certificate response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_analyze_response"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Analyze Certificate"], "x-api-token-group": ["Access: Mutual TLS Certificates Write", "Access: Mutual TLS Certificates Read", "SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ssl.analyze", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
