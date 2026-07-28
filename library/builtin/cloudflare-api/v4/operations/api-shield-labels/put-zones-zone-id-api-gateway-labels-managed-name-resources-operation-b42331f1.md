---
title: Replace operation(s) attached to a managed label
page_id: operation-put-zones-zone-id-api-gateway-labels-managed-name-resources-operation-917b0b65
path: operations/api-shield-labels
description: Replace all operations(s) attached to a managed label
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/api_gateway/labels/managed/{name}/resources/operation
operation_ids:
    - api-shield-labels-replace-operations-attached-to-managed-label
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Replace operation(s) attached to a managed label

`PUT /zones/{zone_id}/api_gateway/labels/managed/{name}/resources/operation`

Operation ID: `api-shield-labels-replace-operations-attached-to-managed-label`

Replace all operations(s) attached to a managed label

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_label_name_parameter"}]
```

## Definition

```yaml
{"operationId": "api-shield-labels-replace-operations-attached-to-managed-label", "summary": "Replace operation(s) attached to a managed label", "description": "Replace all operations(s) attached to a managed label", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_replace_operations_attached_to_label_request"}}}}, "responses": {"200": {"description": "Replace all operations(s) attached to a managed label response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/api-shield_full_managed_label"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Replace all operations(s) attached to a managed label failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield Labels"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.labels.managed.resources.operation", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
