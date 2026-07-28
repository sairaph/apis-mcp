---
title: Get IP Overview
page_id: operation-get-accounts-account-id-intel-ip-76df342a
path: operations/ip-intelligence
description: Gets the geolocation, ASN, infrastructure type of the ASN, and any security threat categories of an IP address. **Must provide ip query parameters.** For example, `/intel/ip?ipv4=1.1.1.1` or `/intel/ip?ipv6=2001:db8::1`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/ip
operation_ids:
    - ip-intelligence-get-ip-overview
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get IP Overview

`GET /accounts/{account_id}/intel/ip`

Operation ID: `ip-intelligence-get-ip-overview`

Gets the geolocation, ASN, infrastructure type of the ASN, and any security threat categories of an IP address. **Must provide ip query parameters.** For example, `/intel/ip?ipv4=1.1.1.1` or `/intel/ip?ipv6=2001:db8::1`.

## Definition

```yaml
{"operationId": "ip-intelligence-get-ip-overview", "summary": "Get IP Overview", "description": "Gets the geolocation, ASN, infrastructure type of the ASN, and any security threat categories of an IP address. **Must provide ip query parameters.** For example, `/intel/ip?ipv4=1.1.1.1` or `/intel/ip?ipv6=2001:db8::1`.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/intel_identifier"}}, {"name": "ipv4", "in": "query", "schema": {"type": "string"}}, {"name": "ipv6", "in": "query", "schema": {"type": "string"}}], "responses": {"200": {"description": "Get IP Overview response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel_schemas-response"}}}}, "4XX": {"description": "Get IP Overview response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/intel_schemas-response"}, {"$ref": "#/components/schemas/intel_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["IP Intelligence"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
