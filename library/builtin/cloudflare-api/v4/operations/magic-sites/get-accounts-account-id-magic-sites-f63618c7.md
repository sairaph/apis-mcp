---
title: List Sites
page_id: operation-get-accounts-account-id-magic-sites-68f9bbd7
path: operations/magic-sites
description: Lists Sites associated with an account. Use connectorid query param to return sites where connectorid matches either site.ConnectorID or site.SecondaryConnectorID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/sites
operation_ids:
    - magic-sites-list-sites
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Sites

`GET /accounts/{account_id}/magic/sites`

Operation ID: `magic-sites-list-sites`

Lists Sites associated with an account. Use connectorid query param to return sites where connectorid matches either site.ConnectorID or site.SecondaryConnectorID.

## Definition

```yaml
{"operationId": "magic-sites-list-sites", "summary": "List Sites", "description": "Lists Sites associated with an account. Use connectorid query param to return sites where connectorid matches either site.ConnectorID or site.SecondaryConnectorID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "connectorid", "in": "query", "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "List Sites response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_sites_collection_response"}}}}, "4XX": {"description": "List Sites response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Sites"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
