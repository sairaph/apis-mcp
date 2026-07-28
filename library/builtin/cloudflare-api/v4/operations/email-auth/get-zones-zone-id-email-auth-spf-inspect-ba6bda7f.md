---
title: Inspect SPF Record
page_id: operation-get-zones-zone-id-email-auth-spf-inspect-8d20685d
path: operations/email-auth
description: |-
    Inspects a specific SPF TXT record and returns a parsed tree structure
    in the spflimit-worker format.

    The record ID must be provided via the `id` query parameter.

    Returns a recursive tree showing:
    - Parsed components with their qualifiers and types
    - Nested includes recursively resolved within components
    - Per-component and total lookup counts
    - Detailed error information with context
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/email/auth/spf/inspect
operation_ids:
    - inspect_spf
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Inspect SPF Record

`GET /zones/{zone_id}/email/auth/spf/inspect`

Operation ID: `inspect_spf`

Inspects a specific SPF TXT record and returns a parsed tree structure
in the spflimit-worker format.

The record ID must be provided via the `id` query parameter.

Returns a recursive tree showing:
- Parsed components with their qualifiers and types
- Nested includes recursively resolved within components
- Per-component and total lookup counts
- Detailed error information with context

## Definition

```yaml
{"operationId": "inspect_spf", "summary": "Inspect SPF Record", "description": "Inspects a specific SPF TXT record and returns a parsed tree structure\nin the spflimit-worker format.\n\nThe record ID must be provided via the `id` query parameter.\n\nReturns a recursive tree showing:\n- Parsed components with their qualifiers and types\n- Nested includes recursively resolved within components\n- Per-component and total lookup counts\n- Detailed error information with context\n", "parameters": [{"$ref": "#/components/parameters/email-auth_zone_id"}, {"$ref": "#/components/parameters/email-auth_record_id"}], "responses": {"200": {"description": "SPF record inspected successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-auth_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-auth_SpfTree"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-auth_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Auth"]}
```
