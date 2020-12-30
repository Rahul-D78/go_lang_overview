                    control Structure
                           |
        ___________________|__________________
       |                   |                  |
       |                   |                  |
       |                   |                  |
       |                   |                  |
     Defer              Recover             Panic

-----Defer----     
. Defer statement defers the executation of a function until   the surrounding function returns.
. Multiple defers are pushed into stack and executes in last in first out order.
>Defer generally used to cleanup resouces like file , db connections etc.

-----Recover----
. Recover is another builtin function in go 
