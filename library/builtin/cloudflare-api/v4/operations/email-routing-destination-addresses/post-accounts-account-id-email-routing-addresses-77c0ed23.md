---
title: Create a destination address
page_id: operation-post-accounts-account-id-email-routing-addresses-be8b02d0
path: operations/email-routing-destination-addresses
description: Create a destination address to forward your emails to. Destination addresses need to be verified before they can be used.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email/routing/addresses
operation_ids:
    - email-routing-destination-addresses-create-a-destination-address
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a destination address

`POST /accounts/{account_id}/email/routing/addresses`

Operation ID: `email-routing-destination-addresses-create-a-destination-address`

Create a destination address to forward your emails to. Destination addresses need to be verified before they can be used.

## Definition

```yaml
{"operationId": "email-routing-destination-addresses-create-a-destination-address", "summary": "Create a destination address", "description": "Create a destination address to forward your emails to. Destination addresses need to be verified before they can be used.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_create_destination_address_properties"}}}}, "responses": {"200": {"description": "Create a destination address response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_destination_address_response_single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing destination addresses"], "x-api-token-group": ["Email Routing Addresses Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.routing.address.create"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.addresses", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
