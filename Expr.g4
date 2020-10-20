// antlr4 -Dlanguage=Go -o parser Expr.g4

// Expr.g4
grammar Expr;

query
   : NOT? SP? '(' query ')'                                                                                # parenExp
   | query SP LOGICAL_OPERATOR SP query                                                                    # logicalExp
   | attrPath SP op=( EQ | NE | GT | LT | GE | LE | CONTAIN | PREFIX | SUFFIX | IN | REGEX) SP value       # compareExp
   ;



NOT: N O T;
LOGICAL_OPERATOR: A N D | O R;
BOOLEAN: T R U E | F A L S E;
NULL: N U L L;
IN:  I N;
EQ : '==' | E Q;
NE : '!=' | N E;
GT : '>' | G T;
LT : '<' | L T;
GE : '>=' | G E;
LE : '<='| L E;
CONTAIN : C O N T A I N;
PREFIX : P R E F I X;
SUFFIX : S U F F I X;
REGEX : R E G E X;

attrPath
   : ATTRNAME subAttr?
   ;

subAttr
   : '.' attrPath
   ;

ATTRNAME
   : ALPHA ATTR_NAME_CHAR* INDEX?;

fragment ATTR_NAME_CHAR
   : '-' | '_' | ':' | DIGIT | ALPHA
   ;

INDEX
   : '[' DIGIT ']';

fragment DIGIT
   : ('0'..'9')
   ;

fragment ALPHA
   : ( 'A'..'Z' | 'a'..'z' )
   ;

value
   : BOOLEAN           #boolean
   | NULL              #null
   | STRING            #string
   | DOUBLE            #double
   | '-'? INT EXP?     #integer
   | listIntegers      #listOfIntegers
   | listDoubles       #listOfDoubles
   | listStrings       #listOfStrings
   ;

STRING
   : '"' (ESC | ~ ["\\])* '"'
   ;

listStrings
   : '[' subListOfStrings
   ;

subListOfStrings
   : STRING COMMA subListOfStrings
   | STRING ']';

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;

fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;

fragment HEX
   : [0-9a-fA-F]
   ;

DOUBLE
   : '-'? INT '.' [0-9] + EXP?
   ;

listDoubles
   : '[' subListOfDoubles
   ;

subListOfDoubles
   : DOUBLE COMMA subListOfDoubles
   | DOUBLE ']';

listIntegers
   : '[' subListOfIntegers
   ;

subListOfIntegers
   : INT COMMA subListOfIntegers
   | INT ']';

// INT no leading zeros.
INT
   : '0' | [1-9] [0-9]*
   ;

// EXP we use "\-" since "-" means "range" inside [...]
EXP
   : [Ee] [+\-]? INT
   ;

NEWLINE
   : '\n' ;

COMMA
   : ',' ' '*;
SP
   : ' ' NEWLINE*
   ;



fragment A: [aA];
fragment B: [bB];
fragment C: [cC];
fragment D: [dD];
fragment E: [eE];
fragment F: [fF];
fragment G: [gG];
fragment H: [hH];
fragment I: [iI];
fragment J: [jJ];
fragment K: [kK];
fragment L: [lL];
fragment M: [mM];
fragment N: [nN];
fragment O: [oO];
fragment P: [pP];
fragment Q: [qQ];
fragment R: [rR];
fragment S: [sS];
fragment T: [tT];
fragment U: [uU];
fragment V: [vV];
fragment W: [wW];
fragment X: [xX];
fragment Y: [yY];
fragment Z: [zZ];