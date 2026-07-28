---
title: Search for available domains
page_id: operation-get-accounts-account-id-registrar-domain-search-7e737f0c
path: operations/domain-discovery
description: |-
    Searches for domain name suggestions based on a keyword, phrase, or partial domain name.
    Returns a list of potentially available domains with pricing information.

    **Important:** Results are non-authoritative and based on cached data. Always use the
    `/domain-check` endpoint to verify real-time availability before attempting registration.

    Suggestions are scoped to extensions supported for programmatic registration
    via this API (`POST /registrations`). Domains on unsupported extensions will
    not appear in results, even if they are available at the registry level.

    ### Use cases
    - Brand name discovery (e.g., "acme corp" → acmecorp.com, acmecorp.dev)
    - Keyword-based suggestions (e.g., "coffee shop" → coffeeshop.com, mycoffeeshop.net)
    - Alternative extension discovery (e.g., "example.com" → example.com, example.app, example.xyz)

    ### Workflow
    1. Call this endpoint with a keyword or domain name.
    2. Present suggestions to the user.
    3. Call `/domain-check` with the user's chosen domains to confirm real-time availability and pricing.
    4. Proceed to `POST /registrations` only for supported non-premium domains
       where the Check response returns `registrable: true`.

    **Note:** Searching with just a domain extension (e.g., "com" or ".app") is not supported. Provide a keyword or domain name.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/registrar/domain-search
operation_ids:
    - registrar-domain-discovery-search
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Search for available domains

`GET /accounts/{account_id}/registrar/domain-search`

Operation ID: `registrar-domain-discovery-search`

Searches for domain name suggestions based on a keyword, phrase, or partial domain name.
Returns a list of potentially available domains with pricing information.

**Important:** Results are non-authoritative and based on cached data. Always use the
`/domain-check` endpoint to verify real-time availability before attempting registration.

Suggestions are scoped to extensions supported for programmatic registration
via this API (`POST /registrations`). Domains on unsupported extensions will
not appear in results, even if they are available at the registry level.

### Use cases
- Brand name discovery (e.g., "acme corp" → acmecorp.com, acmecorp.dev)
- Keyword-based suggestions (e.g., "coffee shop" → coffeeshop.com, mycoffeeshop.net)
- Alternative extension discovery (e.g., "example.com" → example.com, example.app, example.xyz)

### Workflow
1. Call this endpoint with a keyword or domain name.
2. Present suggestions to the user.
3. Call `/domain-check` with the user's chosen domains to confirm real-time availability and pricing.
4. Proceed to `POST /registrations` only for supported non-premium domains
   where the Check response returns `registrable: true`.

**Note:** Searching with just a domain extension (e.g., "com" or ".app") is not supported. Provide a keyword or domain name.

## Definition

