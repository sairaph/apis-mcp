---
title: Retrieve all operations from the schema.
page_id: operation-get-zones-zone-id-schema-validation-schemas-schema-id-operations-44f3a637
path: operations/schema-validation
description: Retrieves all operations from the schema. Operations that already exist in API Shield Endpoint Management will be returned as full operations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/schema_validation/schemas/{schema_id}/operations
operation_ids:
    - schema-validation-extract-operations-from-schema
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve all operations from the schema.

`GET /zones/{zone_id}/schema_validation/schemas/{schema_id}/operations`

Operation ID: `schema-validation-extract-operations-from-schema`

Retrieves all operations from the schema. Operations that already exist in API Shield Endpoint Management will be returned as full operations.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_schema_id_path"}]
```

## Definition

```yaml
{"operationId": "schema-validation-extract-operations-from-schema", "summary": "Retrieve all operations from the schema.", "description": "Retrieves all operations from the schema. Operations that already exist in API Shield Endpoint Management will be returned as full operations.", "parameters": [{"$ref": "#/components/parameters/api-shield_operation_feature_parameter"}, {"$ref": "#/components/parameters/api-shield_host_parameter"}, {"$ref": "#/components/parameters/api-shield_method_parameter"}, {"$ref": "#/components/parameters/api-shield_endpoint_parameter"}, {"$ref": "#/components/parameters/api-shield_page"}, {"$ref": "#/components/parameters/api-shield_per_page"}, {"name": "operation_status", "in": "query", "description": "Filter results by whether operations exist in Web Asset Management or not. `new` will just return operations from the schema that do not exist otherwise. `existing` will just return operations from the schema that already exist.", "schema": {"type": "string", "example": "new", "enum": ["new", "existing"]}}], "responses": {"200": {"$ref": "#/components/responses/api-shield_schemas_extract_operations_get_success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Schema Validation"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "schema-validation.schemas.operations", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
