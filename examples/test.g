TNAME = "[A-Z_]+";
RNAME = "[a-z_]+";

REGEX = "\".*\"";

EQ = "=";
BAR = "|";
SCOLON = ";";

WS = "[ \t\n]+";

grammar = token_def grammar 
    | rule_def grammar 
    | token_def 
    | rule_def 
    ;
token_def = TNAME EQ REGEX SCOLON;
rule_def = RNAME EQ rule_set SCOLON;
rule_set = TNAME rule_set
    | RNAME rule_set
    | BAR rule_set 
    | TNAME 
    | RNAME
    ;