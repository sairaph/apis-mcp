---
title: List destination addresses
page_id: operation-get-accounts-account-id-email-routing-addresses-e1a20eff
path: operations/email-routing-destination-addresses
description: Lists existing destination addresses.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email/routing/addresses
operation_ids:
    - email-routing-destination-addresses-list-destination-addresses
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List destination addresses

`GET /accounts/{account_id}/email/routing/addresses`

Operation ID: `email-routing-destination-addresses-list-destination-addresses`

Lists existing destination addresses.

## Definition

```yaml
{"operationId": "email-routing-destination-addresses-list-destination-addresses", "summary": "List destination addresses", "description": "Lists existing destination addresses.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of results per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "direction", "in": "query", "schema": {"description": "Sorts results in an ascending or descending order.", "type": "string", "example": "asc", "default": "asc", "enum": ["asc", "desc"]}}, {"name": "verified", "in": "query", "schema": {"description": "Filter by verified destination addresses.", "type": "boolean", "example": true, "default": true, "enum": [true, false]}}], "responses": {"200": {"description": "List destination addresses response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email_destination_addresses_response_collection"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Email Routing destination addresses"], "x-api-token-group": ["Email Routing Addresses Write", "Email Routing Addresses Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.routing.address.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-routing.addresses", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
