---
title: Update your Zero Trust organization
page_id: operation-put-zones-zone-id-access-organizations-5c3a58e3
path: operations/zone-level-zero-trust-organization
description: Updates the configuration for your Zero Trust organization.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/access/organizations
operation_ids:
    - zone-level-zero-trust-organization-update-your-zero-trust-organization
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update your Zero Trust organization

`PUT /zones/{zone_id}/access/organizations`

Operation ID: `zone-level-zero-trust-organization-update-your-zero-trust-organization`

Updates the configuration for your Zero Trust organization.

## Definition

```yaml
{"operationId": "zone-level-zero-trust-organization-update-your-zero-trust-organization", "summary": "Update your Zero Trust organization", "description": "Updates the configuration for your Zero Trust organization.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier-4"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"auth_domain": {"$ref": "#/components/schemas/access_auth_domain-2"}, "is_ui_read_only": {"$ref": "#/components/schemas/access_is_ui_read_only-2"}, "login_design": {"$ref": "#/components/schemas/access_login_design-2"}, "name": {"$ref": "#/components/schemas/access_name-15"}, "ui_read_only_toggle_reason": {"$ref": "#/components/schemas/access_ui_read_only_toggle_reason"}, "user_seat_expiration_inactive_time": {"$ref": "#/components/schemas/access_user_seat_expiration_inactive_time-2"}}}}}}, "responses": {"200": {"description": "Update your Zero Trust organization response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-15"}}}}, "4XX": {"description": "Update your Zero Trust organization response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Zero Trust organization"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.zone-organizations", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
