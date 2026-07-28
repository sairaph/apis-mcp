---
title: Retrieve information about all schemas on a zone
page_id: operation-get-zones-zone-id-api-gateway-user-schemas-5c4a9ffe
path: operations/api-shield-schema-validation-2-0
description: Lists all OpenAPI schemas uploaded to API Shield for the zone, including their validation status and associated operations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/user_schemas
operation_ids:
    - api-shield-schema-validation-retrieve-information-about-all-schemas
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve information about all schemas on a zone

`GET /zones/{zone_id}/api_gateway/user_schemas`

Operation ID: `api-shield-schema-validation-retrieve-information-about-all-schemas`

Lists all OpenAPI schemas uploaded to API Shield for the zone, including their validation status and associated operations.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-schema-validation-retrieve-information-about-all-schemas", "summary": "Retrieve information about all schemas on a zone", "description": "Lists all OpenAPI schemas uploaded to API Shield for the zone, including their validation status and associated operations.", "parameters": [{"$ref": "#/components/parameters/api-shield_page"}, {"$ref": "#/components/parameters/api-shield_per_page"}, {"$ref": "#/components/parameters/api-shield_old_omit_source"}, {"name": "validation_enabled", "in": "query", "schema": {"$ref": "#/components/schemas/api-shield_old_validation_enabled"}}], "responses": {"200": {"description": "Retrieve information about all schemas on a zone response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_old_public_schema"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Retrieve information about all schemas on a zone response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Schema Validation 2.0"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "api-gateway.user-schemas", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-stainless-deprecation-message": "Use [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/) instead."}
```
