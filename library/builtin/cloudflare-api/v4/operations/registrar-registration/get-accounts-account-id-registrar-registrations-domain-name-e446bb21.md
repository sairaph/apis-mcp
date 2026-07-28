---
title: Get Registration
page_id: operation-get-accounts-account-id-registrar-registrations-domain-name-18f4f625
path: operations/registrar-registration
description: |-
    Returns the current state of a domain registration.

    This is the canonical read endpoint for a domain you own. It returns
    the full registration resource including current settings and expiration.
    When the registration resource is ready, both `created_at` and `expires_at`
    are present in the response.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/registrar/registrations/{domain_name}
operation_ids:
    - registrar-domain-registration-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Registration

`GET /accounts/{account_id}/registrar/registrations/{domain_name}`

Operation ID: `registrar-domain-registration-get`

Returns the current state of a domain registration.

This is the canonical read endpoint for a domain you own. It returns
the full registration resource including current settings and expiration.
When the registration resource is ready, both `created_at` and `expires_at`
are present in the response.

## Definition

```yaml
{"operationId": "registrar-domain-registration-get", "summary": "Get Registration", "description": "Returns the current state of a domain registration.\n\nThis is the canonical read endpoint for a domain you own. It returns\nthe full registration resource including current settings and expiration.\nWhen the registration resource is ready, both `created_at` and `expires_at`\nare present in the response.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api_identifier"}}, {"name": "domain_name", "in": "path", "description": "Domain name to retrieve.", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api_domain_name"}}], "responses": {"200": {"description": "Registration details.", "content": {"application/json": {"examples": {"active": {"summary": "Active domain", "value": {"errors": [], "messages": [], "result": {"auto_renew": true, "created_at": "2025-01-15T10:00:00Z", "domain_name": "example.com", "expires_at": "2026-01-15T10:00:00Z", "locked": true, "privacy_mode": "redaction", "status": "active"}, "success": true}}}, "schema": {"$ref": "#/components/schemas/registrar-api_registration-response-single"}}}}, "4XX": {"description": "Get registration failure.", "content": {"application/json": {"examples": {"domain_not_found": {"summary": "Domain not found", "value": {"errors": [{"code": 10000, "message": "Domain not found"}], "messages": [], "result": null, "success": false}}, "invalid_domain_name": {"summary": "Invalid domain name", "value": {"errors": [{"code": 10000, "message": "Invalid domain name"}], "messages": [], "result": null, "success": false}}}, "schema": {"$ref": "#/components/schemas/registrar-api_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Registrar Registration"]}
```
