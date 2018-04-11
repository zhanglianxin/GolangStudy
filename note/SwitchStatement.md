In Go programming, switch statements are of two types:

* **Expression Switch**

  In expression switch, a case contains expression, which is compared against the value of the switch expression.

  The syntax is as follows:

  ```go
  switch (boolean-expression or integral type) {
      case boolean-expression or integral type:
          statement(s);
      case boolean-expression or integral type:
          statement(s);
      default:
          statement(s);
  }
  ```

  The following rules apply to a switch statement:

  * The expression used in a switch statement must have an integral or boolean expression, or be of a class type in which the class has a single conversion function to an integral or boolean value. If the expression is not passed then the default value is `true`.

  * You can have any number of case statements within a switch. Each case is followed by the value to be compared to and a colon.

  * The **constant-expression** for a case must be the same data type as the variable in the switch, and it must be a constant of a literal.

  * When the variable being switched on is equal to a case, the statements following that case will execute. <u>No **break** is needed in the case statement</u>.

  * A switch statement can have an optional default case, which must appear at the end of the switch. The default case can be used for performing a task when none of the cases is true. <u>No **break** is needed in the default case</u>.

* **Type Switch**

  In type switch, a case contains type, which is compared against the type of a specially annotated switch expression.

  The syntax is as follows:

  ```go
  switch x.(type) {
      case type:
          statement(s);
      case type:
          statement(s);
      default:
          statement(s);
  }
  ```

  The following rules apply to a switch statement:

  * The expression used in a switch statement must have an variable of interface{} type.

  * You can have any number of case statements within a switch. Each case is followed by the value to be compared to and a colon.

  * The type for a case must be the same data type as the variable in the switch, and it must be a valid data type.

  * When the variable being switched on is equal to a case, the statements following that case will execute. <u>No break is needed in the case statement</u>.

  * A switch statement can have an optional default case, which must appear at the end of the switch. The default case can be used for performing a task when none of the cases is true. <u>No break is needed in the default case</u>.
