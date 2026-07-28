---
title: List sourcing kit sources
page_id: operation-get-accounts-account-id-images-v2-sourcingkit-sources-ca8ba6c2
path: operations/cloudflare-images-sourcing-kit
description: List all configured migration sources for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/images/v2/sourcingkit/sources
operation_ids:
    - cloudflare-images-sourcingkit-list-sources
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List sourcing kit sources

`GET /accounts/{account_id}/images/v2/sourcingkit/sources`

Operation ID: `cloudflare-images-sourcingkit-list-sources`

List all configured migration sources for the account.

## Definition

```yaml
{"operationId": "cloudflare-images-sourcingkit-list-sources", "summary": "List sourcing kit sources", "description": "List all configured migration sources for the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}, {"name": "offset", "in": "query", "schema": {"description": "Number of items to skip before returning results.", "type": "integer", "default": 0, "minimum": 0}}, {"name": "limit", "in": "query", "schema": {"description": "Maximum number of items to return.", "type": "integer", "default": 25, "maximum": 100, "minimum": 1}}, {"name": "name", "in": "query", "schema": {"description": "Filter sources by name (partial match).", "type": "string"}}], "responses": {"200": {"description": "List sourcing kit sources response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_sourcingkit_source_list_response"}}}}, "4XX": {"description": "List sourcing kit sources response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_sourcingkit_source_list_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Sourcing Kit"], "x-api-token-group": ["Images Read", "Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.sourcing-kit.sources", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
