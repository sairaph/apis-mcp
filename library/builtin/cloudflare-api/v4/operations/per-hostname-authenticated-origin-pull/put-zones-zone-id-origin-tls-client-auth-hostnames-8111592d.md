---
title: Enable or Disable a Hostname for Client Authentication
page_id: operation-put-zones-zone-id-origin-tls-client-auth-hostnames-d669c06c
path: operations/per-hostname-authenticated-origin-pull
description: 'Associate a hostname to a certificate and enable, disable or invalidate the association. If disabled, client certificate will not be sent to the hostname even if activated at the zone level. 100 maximum associations on a single certificate are allowed. Note: Use a null value for parameter *enabled* to invalidate the association.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/origin_tls_client_auth/hostnames
operation_ids:
    - per-hostname-authenticated-origin-pull-enable-or-disable-a-hostname-for-client-authentication
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Enable or Disable a Hostname for Client Authentication

`PUT /zones/{zone_id}/origin_tls_client_auth/hostnames`

Operation ID: `per-hostname-authenticated-origin-pull-enable-or-disable-a-hostname-for-client-authentication`

Associate a hostname to a certificate and enable, disable or invalidate the association. If disabled, client certificate will not be sent to the hostname even if activated at the zone level. 100 maximum associations on a single certificate are allowed. Note: Use a null value for parameter *enabled* to invalidate the association.

## Definition

```yaml
{"operationId": "per-hostname-authenticated-origin-pull-enable-or-disable-a-hostname-for-client-authentication", "summary": "Enable or Disable a Hostname for Client Authentication", "description": "Associate a hostname to a certificate and enable, disable or invalidate the association. If disabled, client certificate will not be sent to the hostname even if activated at the zone level. 100 maximum associations on a single certificate are allowed. Note: Use a null value for parameter *enabled* to invalidate the association.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"config": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_config"}}, "required": ["config"]}}}}, "responses": {"200": {"description": "Enable or Disable a Hostname for Client Authentication response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname_aop_response_collection"}}}}, "4XX": {"description": "Enable or Disable a Hostname for Client Authentication response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostname_aop_response_collection"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Per-hostname Authenticated Origin Pull"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-tls-client-auth.hostnames", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
