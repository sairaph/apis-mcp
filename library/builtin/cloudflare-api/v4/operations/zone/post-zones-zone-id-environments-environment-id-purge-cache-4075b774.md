---
title: Purge Cached Content by Environment
page_id: operation-post-zones-zone-id-environments-environment-id-purge-cache-bb103590
path: operations/zone
description: |-
    Purge cached content scoped to a specific environment. Supports the same purge types as the zone-level endpoint (purge everything, by URL, by tag, host, or prefix).

    ### Availability and limits
    Please refer to [purge cache availability and limits documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/#availability-and-limits).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/environments/{environment_id}/purge_cache
operation_ids:
    - zone-environment-purge
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Purge Cached Content by Environment

`POST /zones/{zone_id}/environments/{environment_id}/purge_cache`

Operation ID: `zone-environment-purge`

Purge cached content scoped to a specific environment. Supports the same purge types as the zone-level endpoint (purge everything, by URL, by tag, host, or prefix).

### Availability and limits
Please refer to [purge cache availability and limits documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/#availability-and-limits).

## Definition

```yaml
{"operationId": "zone-environment-purge", "summary": "Purge Cached Content by Environment", "description": "Purge cached content scoped to a specific environment. Supports the same purge types as the zone-level endpoint (purge everything, by URL, by tag, host, or prefix).\n\n### Availability and limits\nPlease refer to [purge cache availability and limits documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/#availability-and-limits).\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-purge_identifier"}}, {"name": "environment_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-purge_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"Flex Purge with Hosts": {"summary": "Flex purge example with hosts list", "value": {"hosts": ["www.example.com", "images.example.com"]}}, "Flex Purge with Prefixes": {"summary": "Flex purge example with prefixes list", "value": {"prefixes": ["www.example.com/foo", "images.example.com/bar/baz"]}}, "Flex Purge with Tags": {"summary": "Flex purge example with tags list", "value": {"tags": ["a-cache-tag", "another-cache-tag"]}}, "Purge Everything": {"summary": "Purge everything example", "value": {"purge_everything": true}}, "Single File Purge": {"summary": "Single file purge example with files list", "value": {"files": ["http://www.example.com/css/styles.css", "http://www.example.com/js/index.js"]}}, "Single File Purge with UrlAndHeaders": {"summary": "Single file purge example with url and headers list", "value": {"files": [{"headers": {"Accept-Language": "zh-CN", "CF-Device-Type": "desktop", "CF-IPCountry": "US"}, "url": "http://www.example.com/cat_picture.jpg"}, {"headers": {"Accept-Language": "en-US", "CF-Device-Type": "mobile", "CF-IPCountry": "EU"}, "url": "http://www.example.com/dog_picture.jpg"}]}}}, "schema": {"anyOf": [{"$ref": "#/components/schemas/cache-purge_FlexPurgeByTags"}, {"$ref": "#/components/schemas/cache-purge_FlexPurgeByHostnames"}, {"$ref": "#/components/schemas/cache-purge_FlexPurgeByPrefixes"}, {"$ref": "#/components/schemas/cache-purge_Everything"}, {"$ref": "#/components/schemas/cache-purge_SingleFile"}, {"$ref": "#/components/schemas/cache-purge_SingleFileWithUrlAndHeaders"}]}}}}, "responses": {"200": {"description": "Request to purge cached content successful.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache-purge_api-response-single-id"}}}}, "4XX": {"description": "Request to purge cached content failed.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cache-purge_api-response-single-id"}, {"$ref": "#/components/schemas/cache-purge_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone"], "x-api-token-group": ["Cache Purge"], "x-cfPermissionsRequired": {"enum": ["#cache_purge:edit"]}}
```
