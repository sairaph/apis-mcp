---
title: Update CT Alerting Subscription
page_id: operation-patch-zones-zone-id-ct-alerting-969f442a
path: operations/ct-alerting
description: |-
    Create or update the Certificate Transparency alerting subscription for a zone. Enables or disables email notifications when certificates are issued for the zone's domains.
    For Free and Pro zones, the subscription is toggled on or off using the enabled field. Notification emails are sent to all users with SSL permissions on the zone.
    For Business and Enterprise zones, the emails field is required and controls which addresses receive alerts. Setting emails to an empty list disables the subscription regardless of the enabled field. A maximum of 10 email addresses may be configured.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/ct/alerting
operation_ids:
    - ct-alerting-update-subscription
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update CT Alerting Subscription

`PATCH /zones/{zone_id}/ct/alerting`

Operation ID: `ct-alerting-update-subscription`

Create or update the Certificate Transparency alerting subscription for a zone. Enables or disables email notifications when certificates are issued for the zone's domains.
For Free and Pro zones, the subscription is toggled on or off using the enabled field. Notification emails are sent to all users with SSL permissions on the zone.
For Business and Enterprise zones, the emails field is required and controls which addresses receive alerts. Setting emails to an empty list disables the subscription regardless of the enabled field. A maximum of 10 email addresses may be configured.

## Definition

```yaml
{"operationId": "ct-alerting-update-subscription", "summary": "Update CT Alerting Subscription", "description": "Create or update the Certificate Transparency alerting subscription for a zone. Enables or disables email notifications when certificates are issued for the zone's domains.\nFor Free and Pro zones, the subscription is toggled on or off using the enabled field. Notification emails are sent to all users with SSL permissions on the zone.\nFor Business and Enterprise zones, the emails field is required and controls which addresses receive alerts. Setting emails to an empty list disables the subscription regardless of the enabled field. A maximum of 10 email addresses may be configured.\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_ct_alerting_subscription_update"}}}}, "responses": {"200": {"description": "Update CT Alerting Subscription response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_ct_alerting_subscription_response_single"}}}}, "4XX": {"description": "Update CT Alerting Subscription response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_ct_alerting_subscription_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["CT Alerting"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
