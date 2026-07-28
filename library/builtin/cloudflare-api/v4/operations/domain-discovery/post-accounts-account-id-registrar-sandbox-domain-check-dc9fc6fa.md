---
title: Check domain availability
page_id: operation-post-accounts-account-id-registrar-sandbox-domain-check-3f955316
path: operations/domain-discovery
description: |-
    Performs real-time, authoritative availability checks directly against domain
    registries. Use this endpoint to verify a domain is available before attempting
    registration via `POST /registrations`.

    **Important:** Unlike the Search endpoint, these results are authoritative and
    reflect current registry status. Always check availability immediately before
    registration as domain status can change rapidly.

    **Note:** This endpoint uses POST to accept a list of domains in the request
    body. It is a read-only operation — it does not create, modify, or reserve
    any domains.

    ### Extension support

    Only domains on extensions supported for programmatic registration by this API
    can be registered. If you check a domain on an unsupported extension, the response
    will include `registrable: false` with a `reason` field explaining why:

    - `extension_not_supported_via_api` — Cloudflare Registrar supports this extension
      in the dashboard, but it is not yet available for programmatic registration via
      this API. Register via `https://dash.cloudflare.com/{account_id}/domains/registrations` instead.
    - `extension_not_supported` — This extension is not supported by Cloudflare
      Registrar.
    - `extension_disallows_registration` — The extension's registry has temporarily
      or permanently frozen new registrations. No registrar can register domains on
      this extension at this time.
    - `domain_premium` — The domain is premium priced. Premium registration is not
      currently supported by this API.
    - `domain_unavailable` — The domain is already registered, reserved, or otherwise
      not available for registration on a supported extension.

    The `reason` field is only present when `registrable` is `false`.

    ### Behavior
    - Maximum 20 domains per request
    - Pricing is only returned for domains where `registrable: true`
    - Results are not cached; each request queries the registry

    ### Workflow
    1. Call this endpoint with domains the user wants to register.
    2. For each domain where `registrable: true`, present pricing to the user.
    3. If `tier: premium`, note that premium registration is not currently
       supported by this API and do not proceed to `POST /registrations`.
    4. Proceed to `POST /registrations` only for supported non-premium domains.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/registrar-sandbox/domain-check
operation_ids:
    - sandbox-registrar-domain-discovery-check
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Check domain availability

`POST /accounts/{account_id}/registrar-sandbox/domain-check`

Operation ID: `sandbox-registrar-domain-discovery-check`

Performs real-time, authoritative availability checks directly against domain
registries. Use this endpoint to verify a domain is available before attempting
registration via `POST /registrations`.

**Important:** Unlike the Search endpoint, these results are authoritative and
reflect current registry status. Always check availability immediately before
registration as domain status can change rapidly.

**Note:** This endpoint uses POST to accept a list of domains in the request
body. It is a read-only operation — it does not create, modify, or reserve
any domains.

### Extension support

Only domains on extensions supported for programmatic registration by this API
can be registered. If you check a domain on an unsupported extension, the response
will include `registrable: false` with a `reason` field explaining why:

- `extension_not_supported_via_api` — Cloudflare Registrar supports this extension
  in the dashboard, but it is not yet available for programmatic registration via
  this API. Register via `https://dash.cloudflare.com/{account_id}/domains/registrations` instead.
- `extension_not_supported` — This extension is not supported by Cloudflare
  Registrar.
- `extension_disallows_registration` — The extension's registry has temporarily
  or permanently frozen new registrations. No registrar can register domains on
  this extension at this time.
- `domain_premium` — The domain is premium priced. Premium registration is not
  currently supported by this API.
- `domain_unavailable` — The domain is already registered, reserved, or otherwise
  not available for registration on a supported extension.

The `reason` field is only present when `registrable` is `false`.

### Behavior
- Maximum 20 domains per request
- Pricing is only returned for domains where `registrable: true`
- Results are not cached; each request queries the registry

