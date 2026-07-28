---
title: List Zero Trust Gateway DNS destination IPv4 address pairs
page_id: operation-get-accounts-account-id-gateway-dns-destination-ips-820e3f8f
path: operations/zero-trust-gateway-dns-destination-ipv4-address-pairs
description: List Zero Trust Gateway IPv4 address pairs for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/dns_destination_ips
operation_ids:
    - zero-trust-dns-destination-ips-list-dns-destination-ips
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Zero Trust Gateway DNS destination IPv4 address pairs

`GET /accounts/{account_id}/gateway/dns_destination_ips`

Operation ID: `zero-trust-dns-destination-ips-list-dns-destination-ips`

List Zero Trust Gateway IPv4 address pairs for an account.

## Definition

```yaml
{"operationId": "zero-trust-dns-destination-ips-list-dns-destination-ips", "summary": "List Zero Trust Gateway DNS destination IPv4 address pairs", "description": "List Zero Trust Gateway IPv4 address pairs for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier"}}], "responses": {"200": {"description": "List Zero Trust Gateway DNS destination IPv4 address pairs response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_response_collection-2"}}}}, "4XX": {"description": "List Zero Trust Gateway DNS destination IPv4 address pairs response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_response_collection-2"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway DNS destination IPv4 address pairs"], "x-api-token-group": ["Cloudflare Zero Trust Secure DNS Locations Write", "Zero Trust Read", "Zero Trust Write"]}
```
