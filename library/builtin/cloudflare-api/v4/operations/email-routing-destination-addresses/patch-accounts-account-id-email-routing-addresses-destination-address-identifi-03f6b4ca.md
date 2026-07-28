---
title: Update destination address
page_id: operation-patch-accounts-account-id-email-routing-addresses-destination-address-id-2036314f
path: operations/email-routing-destination-addresses
description: Updates the status of a specific destination address.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}
operation_ids:
    - email-routing-destination-addresses-update-destination-address
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update destination address

`PATCH /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}`

Operation ID: `email-routing-destination-addresses-update-destination-address`

Updates the status of a specific destination address.

## Definition

```yaml
{"operationId": "email-routing-destination-addresses-update-destination-address", "summary": "Update destination address", "description": "Updates the status of a specific destination address.", "parameters": [{"name": "destination_address_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_destination_address_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_update_destination_address_properties"}}}}, "responses": {"200": {"description": "Update destination address response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_destination_address_response_single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing destination addresses"], "x-api-token-group": ["Email Routing Addresses Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.routing.address.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.addresses", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
