---
title: Upsert zone environments
page_id: operation-put-zones-zone-id-environments-ce3fea20
path: operations/zone-environments
description: Replaces the full environment configuration for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/environments
operation_ids:
    - zonesEnvironmentsUpdate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upsert zone environments

`PUT /zones/{zone_id}/environments`

Operation ID: `zonesEnvironmentsUpdate`

Replaces the full environment configuration for a zone.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/kamino_zone_id"}]
```

## Definition

```yaml
{"operationId": "zonesEnvironmentsUpdate", "summary": "Upsert zone environments", "description": "Replaces the full environment configuration for a zone.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/kamino_environments_request"}}}}, "responses": {"200": {"description": "Successfully upserted zone environments.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/kamino_environments_response"}}}}, "4XX": {"$ref": "#/components/responses/kamino_client_error"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Environments"], "x-api-token-group": ["Zone Versioning Write"]}
```
