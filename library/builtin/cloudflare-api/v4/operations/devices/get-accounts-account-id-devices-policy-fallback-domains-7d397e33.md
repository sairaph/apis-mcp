---
title: Get your Local Domain Fallback list
page_id: operation-get-accounts-account-id-devices-policy-fallback-domains-550e73d3
path: operations/devices
description: Fetches a list of domains to bypass Gateway DNS resolution. These domains will use the specified local DNS resolver instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/policy/fallback_domains
operation_ids:
    - devices-get-local-domain-fallback-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get your Local Domain Fallback list

`GET /accounts/{account_id}/devices/policy/fallback_domains`

Operation ID: `devices-get-local-domain-fallback-list`

Fetches a list of domains to bypass Gateway DNS resolution. These domains will use the specified local DNS resolver instead.

## Definition

```yaml
{"operationId": "devices-get-local-domain-fallback-list", "summary": "Get your Local Domain Fallback list", "description": "Fetches a list of domains to bypass Gateway DNS resolution. These domains will use the specified local DNS resolver instead.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Get your Local Domain Fallback list response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_fallback_domain_response_collection"}}}}, "4XX": {"description": "Get your Local Domain Fallback list response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_fallback_domain_response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.default.fallback-domains", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
