---
title: Precheck source connectivity
page_id: operation-post-accounts-account-id-images-v2-sourcingkit-sources-connectivity-prec-dc86b6c7
path: operations/cloudflare-images-sourcing-kit
description: |-
    Verify connectivity to a storage bucket before creating a source. Returns
    connectivity status without persisting any state.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/images/v2/sourcingkit/sources/connectivity-precheck
operation_ids:
    - cloudflare-images-sourcingkit-precheck-source-connectivity
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Precheck source connectivity

`POST /accounts/{account_id}/images/v2/sourcingkit/sources/connectivity-precheck`

Operation ID: `cloudflare-images-sourcingkit-precheck-source-connectivity`

Verify connectivity to a storage bucket before creating a source. Returns
connectivity status without persisting any state.

## Definition

```yaml
{"operationId": "cloudflare-images-sourcingkit-precheck-source-connectivity", "summary": "Precheck source connectivity", "description": "Verify connectivity to a storage bucket before creating a source. Returns\nconnectivity status without persisting any state.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_sourcingkit_connectivity_precheck_request"}}}}, "responses": {"200": {"description": "Connectivity precheck response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_sourcingkit_connectivity_check_response"}}}}, "4XX": {"description": "Connectivity precheck response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_sourcingkit_connectivity_check_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Sourcing Kit"], "x-api-token-group": ["Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.sourcing-kit.sources", "x-fern-sdk-method-name": "pre-check", "x-forge-hidden": true}
```
