---
title: Search email messages
page_id: operation-get-accounts-account-id-email-security-investigate-87ad40a8
path: operations/email-security
description: Returns information for each email that matches the search parameter(s).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/investigate
operation_ids:
    - email_security_investigate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Search email messages

`GET /accounts/{account_id}/email-security/investigate`

Operation ID: `email_security_investigate`

Returns information for each email that matches the search parameter(s).

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_investigate", "summary": "Search email messages", "description": "Returns information for each email that matches the search parameter(s).", "parameters": [{"name": "start", "in": "query", "description": "The beginning of the search date range. Defaults to `now - 30 days`. Must not be in the future.", "schema": {"type": "string", "format": "date-time"}, "example": "2022-06-25T14:30:00Z"}, {"name": "end", "in": "query", "description": "The end of the search date range. Defaults to `now`.", "schema": {"type": "string", "format": "date-time"}, "example": "2022-07-25T14:30:00Z"}, {"name": "query", "in": "query", "description": "Space-delimited search term. Case-insensitive.", "schema": {"type": "string"}, "example": "bob jones"}, {"name": "detections_only", "in": "query", "description": "Whether to include only detections in search results.", "schema": {"type": "boolean", "default": true}}, {"name": "final_disposition", "in": "query", "description": "Dispositions to filter by.", "schema": {"type": "string", "enum": ["MALICIOUS", "SUSPICIOUS", "SPOOF", "SPAM", "BULK", "NONE"]}}, {"name": "metric", "in": "query", "schema": {"type": "string"}}, {"name": "message_action", "in": "query", "description": "Message actions to filter by.", "schema": {"type": "string", "enum": ["PREVIEW", "QUARANTINE_RELEASED", "MOVED"]}}, {"name": "recipient", "in": "query", "schema": {"type": "string"}, "example": "me@example.com"}, {"name": "sender", "in": "query", "schema": {"type": "string"}, "example": "noreply@example.com"}, {"name": "alert_id", "in": "query", "schema": {"type": "string"}}, {"name": "domain", "in": "query", "description": "Sender domains to filter by.", "schema": {"type": "string"}, "example": "example.com"}, {"name": "message_id", "in": "query", "schema": {"type": "string"}}, {"name": "subject", "in": "query", "schema": {"type": "string"}}, {"name": "delivery_status", "in": "query", "description": "Delivery status to filter by.", "schema": {"$ref": "#/components/schemas/email-security_MessageDeliveryStatus"}}, {"name": "cursor", "in": "query", "schema": {"type": "string"}}, {"$ref": "#/components/parameters/email-security_per_page"}, {"name": "page", "in": "query", "description": "Deprecated: Use cursor pagination instead. End of life: November 1, 2026.", "schema": {"type": "integer", "default": 1, "minimum": 1, "nullable": true}, "deprecated": true, "x-stainless-deprecation-message": "Use cursor pagination instead.", "x-sunset": "2026-11-01"}], "responses": {"200": {"description": "Search results for the provided query.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_MessageDetails"}}, "result_info": {"$ref": "#/components/schemas/email-security_CursorResultInfo"}}, "required": ["result", "result_info"], "type": "object"}]}}}}, "202": {"description": "The search is taking longer than expected. Use the Location header to poll for results.", "headers": {"Location": {"description": "URL to poll for search results.", "schema": {"type": "string"}}}, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"description": "Always empty on 202; follow the Location header to retrieve results.", "type": "array", "items": {"type": "string"}, "maxItems": 0}, "result_info": {"$ref": "#/components/schemas/email-security_CursorResultInfo"}}, "required": ["result", "result_info"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.investigate", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-stability": "beta"}
```
