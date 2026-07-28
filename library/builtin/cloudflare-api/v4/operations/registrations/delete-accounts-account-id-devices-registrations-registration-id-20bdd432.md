---
title: Delete registration
page_id: operation-delete-accounts-account-id-devices-registrations-registration-id-038323ab
path: operations/registrations
description: Deletes a WARP registration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/devices/registrations/{registration_id}
operation_ids:
    - delete-registration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete registration

`DELETE /accounts/{account_id}/devices/registrations/{registration_id}`

Operation ID: `delete-registration`

Deletes a WARP registration.

## Definition

```yaml
{"operationId": "delete-registration", "summary": "Delete registration", "description": "Deletes a WARP registration.", "parameters": [{"name": "registration_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Registration deleted response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"$ref": "#/components/schemas/teams-devices_empty_body"}, "success": {"description": "Whether the API call was successful.", "type": "boolean"}}, "required": ["success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Registrations"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.registrations", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
