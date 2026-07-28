---
title: List SAML certificate sets
page_id: operation-get-accounts-account-id-access-saml-certificates-59693a0b
path: operations/access-saml-encryption-certificates
description: |-
    Returns a paginated list of the organization's SAML encryption certificate sets.
    Each certificate set includes the current and (if present) previous certificates.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/saml_certificates
operation_ids:
    - access-saml-certificates-list-certificate-sets
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List SAML certificate sets

`GET /accounts/{account_id}/access/saml_certificates`

Operation ID: `access-saml-certificates-list-certificate-sets`

Returns a paginated list of the organization's SAML encryption certificate sets.
Each certificate set includes the current and (if present) previous certificates.

## Definition

```yaml
{"operationId": "access-saml-certificates-list-certificate-sets", "summary": "List SAML certificate sets", "description": "Returns a paginated list of the organization's SAML encryption certificate sets.\nEach certificate set includes the current and (if present) previous certificates.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "page", "in": "query", "description": "Page number of paginated results.", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Maximum number of results per page.", "schema": {"type": "integer", "default": 25, "maximum": 1000, "minimum": 1}}, {"name": "id", "in": "query", "description": "Filter by SAML certificate set UID. Accepts a comma-separated list of UIDs.", "schema": {"type": "string", "example": "a5bb4b3f-c2d1-4e6a-8f9b-1d3e4f5a6b7c,f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}}], "responses": {"200": {"description": "List SAML certificate sets response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_certificate_set_list_response"}}}}, "4XX": {"description": "List SAML certificate sets response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access SAML encryption certificates"]}
```
