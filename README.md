# NeverSitUpTest

## Golang
### how to run (ACK- that you already install Golang 1.22.2)
```bash
> cd tests/
> tests/: go test -v
```

*ex : response
```bash
➜  tests git:(master) ✗ go test -v
=== RUN   TestFindOddInt
--- PASS: TestFindOddInt (0.00s)
=== RUN   TestPermutationsFn
--- PASS: TestPermutationsFn (0.00s)
=== RUN   TestCountSmileys
--- PASS: TestCountSmileys (0.00s)
PASS
ok      not_workout/tests       0.001s
```

## JS
### how to run (ACK- that you already install Node 18.19.1)
```bash
> cd javascript/
> javascript/: node oddInt.js
> javascript/: node permut.js
> javascript/: node smiley.js
```

*ex : response
```bash
➜  javascript git:(master) ✗ node permut.js
[ 'aabb', 'abab', 'abba', 'baab', 'baba', 'bbaa' ]
➜  javascript git:(master) ✗ node oddInt.js 
Odd integer: 3
➜  javascript git:(master) ✗ node smiley.js 
Count of smiley faces: 2
```

## TS
* Same with JS
