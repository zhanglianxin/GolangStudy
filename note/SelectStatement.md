The syntax is as follows:

```go
select {
    case communication clause:
        statement(s);
    case communication clause:
        statement(s);
    default:
        statement(s);
}
```

The following rules apply to a select statement:

* You can have any number of case statements within a select. Each case is followed by the value to be compared to and a colon.

* The type for a case must be the communication channel operation.


* When the channel operation occurred the statements following that case will execute. <u>No **break** is needed in the case statement</u>.

* A select statement can have an optional default case, which must appear at the end of the select. The default case can be used for performing a task when none of the cases is true. <u>No **break** is needed in the default case</u>.