### Workflow
1. Call this endpoint with domains the user wants to register.
2. For each domain where `registrable: true`, present pricing to the user.
3. If `tier: premium`, note that premium registration is not currently
   supported by this API and do not proceed to `POST /registrations`.
4. Proceed to `POST /registrations` only for supported non-premium domains.

## Definition

```yaml
{"operationId": "sandbox-registrar-domain-discovery-check", "summary": "Check domain availability", "description": "Performs real-time, authoritative availability checks directly against domain\nregistries. Use this endpoint to verify a domain is available before attempting\nregistration via `POST /registrations`.\n\n**Important:** Unlike the Search endpoint, these results are authoritative and\nreflect current registry status. Always check availability immediately before\nregistration as domain status can change rapidly.\n\n**Note:** This endpoint uses POST to accept a list of domains in the request\nbody. It is a read-only operation — it does not create, modify, or reserve\nany domains.\n\n### Extension support\n\nOnly domains on extensions supported for programmatic registration by this API\ncan be registered. If you check a domain on an unsupported extension, the response\nwill include `registrable: false` with a `reason` field explaining why:\n\n- `extension_not_supported_via_api` — Cloudflare Registrar supports this extension\n  in the dashboard, but it is not yet available for programmatic registration via\n  this API. Register via `https://dash.cloudflare.com/{account_id}/domains/registrations` instead.\n- `extension_not_supported` — This extension is not supported by Cloudflare\n  Registrar.\n- `extension_disallows_registration` — The extension's registry has temporarily\n  or permanently frozen new registrations. No registrar can register domains on\n  this extension at this time.\n- `domain_premium` — The domain is premium priced. Premium registration is not\n  currently supported by this API.\n- `domain_unavailable` — The domain is already registered, reserved, or otherwise\n  not available for registration on a supported extension.\n\nThe `reason` field is only present when `registrable` is `false`.\n\n### Behavior\n- Maximum 20 domains per request\n- Pricing is only returned for domains where `registrable: true`\n- Results are not cached; each request queries the registry\n\n### Workflow\n1. Call this endpoint with domains the user wants to register.\n2. For each domain where `registrable: true`, present pricing to the user.\n3. If `tier: premium`, note that premium registration is not currently\n   supported by this API and do not proceed to `POST /registrations`.\n4. Proceed to `POST /registrations` only for supported non-premium domains.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID. Required for all Registrar API operations.", "required": true, "schema": {"$ref": "#/components/schemas/registrar-api-sandbox_identifier"}}], "requestBody": {"description": "List of fully qualified domain names (FQDNs) to check. Each domain must include the extension (e.g., \"example.com\", not \"example\").", "required": true, "content": {"application/json": {"examples": {"bulk_check": {"description": "Check the same name across multiple supported extensions", "summary": "Bulk check across domain extensions", "value": {"domains": ["myawesomebrand.com", "myawesomebrand.net", "myawesomebrand.org", "myawesomebrand.app", "myawesomebrand.dev"]}}, "multiple_domains": {"summary": "Check multiple domains", "value": {"domains": ["example.com", "example.net", "mycompany.dev"]}}, "single_domain": {"summary": "Check single domain", "value": {"domains": ["example.com"]}}}, "schema": {"$ref": "#/components/schemas/registrar-api-sandbox_domain_check_request"}}}}, "responses": {"200": {"description": "Successfully returned availability results. Each requested domain appears in the `domains` array with its current availability status and pricing (if available).", "content": {"application/json": {"examples": {"all_available": {"description": "Unique names may be available across domain extensions", "summary": "All domains available", "value": {"errors": [], "messages": [], "result": {"domains": [{"name": "xq7mz9brand.com", "pricing": {"currency": "USD", "registration_cost": "8.57", "renewal_cost": "8.57"}, "registrable": true, "tier": "standard"}, {"name": "xq7mz9brand.net", "pricing": {"currency": "USD", "registration_cost": "9.95", "renewal_cost": "9.95"}, "registrable": true, "tier": "standard"}]}, "success": true}}, "all_unavailable": {"description": "Common names are often already registered", "summary": "All domains taken", "value": {"errors": [], "messages": [], "result": {"domains": [{"name": "example.com", "reason": "domain_unavailable", "registrable": false}, {"name": "example.net", "reason": "domain_unavailable", "registrable": false}, {"name": "example.org", "reason": "domain_unavailable", "registrable": false}]}, "success": true}}, "mixed_availability": {"description": "Some domains available, some taken", "summary": "Mixed availability results", "value": {"errors": [], "messages": [], "result": {"domains": [{"name": "cloudflare.com", "reason": "domain_unavailable", "registrable": false}, {"name": "my-unique-startup-name-2024.com", "pricing": {"currency": "USD", "registration_cost": "8.57", "renewal_cost": "8.57"}, "registrable": true, "tier": "standard"}, {"name": "my-unique-startup-name-2024.dev", "pricing": {"currency": "USD", "registration_cost": "10.11", "renewal_cost": "10.11"}, "registrable": true, "tier": "standard"}]}, "success": true}}, "premium_domain": {"description": "Premium pricing may be surfaced by this endpoint, but premium registration is not currently supported by this API.", "summary": "Premium domain surfaced but not registrable", "value": {"errors": [], "messages": [], "result": {"domains": [{"name": "coffee.xyz", "reason": "domain_premium", "registrable": false, "tier": "premium"}]}, "success": true}}, "registry_frozen": {"description": "The extension's registry has temporarily or permanently stopped accepting new registrations.", "summary": "Registry has frozen new registrations", "value": {"errors": [], "messages": [], "result": {"domains": [{"name": "example.py", "reason": "extension_disallows_registration", "registrable": false}]}, "success": true}}, "unsupported_extension": {"description": "The extension is not supported by Cloudflare Registrar at all.", "summary": "Extension not supported by Cloudflare Registrar", "value": {"errors": [], "messages": [], "result": {"domains": [{"name": "example.horse", "reason": "extension_not_supported", "registrable": false}]}, "success": true}}, "unsupported_via_api": {"description": "The .uk extension is supported by Cloudflare Registrar in the dashboard but not yet available for programmatic registration via this API.", "summary": "Extension supported in dashboard but not via API", "value": {"errors": [], "messages": [], "result": {"domains": [{"name": "mybrand.uk", "reason": "extension_not_supported_via_api", "registrable": false}, {"name": "mybrand.com", "pricing": {"currency": "USD", "registration_cost": "8.57", "renewal_cost": "8.57"}, "registrable": true, "tier": "standard"}]}, "success": true}}}, "schema": {"$ref": "#/components/schemas/registrar-api-sandbox_domain_check_response"}}}}, "400": {"description": "Invalid request parameters. Common causes:\n- Empty `domains` array\n- Exceeds maximum of 20 domains per request\n- Malformed request body\n- None of the provided domains are valid\n", "content": {"application/json": {"examples": {"domains_required": {"description": "The `domains` array must contain at least one domain name to check.", "summary": "Domains array missing or empty", "value": {"errors": [{"code": 1006, "message": "domains array must contain at least one domain"}], "messages": [], "result": null, "success": false}}, "no_valid_domains": {"description": "None of the provided domains are valid after normalization and validation.", "summary": "No valid domains", "value": {"errors": [{"code": 1008, "message": "None of the provided domains are valid or have supported extensions"}], "messages": [], "result": null, "success": false}}, "too_many_domains": {"description": "The request exceeds the maximum limit of 20 domains per request. Split into multiple requests.", "summary": "Too many domains", "value": {"errors": [{"code": 1007, "message": "domains array exceeds maximum of 20 domains"}], "messages": [], "result": null, "success": false}}}, "schema": {"$ref": "#/components/schemas/registrar-api-sandbox_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Domain Discovery"]}
```
