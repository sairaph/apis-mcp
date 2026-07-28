---
title: Billing History Details
page_id: operation-get-user-billing-history-e4a5f2ee
path: operations/user-billing-history
description: Accesses your billing history object.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/billing/history
operation_ids:
    - user-billing-history-(-deprecated)-billing-history-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Billing History Details

`GET /user/billing/history`

Operation ID: `user-billing-history-(-deprecated)-billing-history-details`

Accesses your billing history object.

## Definition

```yaml
{"operationId": "user-billing-history-(-deprecated)-billing-history-details", "summary": "Billing History Details", "description": "Accesses your billing history object.", "parameters": [{"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of items per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "order", "in": "query", "schema": {"description": "Field to order billing history by.", "type": "string", "example": "occurred_at", "enum": ["type", "occurred_at", "action"]}}, {"name": "occurred_at", "in": "query", "schema": {"$ref": "#/components/schemas/bill-subs-api_occurred_at"}}, {"name": "type", "in": "query", "schema": {"description": "The billing item type.", "type": "string", "example": "charge", "maxLength": 30, "readOnly": true}}, {"name": "action", "in": "query", "schema": {"description": "The billing item action.", "type": "string", "example": "subscription", "maxLength": 30, "readOnly": true}}], "responses": {"200": {"description": "Billing History Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_billing_history_collection"}}}}, "4XX": {"description": "Billing History Details response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_billing_history_collection"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["User Billing History"], "x-api-token-group": ["Billing Write", "Billing Read"], "x-cfPermissionsRequired": {"enum": ["#billing:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "user.billing.history", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
