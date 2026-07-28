---
title: Update domain
page_id: operation-put-accounts-account-id-registrar-domains-domain-name-ee995f46
path: operations/registrar-domains
description: Update individual domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/registrar/domains/{domain_name}
operation_ids:
    - registrar-domains-update-domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update domain

`PUT /accounts/{account_id}/registrar/domains/{domain_name}`

Operation ID: `registrar-domains-update-domain`

Update individual domain.

## Definition

```yaml
{"operationId": "registrar-domains-update-domain", "summary": "Update domain", "description": "Update individual domain.", "parameters": [{"name": "domain_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api_domain_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/registrar-api_domain_update_properties"}]}}}}, "responses": {"200": {"description": "Update domain response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/registrar-api_domain_response_single"}}}}, "4XX": {"description": "Update domain response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/registrar-api_domain_response_single"}, {"$ref": "#/components/schemas/registrar-api_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Registrar Domains"], "x-api-token-group": null, "x-stainless-deprecation-message": "This operation is deprecated and will reach end of life on September 27, 2026. Use the new Registrar API endpoints (domain-search, domain-check, registrations) instead. Refer to https://developers.cloudflare.com/fundamentals/api/reference/deprecations/ for details."}
```
