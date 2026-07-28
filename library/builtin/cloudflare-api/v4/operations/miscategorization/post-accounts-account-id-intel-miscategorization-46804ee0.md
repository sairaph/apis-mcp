---
title: Create Miscategorization
page_id: operation-post-accounts-account-id-intel-miscategorization-8a859e2d
path: operations/miscategorization
description: |-
    Allows you to submit requests to change a domain’s category.

    Requests that include category `169` (New Domains) or category `177` (Newly Seen)
    in any of `content_adds`, `content_removes`, `security_adds`, or `security_removes`
    will be rejected with a `400 Bad Request`. These categories are automatically
    managed and fall off 30 days after they are applied.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/intel/miscategorization
operation_ids:
    - miscategorization-create-miscategorization
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Miscategorization

`POST /accounts/{account_id}/intel/miscategorization`

Operation ID: `miscategorization-create-miscategorization`

Allows you to submit requests to change a domain’s category.

Requests that include category `169` (New Domains) or category `177` (Newly Seen)
in any of `content_adds`, `content_removes`, `security_adds`, or `security_removes`
will be rejected with a `400 Bad Request`. These categories are automatically
managed and fall off 30 days after they are applied.

## Definition

```yaml
{"operationId": "miscategorization-create-miscategorization", "summary": "Create Miscategorization", "description": "Allows you to submit requests to change a domain’s category.\n\nRequests that include category `169` (New Domains) or category `177` (Newly Seen)\nin any of `content_adds`, `content_removes`, `security_adds`, or `security_removes`\nwill be rejected with a `400 Bad Request`. These categories are automatically\nmanaged and fall off 30 days after they are applied.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/intel_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel_miscategorization"}}}}, "responses": {"200": {"description": "Create Miscategorization response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel_api-response-single"}}}}, "4XX": {"description": "Create Miscategorization response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/intel_api-response-single"}, {"$ref": "#/components/schemas/intel_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Miscategorization"]}
```
