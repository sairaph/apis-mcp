---
title: Get resource change history from an organization audit log entry (Version 2)
page_id: operation-get-organizations-organization-id-logs-audit-id-history-f52211b5
path: operations/audit-logs
description: |-
    Returns the chronological change history for the resource identified by the given organization-scoped audit log entry.

    The endpoint first locates the source audit log entry by `id` (using `action_time` to narrow the lookup window), derives identifying filters from that entry, and then returns matching audit logs within the `since`/`before` window.

    The `result_info.history_status` field indicates the quality of the resource identification used:
    - `exact`: Resource was identified by the resource URI.
    - `approximate`: Resource was identified without the resource URI.
    - `unavailable`: The source audit log entry did not contain enough information to identify the resource; an empty result is returned.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /organizations/{organization_id}/logs/audit/{id}/history
operation_ids:
    - audit-logs-v2-get-organization-audit-log-history
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get resource change history from an organization audit log entry (Version 2)

`GET /organizations/{organization_id}/logs/audit/{id}/history`

Operation ID: `audit-logs-v2-get-organization-audit-log-history`

Returns the chronological change history for the resource identified by the given organization-scoped audit log entry.

The endpoint first locates the source audit log entry by `id` (using `action_time` to narrow the lookup window), derives identifying filters from that entry, and then returns matching audit logs within the `since`/`before` window.

The `result_info.history_status` field indicates the quality of the resource identification used:
- `exact`: Resource was identified by the resource URI.
- `approximate`: Resource was identified without the resource URI.
- `unavailable`: The source audit log entry did not contain enough information to identify the resource; an empty result is returned.

## Definition

```yaml
{"operationId": "audit-logs-v2-get-organization-audit-log-history", "summary": "Get resource change history from an organization audit log entry (Version 2)", "description": "Returns the chronological change history for the resource identified by the given organization-scoped audit log entry.\n\nThe endpoint first locates the source audit log entry by `id` (using `action_time` to narrow the lookup window), derives identifying filters from that entry, and then returns matching audit logs within the `since`/`before` window.\n\nThe `result_info.history_status` field indicates the quality of the resource identification used:\n- `exact`: Resource was identified by the resource URI.\n- `approximate`: Resource was identified without the resource URI.\n- `unavailable`: The source audit log entry did not contain enough information to identify the resource; an empty result is returned.\n", "parameters": [{"name": "organization_id", "in": "path", "required": true, "schema": {"description": "The unique ID that identifies the organization.", "type": "string", "example": "a67e14daa5f8dceeb91fe5449ba496ef"}}, {"name": "id", "in": "path", "required": true, "schema": {"description": "The ID of the audit log to fetch resource history for.", "type": "string", "format": "uuid", "example": "f174be97-19b1-40d6-954d-70cd5fbd52db"}}, {"name": "action_time", "in": "query", "description": "RFC3339 timestamp of the source audit log entry's action time. Used to narrow the source-entry lookup window. Provide the `action.time` value from the audit log identified by `id`.", "required": true, "schema": {"type": "string", "format": "date-time", "example": "2024-10-30T15:00:00Z"}}, {"name": "since", "in": "query", "description": "Limits the returned results to logs newer than the specified date. This can be a date string 2019-04-30 (interpreted in UTC) or an absolute timestamp that conforms to RFC3339.", "required": true, "schema": {"type": "string", "format": "date", "example": "2024-10-30"}}, {"name": "before", "in": "query", "description": "Limits the returned results to logs older than the specified date. This can be a date string 2019-04-30 (interpreted in UTC) or an absolute timestamp that conforms to RFC3339.", "required": true, "schema": {"type": "string", "format": "date", "example": "2024-10-31"}}, {"name": "direction", "in": "query", "schema": {"description": "Sets sorting order.", "type": "string", "example": "desc", "default": "desc", "enum": ["desc", "asc"]}}, {"name": "limit", "in": "query", "schema": {"description": "The number limits the objects to return. The cursor attribute may be used to iterate over the next batch of objects if there are more than the limit.", "type": "number", "example": 25, "default": 100, "maximum": 1000, "minimum": 1}}, {"name": "cursor", "in": "query", "schema": {"description": "The cursor is an opaque token used to paginate through large sets of records. It indicates the position from which to continue when requesting the next set of records. A valid cursor value can be obtained from the cursor object in the result_info structure of a previous response.", "type": "string", "example": "Q1buH-__DQqqig7SVYXT-SsMOTGY2Z3Y80W-fGgva7yaDdmPKveucH5ddOcHsJRhNb-xUK8agZQqkJSMAENGO8NU6g=="}}], "responses": {"200": {"description": "Get organization resource history successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_audit-logs-v2-org-history-response-collection"}}}}, "404": {"description": "Audit log entry not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_audit-logs-v2-history-not-found"}}}}, "4XX": {"description": "Get organization resource history failed response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_api-response-common-failure-2"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Audit Logs"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
