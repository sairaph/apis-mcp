---
title: List all targets
page_id: operation-get-accounts-account-id-infrastructure-targets-7d4573dc
path: operations/infrastructure-access-targets
description: |-
    Lists and sorts an account’s targets. Filters are optional and are ANDed
    together.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/infrastructure/targets
operation_ids:
    - infra-targets-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all targets

`GET /accounts/{account_id}/infrastructure/targets`

Operation ID: `infra-targets-list`

Lists and sorts an account’s targets. Filters are optional and are ANDed
together.

## Definition

```yaml
{"operationId": "infra-targets-list", "summary": "List all targets", "description": "Lists and sorts an account’s targets. Filters are optional and are ANDed\ntogether.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/infra_AccountTag"}}, {"name": "hostname", "in": "query", "description": "Hostname of a target", "schema": {"type": "string", "nullable": true}}, {"name": "hostname_contains", "in": "query", "description": "Partial match to the hostname of a target", "schema": {"type": "string", "nullable": true}}, {"name": "virtual_network_id", "in": "query", "description": "Private virtual network identifier of the target", "schema": {"type": "string", "format": "uuid", "nullable": true}}, {"name": "ip_v4", "in": "query", "description": "IPv4 address of the target", "schema": {"type": "string", "nullable": true}}, {"name": "ip_v6", "in": "query", "description": "IPv6 address of the target", "schema": {"type": "string", "nullable": true}}, {"name": "created_before", "in": "query", "description": "Date and time at which the target was created before (inclusive)", "schema": {"type": "string", "format": "date-time", "nullable": true}}, {"name": "created_after", "in": "query", "description": "Date and time at which the target was created after (inclusive)", "schema": {"type": "string", "format": "date-time", "nullable": true}}, {"name": "modified_before", "in": "query", "description": "Date and time at which the target was modified before (inclusive)", "schema": {"type": "string", "format": "date-time", "nullable": true}}, {"name": "modified_after", "in": "query", "description": "Date and time at which the target was modified after (inclusive)", "schema": {"type": "string", "format": "date-time", "nullable": true}}, {"name": "ips", "in": "query", "description": "Filters for targets that have any of the following IP addresses. Specify\n`ips` multiple times in query parameter to build list of candidates.", "schema": {"type": "array", "items": {"type": "string"}}}, {"name": "target_ids", "in": "query", "description": "Filters for targets that have any of the following UUIDs. Specify\n`target_ids` multiple times in query parameter to build list of\ncandidates.", "schema": {"type": "array", "items": {"format": "uuid", "type": "string"}}}, {"name": "ip_like", "in": "query", "description": "Filters for targets whose IP addresses look like the specified string.\nSupports `*` as a wildcard character", "schema": {"type": "string", "nullable": true}}, {"name": "ipv4_start", "in": "query", "description": "Defines an IPv4 filter range's starting value (inclusive). Requires\n`ipv4_end` to be specified as well.", "schema": {"type": "string", "nullable": true}}, {"name": "ipv4_end", "in": "query", "description": "Defines an IPv4 filter range's ending value (inclusive). Requires\n`ipv4_start` to be specified as well.", "schema": {"type": "string", "nullable": true}}, {"name": "ipv6_start", "in": "query", "description": "Defines an IPv6 filter range's starting value (inclusive). Requires\n`ipv6_end` to be specified as well.", "schema": {"type": "string", "nullable": true}}, {"name": "ipv6_end", "in": "query", "description": "Defines an IPv6 filter range's ending value (inclusive). Requires\n`ipv6_start` to be specified as well.", "schema": {"type": "string", "nullable": true}}, {"name": "page", "in": "query", "description": "Current page in the response", "schema": {"type": "integer", "format": "int32", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Max amount of entries returned per page", "schema": {"type": "integer", "format": "int32", "default": 1000, "maximum": 1000, "minimum": 1}}, {"name": "order", "in": "query", "description": "The field to sort by.", "schema": {"type": "string", "enum": ["hostname", "created_at"]}}, {"name": "direction", "in": "query", "description": "The sorting direction.", "schema": {"allOf": [{"$ref": "#/components/schemas/infra_SortingDirection"}]}}], "responses": {"200": {"description": "Successfully retrieved all targets in the account", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/infra_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/infra_TargetArray"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to retrieve all targets in the account", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/infra_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Infrastructure Access Targets"]}
```
