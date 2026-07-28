---
title: Delete Subdomain
page_id: operation-delete-accounts-account-id-workers-subdomain-b201ecdc
path: operations/worker-subdomain
description: Deletes a Workers subdomain for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/workers/subdomain
operation_ids:
    - worker-subdomain-delete-subdomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Subdomain

`DELETE /accounts/{account_id}/workers/subdomain`

Operation ID: `worker-subdomain-delete-subdomain`

Deletes a Workers subdomain for an account.

## Definition

```yaml
{"operationId": "worker-subdomain-delete-subdomain", "summary": "Delete Subdomain", "description": "Deletes a Workers subdomain for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "responses": {"204": {"description": "Subdomain deleted successfully."}, "4XX": {"description": "Delete Subdomain response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Subdomain"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.subdomains", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
