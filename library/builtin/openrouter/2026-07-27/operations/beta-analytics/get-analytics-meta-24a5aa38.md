---
title: Get available analytics metrics and dimensions
page_id: operation-get-analytics-meta-f571d5d6
path: operations/beta-analytics
description: Returns the available metrics, dimensions, filter operators, and granularities for the analytics query endpoint. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /analytics/meta
operation_ids:
    - getAnalyticsMeta
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Get available analytics metrics and dimensions

`GET /analytics/meta`

Operation ID: `getAnalyticsMeta`

Returns the available metrics, dimensions, filter operators, and granularities for the analytics query endpoint. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Returns the available metrics, dimensions, filter operators, and granularities for the analytics query endpoint. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "getAnalyticsMeta", "responses": {"200": {"content": {"application/json": {"example": {"data": {"dimensions": [{"display_label": "Model", "name": "model"}], "granularities": [{"display_label": "Day", "name": "day"}], "metrics": [{"display_format": "number", "display_label": "Request Count", "is_rate": false, "name": "request_count"}], "operators": [{"name": "eq", "value_type": "scalar"}]}}, "schema": {"properties": {"data": {"properties": {"dimensions": {"items": {"properties": {"display_label": {"description": "Human-readable label", "example": "Model", "type": "string"}, "name": {"description": "Dimension identifier used in query requests", "example": "model", "type": "string"}}, "required": ["name", "display_label"], "type": "object"}, "type": "array"}, "granularities": {"items": {"properties": {"display_label": {"description": "Human-readable label", "example": "Day", "type": "string"}, "name": {"description": "Granularity identifier", "enum": ["minute", "hour", "day", "week", "month"], "example": "day", "type": "string", "x-speakeasy-unknown-values": "allow"}}, "required": ["name", "display_label"], "type": "object"}, "type": "array"}, "metrics": {"items": {"properties": {"display_format": {"description": "How this metric value should be formatted for display (e.g. percent → multiply by 100 and append %, currency → prefix with $)", "enum": ["number", "currency", "percent", "latency", "throughput"], "example": "number", "type": "string", "x-speakeasy-unknown-values": "allow"}, "display_label": {"description": "Human-readable label", "example": "Request Count", "type": "string"}, "is_rate": {"description": "Whether this metric is a rate/ratio (averaged, not summed)", "type": "boolean"}, "name": {"description": "Metric identifier used in query requests", "example": "request_count", "type": "string"}}, "required": ["name", "display_label", "is_rate", "display_format"], "type": "object"}, "type": "array"}, "operators": {"items": {"properties": {"name": {"description": "Operator identifier used in filter definitions", "enum": ["eq", "neq", "in", "not_in", "gt", "gte", "lt", "lte"], "example": "eq", "type": "string", "x-speakeasy-unknown-values": "allow"}, "value_type": {"description": "Whether the operator expects a single value or an array", "enum": ["scalar", "array"], "type": "string", "x-speakeasy-unknown-values": "allow"}}, "required": ["name", "value_type"], "type": "object"}, "type": "array"}}, "required": ["metrics", "dimensions", "operators", "granularities"], "type": "object"}}, "required": ["data"], "type": "object"}}}, "description": "Returns analytics query metadata"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Get available analytics metrics and dimensions", "tags": ["beta.Analytics"]}
```
