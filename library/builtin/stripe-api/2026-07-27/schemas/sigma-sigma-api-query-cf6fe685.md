---
title: sigma.sigma_api_query
page_id: schema-sigma-sigma-api-query-cf6fe685
path: schemas
description: A saved query object represents a query that can be executed for a run.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# sigma.sigma_api_query

A saved query object represents a query that can be executed for a run.

```yaml
{"title": "SigmaSigmaResourcesSigmaAPIQuery", "required": ["created", "id", "livemode", "name", "object", "sql"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch."}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "name": {"maxLength": 5000, "type": "string", "description": "The name of the query."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["sigma.sigma_api_query"]}, "sql": {"maxLength": 5000, "type": "string", "description": "The sql statement for the query."}}, "description": "A saved query object represents a query that can be executed for a run.", "x-expandableFields": [], "x-resourceId": "sigma.sigma_api_query"}
```
