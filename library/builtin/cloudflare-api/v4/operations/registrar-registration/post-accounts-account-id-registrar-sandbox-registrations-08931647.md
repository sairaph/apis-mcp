---
title: Create Registration
page_id: operation-post-accounts-account-id-registrar-sandbox-registrations-90aae245
path: operations/registrar-registration
description: |-
    Starts a domain registration workflow.

    ### Prerequisites
    - The account must not already be at the maximum supported domain limit.
        A single account may own up to 100 domains in total across registrations
        created through either the dashboard or this API.
    - The domain must be on a supported extension for programmatic registration.
    - Use `POST /domain-check` immediately before calling this endpoint to confirm
        real-time availability and pricing.

    ### Defaults
    - `years`: defaults to the extension's minimum registration period (1 year for
        most extensions, but varies — for example, `.ai` (if supported) requires a minimum of 2 years).
    - `auto_renew`: defaults to `false`. Setting it to `true` is an explicit
        opt-in authorizing Cloudflare to charge the account's default payment
        method up to 30 days before domain expiry to renew the registration.
        Renewal pricing may change over time based on registry pricing.
    - `privacy_mode`: defaults to `redaction`.

    ### Premium domains
    Premium domain registration is not currently supported by this API.
    If `POST /domain-check` returns `tier: premium`, do not call this
    endpoint for that domain.

    ### Response behavior
    By default, the server holds the connection for a bounded, server-defined
    amount of time while the registration completes. Most registrations finish
    within this window and return `201 Created` with a completed workflow status.

    If the registration is still processing after this synchronous wait window,
    the server returns `202 Accepted`. Poll the URL in `links.self` to track progress.

    To skip the wait and receive an immediate `202`, send `Prefer: respond-async`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/registrar-sandbox/registrations
operation_ids:
    - sandbox-registrar-domain-registration-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Registration

`POST /accounts/{account_id}/registrar-sandbox/registrations`

Operation ID: `sandbox-registrar-domain-registration-create`

Starts a domain registration workflow.

### Prerequisites
- The account must not already be at the maximum supported domain limit.
    A single account may own up to 100 domains in total across registrations
    created through either the dashboard or this API.
- The domain must be on a supported extension for programmatic registration.
- Use `POST /domain-check` immediately before calling this endpoint to confirm
    real-time availability and pricing.

### Defaults
- `years`: defaults to the extension's minimum registration period (1 year for
    most extensions, but varies — for example, `.ai` (if supported) requires a minimum of 2 years).
- `auto_renew`: defaults to `false`. Setting it to `true` is an explicit
    opt-in authorizing Cloudflare to charge the account's default payment
    method up to 30 days before domain expiry to renew the registration.
    Renewal pricing may change over time based on registry pricing.
- `privacy_mode`: defaults to `redaction`.

### Premium domains
Premium domain registration is not currently supported by this API.
If `POST /domain-check` returns `tier: premium`, do not call this
endpoint for that domain.

### Response behavior
By default, the server holds the connection for a bounded, server-defined
amount of time while the registration completes. Most registrations finish
within this window and return `201 Created` with a completed workflow status.

If the registration is still processing after this synchronous wait window,
the server returns `202 Accepted`. Poll the URL in `links.self` to track progress.

To skip the wait and receive an immediate `202`, send `Prefer: respond-async`.

## Definition

