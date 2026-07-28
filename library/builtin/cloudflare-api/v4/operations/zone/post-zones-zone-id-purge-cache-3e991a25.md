---
title: Purge Cached Content
page_id: operation-post-zones-zone-id-purge-cache-c9344b85
path: operations/zone
description: |-
    ### Purge All Cached Content
    Removes ALL files from Cloudflare's cache. All tiers can purge everything.
    ```
    {"purge_everything": true}
    ```

    ### Purge Cached Content by URL
    Granularly removes one or more files from Cloudflare's cache by specifying URLs. All tiers can purge by URL.

    To purge files with custom cache keys, include the headers used to compute the cache key as in the example. If you have a device type or geo in your cache key, you will need to include the CF-Device-Type or CF-IPCountry headers. If you have lang in your cache key, you will need to include the Accept-Language header.

    **NB:** When including the Origin header, be sure to include the **scheme** and **hostname**. The port number can be omitted if it is the default port (80 for http, 443 for https), but must be included otherwise.

    Single file purge example with files:
    ```
    {"files": ["http://www.example.com/css/styles.css", "http://www.example.com/js/index.js"]}
    ```
    Single file purge example with url and header pairs:
    ```
    {"files": [{"url": "http://www.example.com/cat_picture.jpg", "headers": {"CF-IPCountry": "US", "CF-Device-Type": "desktop", "Accept-Language": "zh-CN"}}, {"url": "http://www.example.com/dog_picture.jpg", "headers": {"CF-IPCountry": "EU", "CF-Device-Type": "mobile", "Accept-Language": "en-US"}}]}
    ```

    ### Purge Cached Content by Tag, Host or Prefix
    Granularly removes one or more files from Cloudflare's cache either by specifying the host, the associated Cache-Tag, or a Prefix.

    Flex purge with tags:
    ```
    {"tags": ["a-cache-tag", "another-cache-tag"]}
    ```
    Flex purge with hosts:
    ```
    {"hosts": ["www.example.com", "images.example.com"]}
    ```
    Flex purge with prefixes:
    ```
    {"prefixes": ["www.example.com/foo", "images.example.com/bar/baz"]}
    ```

    ### Availability and limits
    Please refer to [purge cache availability and limits documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/#availability-and-limits).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/purge_cache
operation_ids:
    - zone-purge
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Purge Cached Content

`POST /zones/{zone_id}/purge_cache`

Operation ID: `zone-purge`

### Purge All Cached Content
Removes ALL files from Cloudflare's cache. All tiers can purge everything.
```
{"purge_everything": true}
```

### Purge Cached Content by URL
Granularly removes one or more files from Cloudflare's cache by specifying URLs. All tiers can purge by URL.

To purge files with custom cache keys, include the headers used to compute the cache key as in the example. If you have a device type or geo in your cache key, you will need to include the CF-Device-Type or CF-IPCountry headers. If you have lang in your cache key, you will need to include the Accept-Language header.

**NB:** When including the Origin header, be sure to include the **scheme** and **hostname**. The port number can be omitted if it is the default port (80 for http, 443 for https), but must be included otherwise.

Single file purge example with files:
```
{"files": ["http://www.example.com/css/styles.css", "http://www.example.com/js/index.js"]}
```
Single file purge example with url and header pairs:
```
{"files": [{"url": "http://www.example.com/cat_picture.jpg", "headers": {"CF-IPCountry": "US", "CF-Device-Type": "desktop", "Accept-Language": "zh-CN"}}, {"url": "http://www.example.com/dog_picture.jpg", "headers": {"CF-IPCountry": "EU", "CF-Device-Type": "mobile", "Accept-Language": "en-US"}}]}
```

### Purge Cached Content by Tag, Host or Prefix
Granularly removes one or more files from Cloudflare's cache either by specifying the host, the associated Cache-Tag, or a Prefix.

Flex purge with tags:
```
{"tags": ["a-cache-tag", "another-cache-tag"]}
```
Flex purge with hosts:
```
{"hosts": ["www.example.com", "images.example.com"]}
```
Flex purge with prefixes:
```
{"prefixes": ["www.example.com/foo", "images.example.com/bar/baz"]}
```

### Availability and limits
Please refer to [purge cache availability and limits documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/#availability-and-limits).

## Definition

```yaml
{"operationId": "zone-purge", "summary": "Purge Cached Content", "description": "### Purge All Cached Content\nRemoves ALL files from Cloudflare's cache. All tiers can purge everything.\n```\n{\"purge_everything\": true}\n```\n\n### Purge Cached Content by URL\nGranularly removes one or more files from Cloudflare's cache by specifying URLs. All tiers can purge by URL.\n\nTo purge files with custom cache keys, include the headers used to compute the cache key as in the example. If you have a device type or geo in your cache key, you will need to include the CF-Device-Type or CF-IPCountry headers. If you have lang in your cache key, you will need to include the Accept-Language header.\n\n**NB:** When including the Origin header, be sure to include the **scheme** and **hostname**. The port number can be omitted if it is the default port (80 for http, 443 for https), but must be included otherwise.\n\nSingle file purge example with files:\n```\n{\"files\": [\"http://www.example.com/css/styles.css\", \"http://www.example.com/js/index.js\"]}\n```\nSingle file purge example with url and header pairs:\n```\n{\"files\": [{\"url\": \"http://www.example.com/cat_picture.jpg\", \"headers\": {\"CF-IPCountry\": \"US\", \"CF-Device-Type\": \"desktop\", \"Accept-Language\": \"zh-CN\"}}, {\"url\": \"http://www.example.com/dog_picture.jpg\", \"headers\": {\"CF-IPCountry\": \"EU\", \"CF-Device-Type\": \"mobile\", \"Accept-Language\": \"en-US\"}}]}\n```\n\n### Purge Cached Content by Tag, Host or Prefix\nGranularly removes one or more files from Cloudflare's cache either by specifying the host, the associated Cache-Tag, or a Prefix.\n\nFlex purge with tags:\n```\n{\"tags\": [\"a-cache-tag\", \"another-cache-tag\"]}\n```\nFlex purge with hosts:\n```\n{\"hosts\": [\"www.example.com\", \"images.example.com\"]}\n```\nFlex purge with prefixes:\n```\n{\"prefixes\": [\"www.example.com/foo\", \"images.example.com/bar/baz\"]}\n```\n\n### Availability and limits\nPlease refer to [purge cache availability and limits documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/#availability-and-limits).\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-purge_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"Flex Purge with Hosts": {"summary": "Flex purge example with hosts list", "value": {"hosts": ["www.example.com", "images.example.com"]}}, "Flex Purge with Prefixes": {"summary": "Flex purge example with prefixes list", "value": {"prefixes": ["www.example.com/foo", "images.example.com/bar/baz"]}}, "Flex Purge with Tags": {"summary": "Flex purge example with tags list", "value": {"tags": ["a-cache-tag", "another-cache-tag"]}}, "Purge Everything": {"summary": "Purge everything example", "value": {"purge_everything": true}}, "Single File Purge": {"summary": "Single file purge example with files list", "value": {"files": ["http://www.example.com/css/styles.css", "http://www.example.com/js/index.js"]}}, "Single File Purge with UrlAndHeaders": {"summary": "Single file purge example with url and headers list", "value": {"files": [{"headers": {"Accept-Language": "zh-CN", "CF-Device-Type": "desktop", "CF-IPCountry": "US"}, "url": "http://www.example.com/cat_picture.jpg"}, {"headers": {"Accept-Language": "en-US", "CF-Device-Type": "mobile", "CF-IPCountry": "EU"}, "url": "http://www.example.com/dog_picture.jpg"}]}}}, "schema": {"anyOf": [{"$ref": "#/components/schemas/cache-purge_FlexPurgeByTags"}, {"$ref": "#/components/schemas/cache-purge_FlexPurgeByHostnames"}, {"$ref": "#/components/schemas/cache-purge_FlexPurgeByPrefixes"}, {"$ref": "#/components/schemas/cache-purge_Everything"}, {"$ref": "#/components/schemas/cache-purge_SingleFile"}, {"$ref": "#/components/schemas/cache-purge_SingleFileWithUrlAndHeaders"}]}}}}, "responses": {"200": {"description": "Request to purge cached content successful.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache-purge_api-response-single-id"}}}}, "4XX": {"description": "Request to purge cached content failed.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cache-purge_api-response-single-id"}, {"$ref": "#/components/schemas/cache-purge_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone"], "x-api-token-group": ["Cache Purge"], "x-cfPermissionsRequired": {"enum": ["#cache_purge:edit"]}}
```
