---
title: workers-observability_filter_leaf
page_id: schema-workers-observability-filter-leaf-83e4c41a
path: schemas
description: A filter condition applied to query results. Use the keys and values endpoints to discover available fields and their values before constructing filters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers-observability_filter_leaf

A filter condition applied to query results. Use the keys and values endpoints to discover available fields and their values before constructing filters.

```yaml
{"description": "A filter condition applied to query results. Use the keys and values endpoints to discover available fields and their values before constructing filters.", "type": "object", "properties": {"key": {"description": "Filter field name. Use verified keys from previous query results or the keys endpoint. Common keys include $metadata.service, $metadata.origin, $metadata.trigger, $metadata.message, and $metadata.error.", "type": "string"}, "kind": {"description": "Discriminator for leaf filter nodes. Always 'filter' when present; may be omitted.", "type": "string", "enum": ["filter"]}, "operation": {"description": "Comparison operator. String operators: includes, not_includes, starts_with, ends_with, regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.", "type": "string", "enum": ["includes", "not_includes", "starts_with", "ends_with", "regex", "exists", "is_null", "in", "not_in", "eq", "neq", "gt", "gte", "lt", "lte", "=", "!=", ">", ">=", "<", "<=", "INCLUDES", "DOES_NOT_INCLUDE", "MATCH_REGEX", "EXISTS", "DOES_NOT_EXIST", "IN", "NOT_IN", "STARTS_WITH", "ENDS_WITH"]}, "type": {"description": "Data type of the filter field. Must match the actual type of the key being filtered.", "type": "string", "enum": ["string", "number", "boolean"]}, "value": {"description": "Comparison value. Must match actual values in your data — verify with the values endpoint. Ensure the value type (string/number/boolean) matches the field type. String comparisons are case-sensitive. Regex uses RE2 syntax (no lookaheads/lookbehinds).", "anyOf": [{"type": "string"}, {"type": "number"}, {"type": "boolean"}]}}, "required": ["key", "operation", "type"]}
```
