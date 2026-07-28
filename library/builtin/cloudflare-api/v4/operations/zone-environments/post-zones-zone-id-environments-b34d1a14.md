---
title: Create zone environments
page_id: operation-post-zones-zone-id-environments-c957ba86
path: operations/zone-environments
description: Creates environments for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/environments
operation_ids:
    - zonesEnvironmentsCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create zone environments

`POST /zones/{zone_id}/environments`

Operation ID: `zonesEnvironmentsCreate`

Creates environments for a zone.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/kamino_zone_id"}]
```

## Definition

```yaml
{"operationId": "zonesEnvironmentsCreate", "summary": "Create zone environments", "description": "Creates environments for a zone.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/kamino_environments_request"}}}}, "responses": {"200": {"description": "Successfully created zone environments.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/kamino_environments_response"}}}}, "4XX": {"$ref": "#/components/responses/kamino_client_error"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Environments"], "x-api-token-group": ["Zone Versioning Write"]}
```
