---
title: Get Custom Hostname Quota
page_id: operation-get-zones-zone-id-custom-hostnames-quota-466e2ab4
path: operations/custom-hostname-for-a-zone
description: Returns custom hostname quota usage for a zone. The allocated quota is a soft limit; creating custom hostnames after usage exceeds this limit can still succeed until the hard cap is reached. Use the exceeded and hard_cap fields to track when usage is above the soft limit and when new custom hostname creation will be rejected.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/custom_hostnames/quota
operation_ids:
    - custom-hostname-for-a-zone-get-custom-hostname-quota
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Custom Hostname Quota

`GET /zones/{zone_id}/custom_hostnames/quota`

Operation ID: `custom-hostname-for-a-zone-get-custom-hostname-quota`

Returns custom hostname quota usage for a zone. The allocated quota is a soft limit; creating custom hostnames after usage exceeds this limit can still succeed until the hard cap is reached. Use the exceeded and hard_cap fields to track when usage is above the soft limit and when new custom hostname creation will be rejected.

## Definition

```yaml
{"operationId": "custom-hostname-for-a-zone-get-custom-hostname-quota", "summary": "Get Custom Hostname Quota", "description": "Returns custom hostname quota usage for a zone. The allocated quota is a soft limit; creating custom hostnames after usage exceeds this limit can still succeed until the hard cap is reached. Use the exceeded and hard_cap fields to track when usage is above the soft limit and when new custom hostname creation will be rejected.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Get Custom Hostname Quota response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_hostname_quota_response"}}}}, "4XX": {"description": "Get Custom Hostname Quota response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_hostname_api_response_failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom Hostname for a Zone"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
