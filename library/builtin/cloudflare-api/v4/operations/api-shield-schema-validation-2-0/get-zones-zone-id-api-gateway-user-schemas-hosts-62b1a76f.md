---
title: Retrieve schema hosts in a zone
page_id: operation-get-zones-zone-id-api-gateway-user-schemas-hosts-fb2462cb
path: operations/api-shield-schema-validation-2-0
description: Lists all unique hosts found in uploaded OpenAPI schemas for the zone. Useful for understanding which domains have schema coverage.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/api_gateway/user_schemas/hosts
operation_ids:
    - api-shield-schema-validation-retrieve-user-schema-hosts
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve schema hosts in a zone

`GET /zones/{zone_id}/api_gateway/user_schemas/hosts`

Operation ID: `api-shield-schema-validation-retrieve-user-schema-hosts`

Lists all unique hosts found in uploaded OpenAPI schemas for the zone. Useful for understanding which domains have schema coverage.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_page"}, {"$ref": "#/components/parameters/api-shield_per_page"}]
```

## Definition

```yaml
{"operationId": "api-shield-schema-validation-retrieve-user-schema-hosts", "summary": "Retrieve schema hosts in a zone", "description": "Lists all unique hosts found in uploaded OpenAPI schemas for the zone. Useful for understanding which domains have schema coverage.", "responses": {"200": {"description": "Retrieve schema hosts in a zone response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_old_response_user_schemas_hosts"}}}}]}}}}, "4XX": {"description": "Retrieve schema hosts in a zone response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Schema Validation 2.0"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "api-gateway.user-schemas.hosts", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-stainless-deprecation-message": "Use [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/) instead."}
```