```yaml
{"operationId": "sandbox-registrar-domain-registration-create", "summary": "Create Registration", "description": "Starts a domain registration workflow.\n\n### Prerequisites\n- The account must not already be at the maximum supported domain limit.\n    A single account may own up to 100 domains in total across registrations\n    created through either the dashboard or this API.\n- The domain must be on a supported extension for programmatic registration.\n- Use `POST /domain-check` immediately before calling this endpoint to confirm\n    real-time availability and pricing.\n\n### Defaults\n- `years`: defaults to the extension's minimum registration period (1 year for\n    most extensions, but varies — for example, `.ai` (if supported) requires a minimum of 2 years).\n- `auto_renew`: defaults to `false`. Setting it to `true` is an explicit\n    opt-in authorizing Cloudflare to charge the account's default payment\n    method up to 30 days before domain expiry to renew the registration.\n    Renewal pricing may change over time based on registry pricing.\n- `privacy_mode`: defaults to `redaction`.\n\n### Premium domains\nPremium domain registration is not currently supported by this API.\nIf `POST /domain-check` returns `tier: premium`, do not call this\nendpoint for that domain.\n\n### Response behavior\nBy default, the server holds the connection for a bounded, server-defined\namount of time while the registration completes. Most registrations finish\nwithin this window and return `201 Created` with a completed workflow status.\n\nIf the registration is still processing after this synchronous wait window,\nthe server returns `202 Accepted`. Poll the URL in `links.self` to track progress.\n\nTo skip the wait and receive an immediate `202`, send `Prefer: respond-async`.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID. Required for all Registrar API operations.", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api-sandbox_identifier"}}, {"name": "Prefer", "in": "header", "description": "Set to `respond-async` to receive an immediate `202 Accepted` without\nwaiting for the operation to complete (RFC 7240).\n\nThe header may be combined with other preferences using standard\ncomma-separated syntax.\n", "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"ca_with_contact_extensions": {"description": ".ca registrations require a registrant legal type. Provide the\nrequired value in `contact_extensions` as described by the live\n.ca registration schema.\n", "summary": "Register a .ca domain with legal type", "value": {"auto_renew": false, "contact_extensions": {"ca_legal_type": "CCT"}, "contacts": {"registrant": {"email": "ada@example.ca", "phone": "+1.6135550100", "postal_info": {"address": {"city": "Ottawa", "country_code": "CA", "postal_code": "K1A 0B1", "state": true, "street": "123 Queen St"}, "name": "Ada Lovelace", "organization": "Example Canada Inc"}}}, "domain_name": "my-brand-example.ca", "years": 1}}, "io_with_role_contacts": {"description": "Some extension schemas accept additional standard contact roles\nsuch as `technical`, `administrator`, and `billing`. When these\nroles are provided explicitly, Cloudflare preserves the supplied\nrole-specific contact data instead of deriving it from\n`contacts.registrant`.\n", "summary": "Register with explicit role contacts", "value": {"auto_renew": false, "contacts": {"administrator": {"email": "katherine@example.io", "phone": "+1.5555550102", "postal_info": {"address": {"city": "San Francisco", "country_code": "US", "postal_code": "94103", "state": "CA", "street": "789 Mission St"}, "name": "Katherine Johnson", "organization": "Example Admin Inc"}}, "billing": {"email": "dorothy@example.io", "phone": "+1.5555550103", "postal_info": {"address": {"city": "San Francisco", "country_code": "US", "postal_code": "94105", "state": "CA", "street": "101 Howard St"}, "name": "Dorothy Vaughan", "organization": "Example Billing Inc"}}, "registrant": {"email": "ada@example.io", "phone": "+1.5555555555", "postal_info": {"address": {"city": "Austin", "country_code": "US", "postal_code": "78701", "state": "TX", "street": "123 Main St"}, "name": "Ada Lovelace", "organization": "Example Inc"}}, "technical": {"email": "grace@example.io", "phone": "+1.5555550101", "postal_info": {"address": {"city": "San Francisco", "country_code": "US", "postal_code": "94105", "state": "CA", "street": "456 Market St"}, "name": "Grace Hopper", "organization": "Example Technical Inc"}}}, "domain_name": "my-brand-example.io", "years": 1}}, "minimal": {"description": "The simplest registration request. All defaults apply:\n- `years` defaults to the extension's minimum registration period (typically 1 year)\n- `auto_renew` defaults to `false`\n- `privacy_mode` defaults to `redaction`\n- Registrant contact falls back to the account's default address book entry\n", "summary": "Minimal registration — just a domain name", "value": {"domain_name": "my-new-startup.com"}}, "multi_year": {"description": "Register a domain for 3 years.", "summary": "Register for multiple years", "value": {"domain_name": "longterm-project.dev", "years": 3}}, "uk_with_contact_extensions": {"description": ".uk registrations can require registrant type information. For\norganization registrants, include the company number when the\nlive registration schema requires it.\n", "summary": "Register a .uk domain with registrant type", "value": {"auto_renew": false, "contact_extensions": {"company_number": "12345678", "registrant_type": "LTD"}, "contacts": {"registrant": {"email": "ada@example.co.uk", "phone": "+44.2071234567", "postal_info": {"address": {"city": "London", "country_code": "GB", "postal_code": "SW1A 1AA", "state": "London", "street": "1 Example Street"}, "name": "Ada Lovelace", "organization": "Example UK Ltd"}}}, "domain_name": "my-brand-example.uk", "years": 1}}, "us_with_contact_extensions": {"description": "Some extensions require registry-specific contact extension\nvalues in addition to the standard registrant contact. The\nextension discovery endpoint is authoritative for the required\nkeys and allowed values.\n", "summary": "Register a .us domain with contact extensions", "value": {"auto_renew": false, "contact_extensions": {"application_purpose": "P3", "nexus_category": "C11"}, "contacts": {"registrant": {"email": "ada@example.com", "phone": "+1.5555555555", "postal_info": {"address": {"city": "Austin", "country_code": "US", "postal_code": "78701", "state": "TX", "street": "123 Main St"}, "name": "Ada Lovelace", "organization": "Example Inc"}}}, "domain_name": "my-brand-example.us", "years": 1}}, "with_contact": {"description": "Provide registrant contact data inline if the account does not have\na default address book entry, or to override it for this registration.\n`postal_info.name` is the complete contact name in one field.\nSome registries require a complete personal name, including a\nfamily or last name where applicable, but this API does not\naccept separate first-name and last-name fields.\n", "summary": "Register with explicit registrant contact", "value": {"auto_renew": true, "contacts": {"registrant": {"email": "ada@example.com", "phone": "+1.5555555555", "postal_info": {"address": {"city": "Austin", "country_code": "US", "postal_code": "78701", "state": "TX", "street": "123 Main St"}, "name": "Ada Lovelace", "organization": "Example Inc"}}}, "domain_name": "example.com", "years": 1}}}, "schema": {"$ref": "#/components/schemas/registrar-api-sandbox_registration_create_request"}}}}, "responses": {"201": {"description": "Registration completed successfully within the synchronous wait window.\nThe workflow status will have `state: succeeded` and `completed: true`.\n", "headers": {"Preference-Applied": {"description": "Echoed when the server honored a `Prefer` header.", "schema": {"type": "string"}}}, "content": {"application/json": {"examples": {"completed": {"summary": "Registration completed within timeout (most common)", "value": {"errors": [], "messages": [], "result": {"completed": true, "context": {"domain_name": "example.com", "registration": {"auto_renew": true, "created_at": "2025-10-27T10:00:00Z", "domain_name": "example.com", "expires_at": "2026-10-27T10:00:00Z", "locked": true, "privacy_mode": "redaction", "status": "active"}}, "created_at": "2025-10-27T10:00:00Z", "links": {"resource": "/accounts/abc/registrar/registrations/example.com", "self": "/accounts/abc/registrar/registrations/example.com/registration-status"}, "state": "succeeded", "updated_at": "2025-10-27T10:00:03Z"}, "success": true}}}, "schema": {"$ref": "#/components/schemas/registrar-api-sandbox_workflow-status-response-single"}}}}, "202": {"description": "Registration is still processing. This occurs when the operation did\nnot complete within the synchronous wait window, or when `Prefer: respond-async`\nwas sent. Poll the URL in `links.self` to track progress.\n", "headers": {"Location": {"description": "URL of the workflow status resource (same as `links.self`).", "schema": {"type": "string"}}, "Preference-Applied": {"description": "Set to `respond-async` when the server honored the preference.", "schema": {"type": "string"}}}, "content": {"application/json": {"examples": {"still_processing": {"summary": "Registration still in progress", "value": {"errors": [], "messages": [], "result": {"completed": false, "context": {"domain_name": "example.com"}, "created_at": "2025-10-27T10:00:00Z", "links": {"resource": "/accounts/abc/registrar/registrations/example.com", "self": "/accounts/abc/registrar/registrations/example.com/registration-status"}, "state": "in_progress", "updated_at": "2025-10-27T10:00:10Z"}, "success": true}}}, "schema": {"$ref": "#/components/schemas/registrar-api-sandbox_workflow-status-response-single"}}}}, "4XX": {"description": "Create registration failure.", "content": {"application/json": {"examples": {"account_domain_limit_exceeded": {"description": "The account already owns the maximum supported number of domains.\nDomains created through either the dashboard or this API count\ntoward the same account-wide limit.\n", "summary": "Account has reached the domain limit", "value": {"errors": [{"code": 10000, "message": "Domain limit reached: you cannot register more than 100 domains.", "source": {"pointer": "/domain_name"}}], "messages": [], "result": null, "success": false}}, "domain_name_required": {"description": "The required `domain_name` field was not provided.", "summary": "Missing domain name", "value": {"errors": [{"code": 10000, "message": "domain_name is required", "source": {"pointer": "/domain_name"}}], "messages": [], "result": null, "success": false}}, "invalid_auto_renew": {"description": "The `auto_renew` field must be a boolean if present.", "summary": "Invalid auto_renew value", "value": {"errors": [{"code": 10000, "message": "Must be a boolean", "source": {"pointer": "/auto_renew"}}], "messages": [], "result": null, "success": false}}, "missing_contact": {"description": "The request did not include a registrant contact, and the account\nhas no default address book entry to fall back to. Either provide\ncontacts.registrant in the request or set up a default address\nbook entry at `https://dash.cloudflare.com/{account_id}/domains/registrations`.\n", "summary": "No registrant contact and no default address book entry", "value": {"errors": [{"code": 10000, "message": "No registrant contact provided and no default address book entry found for this account."}], "messages": [], "result": null, "success": false}}, "registration_not_supported": {"description": "The extension is not supported for registration through this API.\nThis includes extensions that are unsupported by the API or do\nnot support the required registration contact flow.\n", "summary": "Registration not supported for extension", "value": {"errors": [{"code": 10000, "message": "Registration is not supported for this extension"}], "messages": [], "result": null, "success": false}}}, "schema": {"$ref": "#/components/schemas/registrar-api-sandbox_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Registrar Registration"]}
```
