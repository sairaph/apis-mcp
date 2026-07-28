---
title: List zone environments
page_id: operation-get-zones-zone-id-environments-e74c8de2
path: operations/zone-environments
description: Lists configured environments for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/environments
operation_ids:
    - zonesEnvironmentsList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List zone environments

`GET /zones/{zone_id}/environments`

Operation ID: `zonesEnvironmentsList`

Lists configured environments for a zone.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/kamino_zone_id"}]
```

## Definition

```yaml
{"operationId": "zonesEnvironmentsList", "summary": "List zone environments", "description": "Lists configured environments for a zone.", "responses": {"200": {"description": "Successfully listed zone environments.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/kamino_environments_response"}}}}, "4XX": {"$ref": "#/components/responses/kamino_client_error"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Environments"], "x-api-token-group": ["Zone Versioning Write", "Zone Versioning Read"]}
```
