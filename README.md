# quiz


Q: Given a list of words like https://github.com/NodePrime/quiz/blob/master/word.list find the longest compound-word in the list, which is also a concatenation of other sub-words that exist in the list. The program should allow the user to input different data. The finished solution shouldn't take more than one hour. Any programming language can be used, but Go is preferred.


Fork this repo, add your solution and documentation on how to compile and run your solution, and then issue a Pull Request. 

Obviously, we are looking for a fresh solution, not based on others' code.

---

## Test with:
```bash
go test
```

## Run with:
```bash
cd cmd
go build longestCompound.go
cat ../words.list | ./longestCompound --stdin
```

## Cmd line help:
```bash
cd cmd
longestCompound -h
```
