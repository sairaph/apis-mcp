---
title: Delete Zone Subscription
page_id: operation-delete-zones-zone-id-subscription-00c25f95
path: operations/zone-subscription
description: Deletes a zone's subscription.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/subscription
operation_ids:
    - zone-subscription-delete-zone-subscription
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Zone Subscription

`DELETE /zones/{zone_id}/subscription`

Operation ID: `zone-subscription-delete-zone-subscription`

Deletes a zone's subscription.

## Definition

```yaml
{"operationId": "zone-subscription-delete-zone-subscription", "summary": "Delete Zone Subscription", "description": "Deletes a zone's subscription.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_identifier"}}], "responses": {"200": {"description": "Delete Zone Subscription response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"subscription_id": {"$ref": "#/components/schemas/bill-subs-api_schemas-identifier"}}}}}]}}}}, "4XX": {"description": "Delete Zone Subscription response failure", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/bill-subs-api_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"subscription_id": {"$ref": "#/components/schemas/bill-subs-api_schemas-identifier"}}}}}]}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zone Subscription"], "x-api-token-group": ["Billing Write"], "x-cfPermissionsRequired": {"enum": ["#billing:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
