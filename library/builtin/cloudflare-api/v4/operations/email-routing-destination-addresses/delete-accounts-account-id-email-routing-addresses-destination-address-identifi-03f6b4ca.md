---
title: Delete destination address
page_id: operation-delete-accounts-account-id-email-routing-addresses-destination-address-i-dcbf5866
path: operations/email-routing-destination-addresses
description: Deletes a specific destination address.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}
operation_ids:
    - email-routing-destination-addresses-delete-destination-address
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete destination address

`DELETE /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}`

Operation ID: `email-routing-destination-addresses-delete-destination-address`

Deletes a specific destination address.

## Definition

```yaml
{"operationId": "email-routing-destination-addresses-delete-destination-address", "summary": "Delete destination address", "description": "Deletes a specific destination address.", "parameters": [{"name": "destination_address_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_destination_address_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "responses": {"200": {"description": "Delete destination address response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_destination_address_response_single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing destination addresses"], "x-api-token-group": ["Email Routing Addresses Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.routing.address.delete"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.addresses", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
