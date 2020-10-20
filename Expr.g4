// antlr4 -Dlanguage=Go -o parser Expr.g4

// Expr.g4
grammar Expr;

query
   : NOT? SP? '(' query ')'                                                                                # parenExp
   | query SP LOGICAL_OPERATOR SP query                                                                    # logicalExp
   | attrPath SP op=( EQ | NE | GT | LT | GE | LE | CONTAIN | PREFIX | SUFFIX | IN | REGEX) SP value       # compareExp
   ;

NOT: 'not' | 'NOT';
LOGICAL_OPERATOR: 'and' | 'or';
BOOLEAN: 'true' | 'false';
NULL: 'NULL' | 'null';
IN:  'IN' | 'in';
EQ : '==';
NE : '!=';
GT : '>';
LT : '<';
GE : '>=';
LE : '<=';
CONTAIN : 'contain';
PREFIX : 'prefix';
SUFFIX : 'suffix';
REGEX : 'regex';

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
