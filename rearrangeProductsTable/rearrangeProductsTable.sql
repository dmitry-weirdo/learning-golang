-- 1795. Rearrange Products Table
-- todo: this can be also solved with UNPIVOT

(
    select
    p.product_id,
    'store1' as "store",
    store1 as "price"
    from Products p
    where (p.store1 is not null)
)
union all
(
    select
    p.product_id,
    'store2' as "store",
    store2 as "price"
    from Products p
    where (p.store2 is not null)
)
union all
(
    select
    p.product_id,
    'store3' as "store",
    store3 as "price"
    from Products p
    where (p.store3 is not null)
)
