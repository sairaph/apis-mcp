---
title: Get the Local Domain Fallback list for a device settings profile
page_id: operation-get-accounts-account-id-devices-policy-policy-id-fallback-domains-70266118
path: operations/devices
description: Fetches the list of domains to bypass Gateway DNS resolution from a specified device settings profile. These domains will use the specified local DNS resolver instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/policy/{policy_id}/fallback_domains
operation_ids:
    - devices-get-local-domain-fallback-list-for-a-device-settings-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the Local Domain Fallback list for a device settings profile

`GET /accounts/{account_id}/devices/policy/{policy_id}/fallback_domains`

Operation ID: `devices-get-local-domain-fallback-list-for-a-device-settings-policy`

Fetches the list of domains to bypass Gateway DNS resolution from a specified device settings profile. These domains will use the specified local DNS resolver instead.

## Definition

```yaml
{"operationId": "devices-get-local-domain-fallback-list-for-a-device-settings-policy", "summary": "Get the Local Domain Fallback list for a device settings profile", "description": "Fetches the list of domains to bypass Gateway DNS resolution from a specified device settings profile. These domains will use the specified local DNS resolver instead.", "parameters": [{"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_schemas-uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Get the Local Domain Fallback list for a device settings profile response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_fallback_domain_response_collection"}}}}, "4XX": {"description": "Get the Local Domain Fallback list for a device settings profile response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_fallback_domain_response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.custom.fallback-domains", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
