/*Package fakeit provides a mechanism for faking data for builtin data types and go-openapi/strfmt extended types.

Fakeit can take any struct and populate data for each field. Any exported fields and nested data structures will be populated and will recurse until all exported entities are populated. One limitation of this is that data structures which include an instance of themselves will cause an infinite recursion. This can be solved via Struct Field Tags and is discussed below.

Support for Struct Field Tags allows for the use of specific fakers, as well as setting recursion limits on a Struct Field.

The choice of faker is particularly useful in generation of string and int data types. Out of the box Fakeit supports both github.com/icrowley/fake and Pallinder/go-randomdata. The reason for this is that both have valuable fakers.
*/
package fakeit
