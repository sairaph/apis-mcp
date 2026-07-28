---
title: Set the Local Domain Fallback list for a device settings profile
page_id: operation-put-accounts-account-id-devices-policy-policy-id-fallback-domains-792cc5dd
path: operations/devices
description: Sets the list of domains to bypass Gateway DNS resolution. These domains will use the specified local DNS resolver instead. This will only apply to the specified device settings profile.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/devices/policy/{policy_id}/fallback_domains
operation_ids:
    - devices-set-local-domain-fallback-list-for-a-device-settings-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set the Local Domain Fallback list for a device settings profile

`PUT /accounts/{account_id}/devices/policy/{policy_id}/fallback_domains`

Operation ID: `devices-set-local-domain-fallback-list-for-a-device-settings-policy`

Sets the list of domains to bypass Gateway DNS resolution. These domains will use the specified local DNS resolver instead. This will only apply to the specified device settings profile.

## Definition

```yaml
{"operationId": "devices-set-local-domain-fallback-list-for-a-device-settings-policy", "summary": "Set the Local Domain Fallback list for a device settings profile", "description": "Sets the list of domains to bypass Gateway DNS resolution. These domains will use the specified local DNS resolver instead. This will only apply to the specified device settings profile.", "parameters": [{"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_schemas-uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_fallback_domain"}}}}}, "responses": {"200": {"description": "Set the Local Domain Fallback list for a device settings profile response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_fallback_domain_response_collection"}}}}, "4XX": {"description": "Set the Local Domain Fallback list for a device settings profile response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_fallback_domain_response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.custom.fallback-domains", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
