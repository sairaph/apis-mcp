---
title: List Spectrum applications
page_id: operation-get-zones-zone-id-spectrum-apps-852af80e
path: operations/spectrum-applications
description: Retrieves a list of currently existing Spectrum applications inside a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/spectrum/apps
operation_ids:
    - spectrum-applications-list-spectrum-applications
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Spectrum applications

`GET /zones/{zone_id}/spectrum/apps`

Operation ID: `spectrum-applications-list-spectrum-applications`

Retrieves a list of currently existing Spectrum applications inside a zone.

## Definition

```yaml
{"operationId": "spectrum-applications-list-spectrum-applications", "summary": "List Spectrum applications", "description": "Retrieves a list of currently existing Spectrum applications inside a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/spectrum-config_zone_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results. This parameter is required in order to use other pagination parameters. If included in the query, `result_info` will be present in the response.", "type": "number", "example": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Sets the maximum number of results per page.", "type": "number", "default": 20, "maximum": 100, "minimum": 1}}, {"name": "direction", "in": "query", "schema": {"description": "Sets the direction by which results are ordered.", "type": "string", "example": "desc", "default": "asc", "enum": ["asc", "desc"]}}, {"name": "order", "in": "query", "schema": {"description": "Application field by which results are ordered.", "type": "string", "example": "protocol", "default": "dns", "enum": ["protocol", "app_id", "created_on", "modified_on", "dns"]}}], "responses": {"200": {"description": "List Spectrum applications response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-config_app_config_collection"}}}}, "4XX": {"description": "List Spectrum applications response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-config_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Spectrum Applications"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "spectrum.apps", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