```yaml
{"operationId": "registrar-domain-discovery-search", "summary": "Search for available domains", "description": "Searches for domain name suggestions based on a keyword, phrase, or partial domain name.\nReturns a list of potentially available domains with pricing information.\n\n**Important:** Results are non-authoritative and based on cached data. Always use the\n`/domain-check` endpoint to verify real-time availability before attempting registration.\n\nSuggestions are scoped to extensions supported for programmatic registration\nvia this API (`POST /registrations`). Domains on unsupported extensions will\nnot appear in results, even if they are available at the registry level.\n\n### Use cases\n- Brand name discovery (e.g., \"acme corp\" → acmecorp.com, acmecorp.dev)\n- Keyword-based suggestions (e.g., \"coffee shop\" → coffeeshop.com, mycoffeeshop.net)\n- Alternative extension discovery (e.g., \"example.com\" → example.com, example.app, example.xyz)\n\n### Workflow\n1. Call this endpoint with a keyword or domain name.\n2. Present suggestions to the user.\n3. Call `/domain-check` with the user's chosen domains to confirm real-time availability and pricing.\n4. Proceed to `POST /registrations` only for supported non-premium domains\n   where the Check response returns `registrable: true`.\n\n**Note:** Searching with just a domain extension (e.g., \"com\" or \".app\") is not supported. Provide a keyword or domain name.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID. Required for all Registrar API operations.", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api_identifier"}}, {"name": "q", "in": "query", "description": "The search term to find domain suggestions. Accepts keywords, phrases, or full domain names.\n- Phrases: \"coffee shop\" returns coffeeshop.com, mycoffeeshop.net, etc.\n- Domain names: \"example.com\" returns example.com and variations across extensions\n", "required": true, "schema": {"type": "string", "maxLength": 100, "minLength": 1}, "example": "coffee shop"}, {"name": "extensions", "in": "query", "description": "Limits results to specific domain extensions from the supported set. If not specified,\nreturns results across all supported extensions. Extensions not in the supported\nset are silently ignored.\n", "schema": {"type": "array", "items": {"type": "string"}, "maxItems": 20}, "example": ["com", "net", "org"], "explode": false, "style": "form"}, {"name": "limit", "in": "query", "description": "Maximum number of domain suggestions to return. Defaults to 20 if not specified.\n", "schema": {"type": "integer", "default": 20, "maximum": 50, "minimum": 1}, "example": 20}], "responses": {"200": {"description": "Successfully returned domain search results.", "content": {"application/json": {"examples": {"brand_search": {"description": "Searching for a brand or company name", "summary": "Brand name search - \"acme corp\"", "value": {"errors": [], "messages": [], "result": {"domains": [{"name": "acmecorp.com", "pricing": {"currency": "USD", "registration_cost": "8.57", "renewal_cost": "8.57"}, "registrable": true, "tier": "standard"}, {"name": "acmecorp.dev", "pricing": {"currency": "USD", "registration_cost": "10.11", "renewal_cost": "10.11"}, "registrable": true, "tier": "standard"}, {"name": "acmecorp.app", "pricing": {"currency": "USD", "registration_cost": "11.00", "renewal_cost": "11.00"}, "registrable": true, "tier": "standard"}]}, "success": true}}, "extension_filtered_search": {"description": "Searching with extensions=com,net limits results to those domain extensions", "summary": "Search with domain extension filter", "value": {"errors": [], "messages": [], "result": {"domains": [{"name": "bestpizza.com", "pricing": {"currency": "USD", "registration_cost": "8.57", "renewal_cost": "8.57"}, "registrable": true, "tier": "standard"}, {"name": "bestpizza.net", "pricing": {"currency": "USD", "registration_cost": "9.95", "renewal_cost": "9.95"}, "registrable": true, "tier": "standard"}, {"name": "bestpizzashop.com", "pricing": {"currency": "USD", "registration_cost": "8.57", "renewal_cost": "8.57"}, "registrable": true, "tier": "standard"}]}, "success": true}}, "keyword_search": {"description": "Searching for a keyword phrase returns relevant domain suggestions", "summary": "Keyword search - \"coffee shop\"", "value": {"errors": [], "messages": [], "result": {"domains": [{"name": "coffeeshop.com", "pricing": {"currency": "USD", "registration_cost": "8.57", "renewal_cost": "8.57"}, "registrable": true, "tier": "standard"}, {"name": "coffeeshoponline.com", "pricing": {"currency": "USD", "registration_cost": "8.57", "renewal_cost": "8.57"}, "registrable": true, "tier": "standard"}, {"name": "mycoffeeshop.net", "pricing": {"currency": "USD", "registration_cost": "9.95", "renewal_cost": "9.95"}, "registrable": true, "tier": "standard"}, {"name": "thecoffeeshop.shop", "pricing": {"currency": "USD", "registration_cost": "11.00", "renewal_cost": "11.00"}, "registrable": true, "tier": "standard"}]}, "success": true}}, "no_results": {"description": "When no registrable domains match the search", "summary": "No available domains found", "value": {"errors": [], "messages": [], "result": {"domains": []}, "success": true}}, "premium_results": {"description": "Some suggestions may be premium domains with higher pricing. Premium registration is not currently supported by this API.", "summary": "Results including premium domains", "value": {"errors": [], "messages": [], "result": {"domains": [{"name": "crypto.com", "pricing": {"currency": "USD", "registration_cost": "100000.00", "renewal_cost": "5000.00"}, "registrable": true, "tier": "premium"}, {"name": "cryptotrading.com", "pricing": {"currency": "USD", "registration_cost": "8.57", "renewal_cost": "8.57"}, "registrable": true, "tier": "standard"}, {"name": "mycrypto.net", "pricing": {"currency": "USD", "registration_cost": "9.95", "renewal_cost": "9.95"}, "registrable": true, "tier": "standard"}]}, "success": true}}}, "schema": {"$ref": "#/components/schemas/registrar-api_domain_search_response"}}}}, "400": {"description": "Invalid request parameters. Common causes:\n- Missing required `q` parameter\n- Query exceeds 100 character limit\n- Invalid extension format\n", "content": {"application/json": {"examples": {"invalid_query": {"description": "The query is invalid for search because it is too long, not supported as IDN/punycode, or becomes empty after sanitization.", "summary": "Invalid search query", "value": {"errors": [{"code": 1002, "message": "Parameter q exceeds maximum length of 100 characters"}], "messages": [], "result": null, "success": false}}, "missing_query": {"description": "The required `q` parameter was not provided in the request.", "summary": "Missing search query", "value": {"errors": [{"code": 1001, "message": "Missing required parameter: q"}], "messages": [], "result": null, "success": false}}}, "schema": {"$ref": "#/components/schemas/registrar-api_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Domain Discovery"]}
```
