---
title: Create a sourcing kit source
page_id: operation-post-accounts-account-id-images-v2-sourcingkit-sources-4e4de1da
path: operations/cloudflare-images-sourcing-kit
description: |-
    Create a new migration source by providing storage credentials. The service
    will verify connectivity to the bucket before accepting the source.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/images/v2/sourcingkit/sources
operation_ids:
    - cloudflare-images-sourcingkit-create-source
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a sourcing kit source

`POST /accounts/{account_id}/images/v2/sourcingkit/sources`

Operation ID: `cloudflare-images-sourcingkit-create-source`

Create a new migration source by providing storage credentials. The service
will verify connectivity to the bucket before accepting the source.

## Definition

```yaml
{"operationId": "cloudflare-images-sourcingkit-create-source", "summary": "Create a sourcing kit source", "description": "Create a new migration source by providing storage credentials. The service\nwill verify connectivity to the bucket before accepting the source.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_sourcingkit_source_create_request"}}}}, "responses": {"200": {"description": "Create sourcing kit source response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_sourcingkit_source_create_response"}}}}, "4XX": {"description": "Create sourcing kit source response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_sourcingkit_source_create_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Sourcing Kit"], "x-api-token-group": ["Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.sourcing-kit.sources", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
