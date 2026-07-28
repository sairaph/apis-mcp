---
title: Detach Domain
page_id: operation-delete-accounts-account-id-workers-domains-domain-id-ee556d73
path: operations/domains
description: Detaches a domain from a Worker. Both the Worker and all of its previews are no longer routable using this domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/workers/domains/{domain_id}
operation_ids:
    - workers.domains.delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Detach Domain

`DELETE /accounts/{account_id}/workers/domains/{domain_id}`

Operation ID: `workers.domains.delete`

Detaches a domain from a Worker. Both the Worker and all of its previews are no longer routable using this domain.

## Definition

```yaml
{"operationId": "workers.domains.delete", "summary": "Detach Domain", "description": "Detaches a domain from a Worker. Both the Worker and all of its previews are no longer routable using this domain.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "domain_id", "in": "path", "required": true, "schema": {"description": "ID of the domain.", "type": "string", "example": "dbe10b4bc17c295377eabd600e1787fd"}}], "responses": {"200": {"description": "Detach domain response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common"}}}}, "4XX": {"description": "Detach domain failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Domains"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.domains", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
