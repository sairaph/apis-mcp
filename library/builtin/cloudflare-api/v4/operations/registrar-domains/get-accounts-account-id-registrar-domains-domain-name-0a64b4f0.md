---
title: Get domain
page_id: operation-get-accounts-account-id-registrar-domains-domain-name-755106a8
path: operations/registrar-domains
description: Show individual domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/registrar/domains/{domain_name}
operation_ids:
    - registrar-domains-get-domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get domain

`GET /accounts/{account_id}/registrar/domains/{domain_name}`

Operation ID: `registrar-domains-get-domain`

Show individual domain.

## Definition

```yaml
{"operationId": "registrar-domains-get-domain", "summary": "Get domain", "description": "Show individual domain.", "parameters": [{"name": "domain_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api_domain_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api_identifier"}}], "responses": {"200": {"description": "Get domain response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/registrar-api_domain_response_single"}}}}, "4XX": {"description": "Get domain response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/registrar-api_domain_response_single"}, {"$ref": "#/components/schemas/registrar-api_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Registrar Domains"], "x-api-token-group": null, "x-stainless-deprecation-message": "This operation is deprecated and will reach end of life on September 27, 2026. Use the new Registrar API endpoints (domain-search, domain-check, registrations) instead. Refer to https://developers.cloudflare.com/fundamentals/api/reference/deprecations/ for details."}
```
