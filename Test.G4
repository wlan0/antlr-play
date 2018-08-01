grammar Test;

expression
    : <assoc=right> expression '^' expression
    | expression ('*' | '/') expression
    | expression ('+' | '-') expression
    | '(' expression ')'
    | ID
    ;
    
ID:     [a-zA-Z] [0-9a-zA-Z]*;
DIGIT:  [0-9]+;
WS:     [ \t\r\n] -> skip;
