---
title: List SSL Configurations
page_id: operation-get-zones-zone-id-custom-certificates-c52feee0
path: operations/custom-ssl-for-a-zone
description: List, search, and filter all of your custom SSL certificates. The higher priority will break ties across overlapping 'legacy_custom' certificates, but 'legacy_custom' certificates will always supercede 'sni_custom' certificates.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/custom_certificates
operation_ids:
    - custom-ssl-for-a-zone-list-ssl-configurations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List SSL Configurations

`GET /zones/{zone_id}/custom_certificates`

Operation ID: `custom-ssl-for-a-zone-list-ssl-configurations`

List, search, and filter all of your custom SSL certificates. The higher priority will break ties across overlapping 'legacy_custom' certificates, but 'legacy_custom' certificates will always supercede 'sni_custom' certificates.

## Definition

```yaml
{"operationId": "custom-ssl-for-a-zone-list-ssl-configurations", "summary": "List SSL Configurations", "description": "List, search, and filter all of your custom SSL certificates. The higher priority will break ties across overlapping 'legacy_custom' certificates, but 'legacy_custom' certificates will always supercede 'sni_custom' certificates.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of zones per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "match", "in": "query", "schema": {"description": "Whether to match all search requirements or at least one (any).", "type": "string", "default": "all", "enum": ["any", "all"]}}, {"name": "status", "in": "query", "schema": {"description": "Status of the zone's custom SSL.", "type": "string", "example": "active", "enum": ["active", "expired", "deleted", "pending", "initializing"], "readOnly": true}}], "responses": {"200": {"description": "List SSL Configurations response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_collection"}}}}, "4XX": {"description": "List SSL Configurations response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_collection"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom SSL for a Zone"], "x-api-token-group": ["Access: Mutual TLS Certificates Write", "Access: Mutual TLS Certificates Read", "SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-certificates", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
