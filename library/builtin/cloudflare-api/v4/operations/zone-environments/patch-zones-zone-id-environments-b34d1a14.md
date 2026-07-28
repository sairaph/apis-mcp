---
title: Partially update zone environments
page_id: operation-patch-zones-zone-id-environments-9eac93b7
path: operations/zone-environments
description: Applies partial updates to zone environments.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/environments
operation_ids:
    - zonesEnvironmentsEdit
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Partially update zone environments

`PATCH /zones/{zone_id}/environments`

Operation ID: `zonesEnvironmentsEdit`

Applies partial updates to zone environments.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/kamino_zone_id"}]
```

## Definition

```yaml
{"operationId": "zonesEnvironmentsEdit", "summary": "Partially update zone environments", "description": "Applies partial updates to zone environments.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/kamino_environments_request"}}}}, "responses": {"200": {"description": "Successfully updated zone environments.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/kamino_environments_response"}}}}, "4XX": {"$ref": "#/components/responses/kamino_client_error"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Environments"], "x-api-token-group": ["Zone Versioning Write"]}
```
