---
title: Get a destination address
page_id: operation-get-accounts-account-id-email-routing-addresses-destination-address-iden-a64734a4
path: operations/email-routing-destination-addresses
description: Gets information for a specific destination email already created.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}
operation_ids:
    - email-routing-destination-addresses-get-a-destination-address
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a destination address

`GET /accounts/{account_id}/email/routing/addresses/{destination_address_identifier}`

Operation ID: `email-routing-destination-addresses-get-a-destination-address`

Gets information for a specific destination email already created.

## Definition

```yaml
{"operationId": "email-routing-destination-addresses-get-a-destination-address", "summary": "Get a destination address", "description": "Gets information for a specific destination email already created.", "parameters": [{"name": "destination_address_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_destination_address_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}], "responses": {"200": {"description": "Get a destination address response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_destination_address_response_single"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing destination addresses"], "x-api-token-group": ["Email Routing Addresses Write", "Email Routing Addresses Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.routing.address.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.addresses", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
