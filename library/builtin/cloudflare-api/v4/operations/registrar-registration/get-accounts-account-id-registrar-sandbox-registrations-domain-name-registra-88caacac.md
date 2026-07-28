---
title: Get Registration Status
page_id: operation-get-accounts-account-id-registrar-sandbox-registrations-domain-name-regi-ca6d284b
path: operations/registrar-registration
description: |-
    Returns the current status of a domain registration workflow.

    Use this endpoint to poll for completion when the POST response
    returned `202 Accepted`. The URL is provided in the `links.self`
    field of the workflow status response.

    Poll this endpoint until the workflow reaches a terminal state or a
    state that requires user attention.

    **Terminal states:** `succeeded` and `failed` are terminal and always
    have `completed: true`.

    **Non-terminal states:**
    - `action_required` has `completed: false` and will not resolve on its
      own. The workflow is paused pending user intervention.
    - `blocked` has `completed: false` and indicates the workflow is waiting
      on a third party such as the extension registry or losing registrar.
      Continue polling while informing the user of the delay.

    Use increasing backoff between polls. When `state: blocked`, use a
    longer polling interval and do not poll indefinitely.

    A naive polling loop that only checks `completed` can run indefinitely
    when `state: action_required`. Break explicitly on `action_required`:

    ```js
    let status;
    do {
      await new Promise(r => setTimeout(r, 2000));
      status = await cloudflare.request({
        method: 'GET',
        path: reg.result.links.self,
      });
    } while (
      !status.result.completed &&
      status.result.state !== 'action_required'
    );

    if (status.result.state === 'action_required') {
      // Surface context.action and context.confirmation_sent_to to the user.
      // Do not re-submit the registration request.
    }
    ```
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/registrar-sandbox/registrations/{domain_name}/registration-status
operation_ids:
    - sandbox-registrar-domain-registration-get-status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Registration Status

`GET /accounts/{account_id}/registrar-sandbox/registrations/{domain_name}/registration-status`

Operation ID: `sandbox-registrar-domain-registration-get-status`

Returns the current status of a domain registration workflow.

Use this endpoint to poll for completion when the POST response
returned `202 Accepted`. The URL is provided in the `links.self`
field of the workflow status response.

Poll this endpoint until the workflow reaches a terminal state or a
state that requires user attention.

**Terminal states:** `succeeded` and `failed` are terminal and always
have `completed: true`.

**Non-terminal states:**
- `action_required` has `completed: false` and will not resolve on its
  own. The workflow is paused pending user intervention.
- `blocked` has `completed: false` and indicates the workflow is waiting
  on a third party such as the extension registry or losing registrar.
  Continue polling while informing the user of the delay.

Use increasing backoff between polls. When `state: blocked`, use a
longer polling interval and do not poll indefinitely.

A naive polling loop that only checks `completed` can run indefinitely
when `state: action_required`. Break explicitly on `action_required`:

```js
let status;
do {
  await new Promise(r => setTimeout(r, 2000));
  status = await cloudflare.request({
    method: 'GET',
    path: reg.result.links.self,
  });
} while (
  !status.result.completed &&
  status.result.state !== 'action_required'
);

if (status.result.state === 'action_required') {
  // Surface context.action and context.confirmation_sent_to to the user.
  // Do not re-submit the registration request.
}
```

## Definition

