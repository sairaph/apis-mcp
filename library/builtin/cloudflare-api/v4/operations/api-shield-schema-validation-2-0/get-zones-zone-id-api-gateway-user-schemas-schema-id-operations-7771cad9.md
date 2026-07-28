---
title: Retrieve all operations from a schema.
page_id: operation-get-zones-zone-id-api-gateway-user-schemas-schema-id-operations-562d2b6e
path: operations/api-shield-schema-validation-2-0
description: Retrieves all operations from the schema. Operations that already exist in API Shield Endpoint Management will be returned as full operations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/user_schemas/{schema_id}/operations
operation_ids:
    - api-shield-schema-validation-extract-operations-from-schema
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve all operations from a schema.

`GET /zones/{zone_id}/api_gateway/user_schemas/{schema_id}/operations`

Operation ID: `api-shield-schema-validation-extract-operations-from-schema`

Retrieves all operations from the schema. Operations that already exist in API Shield Endpoint Management will be returned as full operations.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_old_schema_id"}, {"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-schema-validation-extract-operations-from-schema", "summary": "Retrieve all operations from a schema.", "description": "Retrieves all operations from the schema. Operations that already exist in API Shield Endpoint Management will be returned as full operations.", "parameters": [{"$ref": "#/components/parameters/api-shield_operation_feature_parameter"}, {"$ref": "#/components/parameters/api-shield_host_parameter"}, {"$ref": "#/components/parameters/api-shield_method_parameter"}, {"$ref": "#/components/parameters/api-shield_endpoint_parameter"}, {"$ref": "#/components/parameters/api-shield_page"}, {"$ref": "#/components/parameters/api-shield_per_page"}, {"name": "operation_status", "in": "query", "description": "Filter results by whether operations exist in API Shield Endpoint Management or not. `new` will just return operations from the schema that do not exist in API Shield Endpoint Management. `existing` will just return operations from the schema that already exist in API Shield Endpoint Management.", "schema": {"type": "string", "example": "new", "enum": ["new", "existing"]}}], "responses": {"200": {"description": "Retrieve all operations from a schema response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"anyOf": [{"$ref": "#/components/schemas/api-shield_operation"}, {"$ref": "#/components/schemas/api-shield_basic_operation"}]}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Retrieve all operations from a schema response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Schema Validation 2.0"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "api-gateway.user-schemas.operations", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-stainless-deprecation-message": "Use [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/) instead."}
```
