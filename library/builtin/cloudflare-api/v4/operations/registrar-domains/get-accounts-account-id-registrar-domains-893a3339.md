---
title: List domains
page_id: operation-get-accounts-account-id-registrar-domains-9dc3b2d9
path: operations/registrar-domains
description: List domains handled by Registrar.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/registrar/domains
operation_ids:
    - registrar-domains-list-domains
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List domains

`GET /accounts/{account_id}/registrar/domains`

Operation ID: `registrar-domains-list-domains`

List domains handled by Registrar.

## Definition

```yaml
{"operationId": "registrar-domains-list-domains", "summary": "List domains", "description": "List domains handled by Registrar.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api_identifier"}}], "responses": {"200": {"description": "List domains response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/registrar-api_domain_response_collection"}}}}, "4XX": {"description": "List domains response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/registrar-api_domain_response_collection"}, {"$ref": "#/components/schemas/registrar-api_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Registrar Domains"], "x-api-token-group": null, "x-stainless-deprecation-message": "This operation is deprecated and will reach end of life on September 27, 2026. Use the new Registrar API endpoints (domain-search, domain-check, registrations) instead. Refer to https://developers.cloudflare.com/fundamentals/api/reference/deprecations/ for details."}
```
