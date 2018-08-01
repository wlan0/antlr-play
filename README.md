# antlr-play

Antlr pretty printing example for the Go runtime

To use it

```bash
$ go get github.com/wlan0/antlr-play && cd $GOPATH/src/github.com/wlan0/antlr-play
$ ./build.sh Test.G4 
$ ./antlr-play input
(expression 
 (expression 
  (expression 
   (expression  'a1')
   '*'
   (expression  '('
    (expression 
     (expression 
      (expression 
       (expression  'a2')
       '+'
       (expression  'a3')
      )
      '+'
      (expression  'a4')
     )
     '-'
     (expression 
      (expression  'a5')
      '^'
      (expression 
       (expression  'a6')
       '^'
       (expression  '('
        (expression 
         (expression  'a8')
         '+'
         (expression 
          (expression  'a9')
          '/'
          (expression  'a10')
         )
        )
        ')')
      )
     )
    )
    ')')
  )
  '*'
  (expression  'a11')
 )
 '*'
 (expression  'a12')
)
```
