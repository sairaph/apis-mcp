---
title: Get CT Alerting Subscription
page_id: operation-get-zones-zone-id-ct-alerting-0410e380
path: operations/ct-alerting
description: Retrieve the Certificate Transparency alerting subscription settings for a zone. Returns whether CT monitoring is enabled and, for Business and Enterprise zones, the list of email addresses that receive alerts.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/ct/alerting
operation_ids:
    - ct-alerting-get-subscription
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get CT Alerting Subscription

`GET /zones/{zone_id}/ct/alerting`

Operation ID: `ct-alerting-get-subscription`

Retrieve the Certificate Transparency alerting subscription settings for a zone. Returns whether CT monitoring is enabled and, for Business and Enterprise zones, the list of email addresses that receive alerts.

## Definition

```yaml
{"operationId": "ct-alerting-get-subscription", "summary": "Get CT Alerting Subscription", "description": "Retrieve the Certificate Transparency alerting subscription settings for a zone. Returns whether CT monitoring is enabled and, for Business and Enterprise zones, the list of email addresses that receive alerts.\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Get CT Alerting Subscription response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_ct_alerting_subscription_response_single"}}}}, "4XX": {"description": "Get CT Alerting Subscription response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_ct_alerting_subscription_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["CT Alerting"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
