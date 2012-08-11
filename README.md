## EBNF

```
<command>   -> [id = ] <exp> eol
<exp>       -> <mulexp> { <addop> <mulexp> }
<mulexp>	-> <powexp> { <mulop> <powexp> }
<powexp>	-> <unaryexp> [ <powop> <unaryexp> ]
<unaryexp>	-> [<addop>] <atom>
<atom>		-> (<exp>) | num | id

<addop>	-> + | -
<mulop>	-> * | /
<powop>	-> **
```

## BUILD

```
cd gocalc
go build
```