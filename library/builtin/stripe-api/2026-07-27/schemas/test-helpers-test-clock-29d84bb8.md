---
title: test_helpers.test_clock
page_id: schema-test-helpers-test-clock-29d84bb8
path: schemas
description: |-
    A test clock enables deterministic control over objects in testmode. With a test clock, you can create
    objects at a frozen time in the past or future, and advance to a specific future time to observe webhooks and state changes. After the clock advances,
    you can either validate the current state of your scenario (and test your assumptions), change the current state of your scenario (and test more complex scenarios), or keep advancing forward in time.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# test_helpers.test_clock

A test clock enables deterministic control over objects in testmode. With a test clock, you can create
objects at a frozen time in the past or future, and advance to a specific future time to observe webhooks and state changes. After the clock advances,
you can either validate the current state of your scenario (and test your assumptions), change the current state of your scenario (and test more complex scenarios), or keep advancing forward in time.

```yaml
{"title": "TestClock", "required": ["created", "deletes_after", "frozen_time", "id", "livemode", "object", "status", "status_details"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "deletes_after": {"type": "integer", "description": "Time at which this clock is scheduled to auto delete.", "format": "unix-time"}, "frozen_time": {"type": "integer", "description": "Time at which all objects belonging to this clock are frozen.", "format": "unix-time"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "name": {"maxLength": 5000, "type": "string", "description": "The custom name supplied at creation.", "nullable": true}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["test_helpers.test_clock"]}, "status": {"type": "string", "description": "The status of the Test Clock.", "enum": ["advancing", "internal_failure", "ready"]}, "status_details": {"$ref": "#/components/schemas/billing_clocks_resource_status_details_status_details"}}, "description": "A test clock enables deterministic control over objects in testmode. With a test clock, you can create\nobjects at a frozen time in the past or future, and advance to a specific future time to observe webhooks and state changes. After the clock advances,\nyou can either validate the current state of your scenario (and test your assumptions), change the current state of your scenario (and test more complex scenarios), or keep advancing forward in time.", "x-expandableFields": ["status_details"], "x-resourceId": "test_helpers.test_clock"}
```
