---
title: List hosts covered by uploaded schemas
page_id: operation-get-zones-zone-id-schema-validation-schemas-hosts-471f1d96
path: operations/schema-validation
description: Lists all unique hosts found in uploaded OpenAPI schemas for the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/schema_validation/schemas/hosts
operation_ids:
    - schema-validation-list-schema-hosts
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List hosts covered by uploaded schemas

`GET /zones/{zone_id}/schema_validation/schemas/hosts`

Operation ID: `schema-validation-list-schema-hosts`

Lists all unique hosts found in uploaded OpenAPI schemas for the zone.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "schema-validation-list-schema-hosts", "summary": "List hosts covered by uploaded schemas", "description": "Lists all unique hosts found in uploaded OpenAPI schemas for the zone.", "parameters": [{"$ref": "#/components/parameters/api-shield_page"}, {"$ref": "#/components/parameters/api-shield_per_page"}], "responses": {"200": {"$ref": "#/components/responses/api-shield_schema_hosts_get_success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.schemas.hosts", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
