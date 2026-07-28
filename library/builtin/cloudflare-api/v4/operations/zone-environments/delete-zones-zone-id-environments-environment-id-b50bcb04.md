---
title: Delete zone environment
page_id: operation-delete-zones-zone-id-environments-environment-id-adb41de6
path: operations/zone-environments
description: Deletes a zone environment by reference identifier.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/environments/{environment_id}
operation_ids:
    - zonesEnvironmentsDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete zone environment

`DELETE /zones/{zone_id}/environments/{environment_id}`

Operation ID: `zonesEnvironmentsDelete`

Deletes a zone environment by reference identifier.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/kamino_zone_id"}, {"$ref": "#/components/parameters/kamino_environment_id"}]
```

## Definition

```yaml
{"operationId": "zonesEnvironmentsDelete", "summary": "Delete zone environment", "description": "Deletes a zone environment by reference identifier.", "responses": {"200": {"description": "Successfully deleted zone environment.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/kamino_environments_response"}}}}, "4XX": {"$ref": "#/components/responses/kamino_client_error"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Environments"], "x-api-token-group": ["Zone Versioning Write"]}
```
