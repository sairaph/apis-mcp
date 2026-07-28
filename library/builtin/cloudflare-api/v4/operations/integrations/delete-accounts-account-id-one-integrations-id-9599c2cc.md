---
title: Delete integration
page_id: operation-delete-accounts-account-id-one-integrations-id-d10cb0b8
path: operations/integrations
description: Delete an integration by soft-deleting it.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/one/integrations/{id}
operation_ids:
    - delete_integration_v2
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete integration

`DELETE /accounts/{account_id}/one/integrations/{id}`

Operation ID: `delete_integration_v2`

Delete an integration by soft-deleting it.

## Definition

```yaml
{"operationId": "delete_integration_v2", "summary": "Delete integration", "description": "Delete an integration by soft-deleting it.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account identifier.", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}, {"name": "id", "in": "path", "description": "Integration ID.", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Integration deleted successfully."}, "400": {"description": "Invalid request."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Integrations"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero_trust.casb.integrations", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true, "x-stability": "beta"}
```
