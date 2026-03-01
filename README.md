vero
===========

> [!NOTE]
> Vero is under active development. Before the stable v1.x.x release, there may be many updates.

_Vero_ provides testing utilities and test data generation.

**What problems does _vero_ solve?**

- Intuitive structure - helps you quickly find what you need;
- Simplified fake data generation - works with both built-in and custom types;
- Functions and packages are designed specifically for testing purposes.

## `random` package

Simple generate struct

### Bool type rules

В скобках указаны возможные значения

- `only=(true,false)` - генерация этого поля не будет случайной. В поле поставится то, что указано в правиле.
При указании нескольких `only`
