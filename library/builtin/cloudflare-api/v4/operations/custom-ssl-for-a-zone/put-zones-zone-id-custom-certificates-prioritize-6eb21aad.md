---
title: Re-prioritize SSL Certificates
page_id: operation-put-zones-zone-id-custom-certificates-prioritize-52cac683
path: operations/custom-ssl-for-a-zone
description: If a zone has multiple SSL certificates, you can set the order in which they should be used during a request. The higher priority will break ties across overlapping 'legacy_custom' certificates.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/custom_certificates/prioritize
operation_ids:
    - custom-ssl-for-a-zone-re-prioritize-ssl-certificates
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Re-prioritize SSL Certificates

`PUT /zones/{zone_id}/custom_certificates/prioritize`

Operation ID: `custom-ssl-for-a-zone-re-prioritize-ssl-certificates`

If a zone has multiple SSL certificates, you can set the order in which they should be used during a request. The higher priority will break ties across overlapping 'legacy_custom' certificates.

## Definition

```yaml
{"operationId": "custom-ssl-for-a-zone-re-prioritize-ssl-certificates", "summary": "Re-prioritize SSL Certificates", "description": "If a zone has multiple SSL certificates, you can set the order in which they should be used during a request. The higher priority will break ties across overlapping 'legacy_custom' certificates.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"certificates": {"description": "Array of ordered certificates.", "type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}, "priority": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_priority"}}, "type": "object"}, "example": [{"id": "5a7805061c76ada191ed06f989cc3dac", "priority": 2}, {"id": "9a7806061c88ada191ed06f989cc3dac", "priority": 1}]}}, "required": ["certificates"]}}}}, "responses": {"200": {"description": "Re-prioritize SSL Certificates response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_collection"}}}}, "4XX": {"description": "Re-prioritize SSL Certificates response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_collection"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom SSL for a Zone"], "x-api-token-group": ["Access: Mutual TLS Certificates Write", "SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-certificates.prioritize", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
