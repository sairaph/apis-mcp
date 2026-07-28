---
title: Create your Zero Trust organization
page_id: operation-post-accounts-account-id-access-organizations-8a8ee7ad
path: operations/zero-trust-organization
description: Sets up a Zero Trust organization for your account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/organizations
operation_ids:
    - zero-trust-organization-create-your-zero-trust-organization
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create your Zero Trust organization

`POST /accounts/{account_id}/access/organizations`

Operation ID: `zero-trust-organization-create-your-zero-trust-organization`

Sets up a Zero Trust organization for your account.

## Definition

```yaml
{"operationId": "zero-trust-organization-create-your-zero-trust-organization", "summary": "Create your Zero Trust organization", "description": "Sets up a Zero Trust organization for your account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"allow_authenticate_via_warp": {"$ref": "#/components/schemas/access_allow_authenticate_via_warp"}, "auth_domain": {"$ref": "#/components/schemas/access_auth_domain"}, "auto_redirect_to_identity": {"$ref": "#/components/schemas/access_auto_redirect_to_identity"}, "deny_unmatched_requests": {"$ref": "#/components/schemas/access_deny_unmatched_requests"}, "deny_unmatched_requests_exempted_zone_names": {"$ref": "#/components/schemas/access_deny_unmatched_requests_exempted_zone_names"}, "is_ui_read_only": {"$ref": "#/components/schemas/access_is_ui_read_only"}, "login_design": {"$ref": "#/components/schemas/access_login_design"}, "mfa_config": {"$ref": "#/components/schemas/access_org_mfa_config"}, "mfa_piv_key_requirements": {"$ref": "#/components/schemas/access_mfa_piv_key_requirements"}, "mfa_required_for_all_apps": {"$ref": "#/components/schemas/access_mfa_required_for_all_apps"}, "name": {"$ref": "#/components/schemas/access_name"}, "session_duration": {"$ref": "#/components/schemas/access_session_duration"}, "ui_read_only_toggle_reason": {"$ref": "#/components/schemas/access_ui_read_only_toggle_reason"}, "user_seat_expiration_inactive_time": {"$ref": "#/components/schemas/access_user_seat_expiration_inactive_time"}, "warp_auth_session_duration": {"$ref": "#/components/schemas/access_warp_auth_session_duration"}}, "required": ["name", "auth_domain"]}}}}, "responses": {"201": {"description": "Create your Zero Trust organization response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response"}}}}, "4XX": {"description": "Create your Zero Trust organization response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust organization"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.organizations", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
