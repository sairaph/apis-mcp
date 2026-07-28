---
title: tax_product_resource_tax_association_transaction_attempts_resource_errored
page_id: schema-tax-product-resource-tax-association-transaction-attempts-resource-error-150110f4
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_resource_tax_association_transaction_attempts_resource_errored

```yaml
{"title": "TaxProductResourceTaxAssociationTransactionAttemptsResourceErrored", "required": ["reason"], "type": "object", "properties": {"reason": {"type": "string", "description": "Details on why we couldn't commit the tax transaction.", "enum": ["another_payment_associated_with_calculation", "calculation_expired", "currency_mismatch", "original_transaction_voided", "unique_reference_violation"]}}, "description": "", "x-expandableFields": []}
```
