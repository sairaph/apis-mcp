---
title: Delete a schema
page_id: operation-delete-zones-zone-id-api-gateway-user-schemas-schema-id-2cf473b8
path: operations/api-shield-schema-validation-2-0
description: Permanently removes an uploaded OpenAPI schema from API Shield schema validation. Operations using this schema will lose their validation rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/api_gateway/user_schemas/{schema_id}
operation_ids:
    - api-shield-schema-delete-a-schema
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a schema

`DELETE /zones/{zone_id}/api_gateway/user_schemas/{schema_id}`

Operation ID: `api-shield-schema-delete-a-schema`

Permanently removes an uploaded OpenAPI schema from API Shield schema validation. Operations using this schema will lose their validation rules.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_old_schema_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-schema-delete-a-schema", "summary": "Delete a schema", "description": "Permanently removes an uploaded OpenAPI schema from API Shield schema validation. Operations using this schema will lose their validation rules.", "responses": {"200": {"description": "Delete a schema response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-single"}}}}, "4XX": {"description": "Delete a schema response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Schema Validation 2.0"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "api-gateway.user-schemas", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true, "x-stainless-deprecation-message": "Use [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/) instead."}
```
