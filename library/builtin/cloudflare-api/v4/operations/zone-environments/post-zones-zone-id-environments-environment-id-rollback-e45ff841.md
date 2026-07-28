---
title: Roll back zone environment
page_id: operation-post-zones-zone-id-environments-environment-id-rollback-a70cad8b
path: operations/zone-environments
description: Rolls a zone environment back to its previous version.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/environments/{environment_id}/rollback
operation_ids:
    - zonesEnvironmentsRollback
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Roll back zone environment

`POST /zones/{zone_id}/environments/{environment_id}/rollback`

Operation ID: `zonesEnvironmentsRollback`

Rolls a zone environment back to its previous version.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/kamino_zone_id"}, {"$ref": "#/components/parameters/kamino_environment_id"}]
```

## Definition

```yaml
{"operationId": "zonesEnvironmentsRollback", "summary": "Roll back zone environment", "description": "Rolls a zone environment back to its previous version.", "responses": {"200": {"description": "Successfully rolled back zone environment.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/kamino_environments_response"}}}}, "4XX": {"$ref": "#/components/responses/kamino_client_error"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Environments"], "x-api-token-group": ["Zone Versioning Write"]}
```
