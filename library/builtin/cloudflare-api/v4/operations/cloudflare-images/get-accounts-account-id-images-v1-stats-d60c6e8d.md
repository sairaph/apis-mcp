---
title: Images usage statistics
page_id: operation-get-accounts-account-id-images-v1-stats-fd985ba1
path: operations/cloudflare-images
description: Fetch image statistics details for Cloudflare Images. The returned statistics detail storage usage, including the current image count vs this account's allowance.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/images/v1/stats
operation_ids:
    - cloudflare-images-images-usage-statistics
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Images usage statistics

`GET /accounts/{account_id}/images/v1/stats`

Operation ID: `cloudflare-images-images-usage-statistics`

Fetch image statistics details for Cloudflare Images. The returned statistics detail storage usage, including the current image count vs this account's allowance.

## Definition

```yaml
{"operationId": "cloudflare-images-images-usage-statistics", "summary": "Images usage statistics", "description": "Fetch image statistics details for Cloudflare Images. The returned statistics detail storage usage, including the current image count vs this account's allowance.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "responses": {"200": {"description": "Images usage statistics response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_images_stats_response"}}}}, "4XX": {"description": "Images usage statistics response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_images_stats_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images"], "x-api-token-group": ["Images Read", "Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images", "x-fern-sdk-method-name": "stats", "x-forge-hidden": false}
```
