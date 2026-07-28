---
title: Get Update Status
page_id: operation-get-accounts-account-id-registrar-registrations-domain-name-update-statu-45cbbc14
path: operations/registrar-registration
description: |-
    Returns the current status of a domain update workflow.

    Use this endpoint to poll for completion when the PATCH response
    returned `202 Accepted`. The URL is provided in the `links.self`
    field of the workflow status response.

    Poll this endpoint until the workflow reaches a terminal state or a
    state that requires user attention.

    Use increasing backoff between polls. When the workflow remains blocked
    on a third party, use a longer polling interval and do not poll indefinitely.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/registrar/registrations/{domain_name}/update-status
operation_ids:
    - registrar-domain-registration-get-update-status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Update Status

`GET /accounts/{account_id}/registrar/registrations/{domain_name}/update-status`

Operation ID: `registrar-domain-registration-get-update-status`

Returns the current status of a domain update workflow.

Use this endpoint to poll for completion when the PATCH response
returned `202 Accepted`. The URL is provided in the `links.self`
field of the workflow status response.

Poll this endpoint until the workflow reaches a terminal state or a
state that requires user attention.

Use increasing backoff between polls. When the workflow remains blocked
on a third party, use a longer polling interval and do not poll indefinitely.

## Definition

```yaml
{"operationId": "registrar-domain-registration-get-update-status", "summary": "Get Update Status", "description": "Returns the current status of a domain update workflow.\n\nUse this endpoint to poll for completion when the PATCH response\nreturned `202 Accepted`. The URL is provided in the `links.self`\nfield of the workflow status response.\n\nPoll this endpoint until the workflow reaches a terminal state or a\nstate that requires user attention.\n\nUse increasing backoff between polls. When the workflow remains blocked\non a third party, use a longer polling interval and do not poll indefinitely.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api_identifier"}}, {"name": "domain_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api_domain_name"}}], "responses": {"200": {"description": "Update workflow status.", "content": {"application/json": {"examples": {"failed": {"summary": "Update failed", "value": {"errors": [], "messages": [], "result": {"completed": true, "context": {"domain_name": "example.com"}, "created_at": "2025-10-27T10:00:00Z", "error": {"code": "registry_rejected", "message": "Registry rejected the update request."}, "links": {"resource": "/accounts/abc/registrar/registrations/example.com", "self": "/accounts/abc/registrar/registrations/example.com/update-status"}, "state": "failed", "updated_at": "2025-10-27T10:00:45Z"}, "success": true}}, "in_progress": {"summary": "Update in progress", "value": {"errors": [], "messages": [], "result": {"completed": false, "context": {"domain_name": "example.com", "fields_updated": ["auto_renew"], "pending_registry_confirmation": true}, "created_at": "2025-10-27T10:00:00Z", "links": {"resource": "/accounts/abc/registrar/registrations/example.com", "self": "/accounts/abc/registrar/registrations/example.com/update-status"}, "state": "in_progress", "updated_at": "2025-10-27T10:00:05Z"}, "success": true}}, "succeeded": {"summary": "Update completed successfully", "value": {"errors": [], "messages": [], "result": {"completed": true, "context": {"domain_name": "example.com", "fields_updated": ["auto_renew"]}, "created_at": "2025-10-27T10:00:00Z", "links": {"resource": "/accounts/abc/registrar/registrations/example.com", "self": "/accounts/abc/registrar/registrations/example.com/update-status"}, "state": "succeeded", "updated_at": "2025-10-27T10:00:30Z"}, "success": true}}}, "schema": {"$ref": "#/components/schemas/registrar-api_workflow-status-response-single"}}}}, "4XX": {"description": "Get update status failure.", "content": {"application/json": {"examples": {"workflow_not_found": {"summary": "No workflow found", "value": {"errors": [{"code": 10000, "message": "No workflow found for example.com"}], "messages": [], "result": null, "success": false}}}, "schema": {"$ref": "#/components/schemas/registrar-api_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Registrar Registration"]}
```
