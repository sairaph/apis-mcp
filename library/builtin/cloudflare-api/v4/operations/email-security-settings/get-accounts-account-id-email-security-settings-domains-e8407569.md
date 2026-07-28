---
title: List protected email domains
page_id: operation-get-accounts-account-id-email-security-settings-domains-bb7bd518
path: operations/email-security-settings
description: Returns a paginated list of email domains protected by Email Security. Includes domain configuration, delivery modes, and authorization status. Supports filtering by delivery mode and integration ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/domains
operation_ids:
    - email_security_list_domains
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List protected email domains

`GET /accounts/{account_id}/email-security/settings/domains`

Operation ID: `email_security_list_domains`

Returns a paginated list of email domains protected by Email Security. Includes domain configuration, delivery modes, and authorization status. Supports filtering by delivery mode and integration ID.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_list_domains", "summary": "List protected email domains", "description": "Returns a paginated list of email domains protected by Email Security. Includes domain configuration, delivery modes, and authorization status. Supports filtering by delivery mode and integration ID.", "parameters": [{"$ref": "#/components/parameters/email-security_page"}, {"$ref": "#/components/parameters/email-security_per_page"}, {"$ref": "#/components/parameters/email-security_search"}, {"name": "order", "in": "query", "description": "Field to sort by.", "schema": {"type": "string", "enum": ["domain", "created_at"]}}, {"$ref": "#/components/parameters/email-security_direction"}, {"name": "allowed_delivery_mode", "in": "query", "description": "Delivery mode to filter by.", "schema": {"$ref": "#/components/schemas/email-security_DeliveryMode"}}, {"name": "domain", "in": "query", "description": "Domain names to filter by.", "schema": {"type": "array", "items": {"type": "string"}}}, {"name": "active_delivery_mode", "in": "query", "description": "Currently active delivery mode to filter by.", "schema": {"$ref": "#/components/schemas/email-security_DeliveryMode"}}, {"name": "integration_id", "in": "query", "description": "Integration ID to filter by.", "schema": {"type": "string", "format": "uuid"}}, {"name": "status", "in": "query", "description": "Filters response to domains with the provided status.", "schema": {"$ref": "#/components/schemas/email-security_DomainStatus"}}], "responses": {"200": {"description": "List of domains.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_DomainList"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-stability": "beta"}
```
