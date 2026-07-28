---
title: Billing Profile Details
page_id: operation-get-user-billing-profile-834ca15d
path: operations/user-billing-profile
description: Accesses your billing profile object.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/billing/profile
operation_ids:
    - user-billing-profile-(-deprecated)-billing-profile-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Billing Profile Details

`GET /user/billing/profile`

Operation ID: `user-billing-profile-(-deprecated)-billing-profile-details`

Accesses your billing profile object.

## Definition

```yaml
{"operationId": "user-billing-profile-(-deprecated)-billing-profile-details", "summary": "Billing Profile Details", "description": "Accesses your billing profile object.", "responses": {"200": {"description": "Billing Profile Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_billing_response_single"}}}}, "4XX": {"description": "Billing Profile Details response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_billing_response_single"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["User Billing Profile"], "x-api-token-group": ["Billing Write", "Billing Read"], "x-cfPermissionsRequired": {"enum": ["#billing:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "user.billing.profile", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
