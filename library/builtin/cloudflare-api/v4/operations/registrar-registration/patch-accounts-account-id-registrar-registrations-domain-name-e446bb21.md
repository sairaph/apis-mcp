---
title: Update Registration
page_id: operation-patch-accounts-account-id-registrar-registrations-domain-name-acdbd19a
path: operations/registrar-registration
description: |-
    Updates an existing domain registration.

    By default, the server holds the connection for a bounded, server-defined
    amount of time while the update completes. Most updates finish within this
    window and return `200 OK` with a completed workflow status.

    If the update is still processing after this synchronous wait window, the
    server returns `202 Accepted`. Poll the URL in `links.self` to track progress.

    To skip the wait and receive an immediate `202`, send `Prefer: respond-async`.

    This endpoint currently supports updating `auto_renew` only.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/registrar/registrations/{domain_name}
operation_ids:
    - registrar-domain-registration-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Registration

`PATCH /accounts/{account_id}/registrar/registrations/{domain_name}`

Operation ID: `registrar-domain-registration-update`

Updates an existing domain registration.

By default, the server holds the connection for a bounded, server-defined
amount of time while the update completes. Most updates finish within this
window and return `200 OK` with a completed workflow status.

If the update is still processing after this synchronous wait window, the
server returns `202 Accepted`. Poll the URL in `links.self` to track progress.

To skip the wait and receive an immediate `202`, send `Prefer: respond-async`.

This endpoint currently supports updating `auto_renew` only.

## Definition

```yaml
{"operationId": "registrar-domain-registration-update", "summary": "Update Registration", "description": "Updates an existing domain registration.\n\nBy default, the server holds the connection for a bounded, server-defined\namount of time while the update completes. Most updates finish within this\nwindow and return `200 OK` with a completed workflow status.\n\nIf the update is still processing after this synchronous wait window, the\nserver returns `202 Accepted`. Poll the URL in `links.self` to track progress.\n\nTo skip the wait and receive an immediate `202`, send `Prefer: respond-async`.\n\nThis endpoint currently supports updating `auto_renew` only.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api_identifier"}}, {"name": "domain_name", "in": "path", "description": "Domain name to update.", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api_domain_name"}}, {"name": "Prefer", "in": "header", "description": "Set to `respond-async` to receive an immediate `202 Accepted` without\nwaiting for the operation to complete (RFC 7240).\n", "schema": {"type": "string", "enum": ["respond-async"]}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"disable_auto_renew": {"summary": "Disable auto-renewal", "value": {"auto_renew": false}}, "enable_auto_renew": {"summary": "Enable auto-renewal", "value": {"auto_renew": true}}}, "schema": {"$ref": "#/components/schemas/registrar-api_registration_update_request"}}}}, "responses": {"200": {"description": "Update completed successfully within the synchronous wait window.\nThe workflow status will have `state: succeeded` and `completed: true`.\n", "headers": {"Preference-Applied": {"description": "Echoed when the server honored a `Prefer` header.", "schema": {"type": "string"}}}, "content": {"application/json": {"examples": {"completed": {"summary": "Update completed within timeout (most common)", "value": {"errors": [], "messages": [], "result": {"completed": true, "context": {"domain_name": "example.com", "registration": {"auto_renew": true, "created_at": "2025-01-15T10:00:00Z", "domain_name": "example.com", "expires_at": "2026-01-15T10:00:00Z", "locked": true, "privacy_mode": "redaction", "status": "active"}}, "created_at": "2025-10-27T10:00:00Z", "links": {"resource": "/accounts/abc/registrar/registrations/example.com", "self": "/accounts/abc/registrar/registrations/example.com/update-status"}, "state": "succeeded", "updated_at": "2025-10-27T10:00:02Z"}, "success": true}}}, "schema": {"$ref": "#/components/schemas/registrar-api_workflow-status-response-single"}}}}, "202": {"description": "Update is still processing. This occurs when the operation did not\ncomplete within the synchronous wait window, or when `Prefer: respond-async`\nwas sent. Poll the URL in `links.self` to track progress.\n", "headers": {"Location": {"description": "URL of the workflow status resource (same as `links.self`).", "schema": {"type": "string"}}, "Preference-Applied": {"description": "Set to `respond-async` when the server honored the preference.", "schema": {"type": "string"}}}, "content": {"application/json": {"examples": {"still_processing": {"summary": "Update still in progress after timeout", "value": {"errors": [], "messages": [], "result": {"completed": false, "context": {"domain_name": "example.com"}, "created_at": "2025-10-27T10:00:00Z", "links": {"resource": "/accounts/abc/registrar/registrations/example.com", "self": "/accounts/abc/registrar/registrations/example.com/update-status"}, "state": "in_progress", "updated_at": "2025-10-27T10:00:10Z"}, "success": true}}}, "schema": {"$ref": "#/components/schemas/registrar-api_workflow-status-response-single"}}}}, "4XX": {"description": "Update registration failure.", "content": {"application/json": {"examples": {"domain_not_found": {"summary": "Domain not found", "value": {"errors": [{"code": 10000, "message": "Domain not found"}], "messages": [], "result": null, "success": false}}, "invalid_auto_renew": {"summary": "Invalid auto_renew value", "value": {"errors": [{"code": 10000, "message": "Must be a boolean if present", "source": {"pointer": "/auto_renew"}}], "messages": [], "result": null, "success": false}}}, "schema": {"$ref": "#/components/schemas/registrar-api_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Registrar Registration"]}
```
