---
title: Retrieves countries information for all countries
page_id: operation-get-accounts-account-id-cloudforce-one-events-countries-24661795
path: operations/country
description: Retrieve country code information for all supported countries.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/countries
operation_ids:
    - get_CountryRead
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieves countries information for all countries

`GET /accounts/{account_id}/cloudforce-one/events/countries`

Operation ID: `get_CountryRead`

Retrieve country code information for all supported countries.

## Definition

```yaml
{"operationId": "get_CountryRead", "summary": "Retrieves countries information for all countries", "description": "Retrieve country code information for all supported countries.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "responses": {"200": {"description": "Returns the long and short country code for every country.", "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"result": {"type": "array", "items": {"properties": {"alpha2": {"type": "string", "example": "AF"}, "alpha3": {"type": "string", "example": "AF"}, "name": {"type": "string", "example": "Afghanistan"}}, "required": ["name", "alpha3", "alpha2"], "type": "object"}}, "success": {"type": "string", "example": "true"}}, "required": ["success", "result"], "type": "object"}}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Country"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
