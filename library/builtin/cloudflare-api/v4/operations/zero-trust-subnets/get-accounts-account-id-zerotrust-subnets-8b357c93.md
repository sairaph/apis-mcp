---
title: List Subnets
page_id: operation-get-accounts-account-id-zerotrust-subnets-097210c8
path: operations/zero-trust-subnets
description: Lists and filters subnets in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/zerotrust/subnets
operation_ids:
    - zero-trust-networks-subnets-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Subnets

`GET /accounts/{account_id}/zerotrust/subnets`

Operation ID: `zero-trust-networks-subnets-list`

Lists and filters subnets in an account.

## Definition

```yaml
{"operationId": "zero-trust-networks-subnets-list", "summary": "List Subnets", "description": "Lists and filters subnets in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "name", "in": "query", "description": "If set, only list subnets with the given name", "schema": {"$ref": "#/components/schemas/tunnel_subnet_query_name"}}, {"name": "comment", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_subnet_query_comment"}}, {"name": "network", "in": "query", "schema": {"description": "If set, only list the subnet whose network exactly matches the given CIDR.", "allOf": [{"$ref": "#/components/schemas/tunnel_ip_network_encoded"}]}}, {"name": "existed_at", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_existed_at"}}, {"name": "address_family", "in": "query", "description": "If set, only include subnets in the given address family - `v4` or `v6`", "schema": {"$ref": "#/components/schemas/tunnel_address_family"}}, {"name": "is_default_network", "in": "query", "schema": {"description": "If `true`, only include default subnets. If `false`, exclude default subnets subnets. If not set, all subnets will be included.", "type": "boolean"}}, {"name": "is_deleted", "in": "query", "schema": {"description": "If `true`, only include deleted subnets. If `false`, exclude deleted subnets. If not set, all subnets will be included.", "type": "boolean"}}, {"name": "sort_order", "in": "query", "schema": {"description": "Sort order of the results. `asc` means oldest to newest, `desc` means newest to oldest. If not set, they will not be in any particular order.", "type": "string", "enum": ["asc", "desc"]}}, {"name": "subnet_types", "in": "query", "schema": {"description": "If set, the types of subnets to include, separated by comma.", "type": "string", "example": "cloudflare_source,warp", "enum": ["cloudflare_source", "initial_resolved_ip", "warp"]}}, {"name": "per_page", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_per_page"}}, {"name": "page", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_page_number"}}], "responses": {"200": {"description": "List subnets response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_subnet_response_collection"}}}}, "4XX": {"description": "List subnets response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_subnet_response_collection"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Zero Trust Subnets"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare One Networks Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.subnets", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