```yaml
{"operationId": "sandbox-registrar-domain-registration-get-status", "summary": "Get Registration Status", "description": "Returns the current status of a domain registration workflow.\n\nUse this endpoint to poll for completion when the POST response\nreturned `202 Accepted`. The URL is provided in the `links.self`\nfield of the workflow status response.\n\nPoll this endpoint until the workflow reaches a terminal state or a\nstate that requires user attention.\n\n**Terminal states:** `succeeded` and `failed` are terminal and always\nhave `completed: true`.\n\n**Non-terminal states:**\n- `action_required` has `completed: false` and will not resolve on its\n  own. The workflow is paused pending user intervention.\n- `blocked` has `completed: false` and indicates the workflow is waiting\n  on a third party such as the extension registry or losing registrar.\n  Continue polling while informing the user of the delay.\n\nUse increasing backoff between polls. When `state: blocked`, use a\nlonger polling interval and do not poll indefinitely.\n\nA naive polling loop that only checks `completed` can run indefinitely\nwhen `state: action_required`. Break explicitly on `action_required`:\n\n```js\nlet status;\ndo {\n  await new Promise(r => setTimeout(r, 2000));\n  status = await cloudflare.request({\n    method: 'GET',\n    path: reg.result.links.self,\n  });\n} while (\n  !status.result.completed &&\n  status.result.state !== 'action_required'\n);\n\nif (status.result.state === 'action_required') {\n  // Surface context.action and context.confirmation_sent_to to the user.\n  // Do not re-submit the registration request.\n}\n```\n", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID. Required for all Registrar API operations.", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api-sandbox_identifier"}}, {"name": "domain_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api-sandbox_domain_name"}}], "responses": {"200": {"description": "Registration workflow status.", "content": {"application/json": {"examples": {"action_required": {"description": "Non-terminal but not transient. The workflow is paused and\nrequires user action to proceed. An automated polling loop\nmust break on this state — it will not resolve on its own.\nSurface context.action to the user.\n", "summary": "User action required", "value": {"errors": [], "messages": [], "result": {"completed": false, "context": {"action": "registrant_email_confirmation_pending", "confirmation_sent_to": "a***@example.com", "domain_name": "example.com"}, "created_at": "2025-10-27T10:00:00Z", "links": {"resource": "/accounts/abc/registrar/registrations/example.com", "self": "/accounts/abc/registrar/registrations/example.com/registration-status"}, "state": "action_required", "updated_at": "2025-10-27T10:01:00Z"}, "success": true}}, "blocked": {"description": "Non-terminal. The workflow cannot make progress due to a\nthird party (e.g., the extension's registry or a losing\nregistrar). Unlike action_required, no user action will help.\nContinue polling — the block may resolve when the third party\nresponds.\n", "summary": "Blocked by third party", "value": {"errors": [], "messages": [], "result": {"completed": false, "context": {"blocked_by": "registry", "detail": "Awaiting registry confirmation. This may take up to 24 hours.", "domain_name": "example.com"}, "created_at": "2025-10-27T10:00:00Z", "links": {"resource": "/accounts/abc/registrar/registrations/example.com", "self": "/accounts/abc/registrar/registrations/example.com/registration-status"}, "state": "blocked", "updated_at": "2025-10-27T10:05:00Z"}, "success": true}}, "failed": {"description": "Terminal state. The registration could not be completed.\nSee error.code and error.message for the reason. Do not\nauto-retry without user review.\n", "summary": "Registration failed", "value": {"errors": [], "messages": [], "result": {"completed": true, "context": {"domain_name": "example.com"}, "created_at": "2025-10-27T10:00:00Z", "error": {"code": "registry_rejected", "message": "Registry rejected the registration request."}, "links": {"resource": "/accounts/abc/registrar/registrations/example.com", "self": "/accounts/abc/registrar/registrations/example.com/registration-status"}, "state": "failed", "updated_at": "2025-10-27T10:00:08Z"}, "success": true}}, "in_progress": {"description": "The registration is still being processed. Continue polling\nthis endpoint. In synchronous mode, the server only waits\nfor a bounded amount of time before returning `202` — this\nstatus is what you see when polling after that timeout. Registrations\nwill not remain in this state indefinitely; the workflow has\nan internal deadline after which it transitions to succeeded,\nfailed, or blocked.\n", "summary": "Registration in progress", "value": {"errors": [], "messages": [], "result": {"completed": false, "context": {"domain_name": "example.com"}, "created_at": "2025-10-27T10:00:00Z", "links": {"resource": "/accounts/abc/registrar/registrations/example.com", "self": "/accounts/abc/registrar/registrations/example.com/registration-status"}, "state": "in_progress", "updated_at": "2025-10-27T10:00:12Z"}, "success": true}}, "succeeded": {"description": "Terminal state. The domain has been registered. The full\nregistration resource is included in context.registration —\nno additional GET request is needed.\n", "summary": "Registration completed successfully", "value": {"errors": [], "messages": [], "result": {"completed": true, "context": {"domain_name": "example.com", "registration": {"auto_renew": true, "created_at": "2025-10-27T10:00:00Z", "domain_name": "example.com", "expires_at": "2026-10-27T10:00:00Z", "locked": true, "privacy_mode": "redaction", "status": "active"}}, "created_at": "2025-10-27T10:00:00Z", "links": {"resource": "/accounts/abc/registrar/registrations/example.com", "self": "/accounts/abc/registrar/registrations/example.com/registration-status"}, "state": "succeeded", "updated_at": "2025-10-27T10:00:03Z"}, "success": true}}}, "schema": {"$ref": "#/components/schemas/registrar-api-sandbox_workflow-status-response-single"}}}}, "4XX": {"description": "Get status failure.", "content": {"application/json": {"examples": {"workflow_not_found": {"summary": "No workflow found", "value": {"errors": [{"code": 10000, "message": "No workflow found for example.com"}], "messages": [], "result": null, "success": false}}}, "schema": {"$ref": "#/components/schemas/registrar-api-sandbox_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Registrar Registration"]}
```
