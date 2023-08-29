## Class file structure

| type           | name              | count                 |
|----------------|-------------------|-----------------------|
| u4             | magic number      | 1                     |
| u2             | minor version     | 1                     |
| u2             | major version     | 1                     |
| cp_info        | constant pool     | constant_pool_count-1 |
| u2             | access flags      | 1                     |
| u2             | this class index  | 1                     |
| u2             | super class index | 1                     |
| u2             | interfaces count  | 1                     |
| u2             | interfaces        | interfaces count      |
| u2             | field_count       | 1                     |
| field_info     | fields            | field count           |
| u2             | methods count     | 1                     |
| method_info    | methods           | methods_count         |
| u2             | attributes count  | 1                     |
| attribute_info | attribute         | attribute count       |
