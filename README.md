# trash
A CLI tool for managing trash files 



## install
```bash
go install github.com/Beriholic/th@latest
```

## usage
### push file to trash
```bash
th push <file>
th push <file> <file> <file> ...
```
>you can use ```p```, ```ps```, ```rm```, ```remove``` instead of ```push```
### clean trash
```bash
th clean
```
### list trash files
```bash
th list
```
it will list all files in trash

like this:
```
+-------------------------------------------------------+
| Trash List                                            |
+----+-----------+----------------+---------------------+
| ID | NAME      | PATH           | DELETIONDATE        |
+----+-----------+----------------+---------------------+
|  0 | 1.txt     | /tmp/1.txt     | 2023-11-18 22:23:06 |
|  1 | 2.txt     | /tmp/2.txt     | 2023-11-18 22:23:06 |
|  2 | 3.txt     | /tmp/3.txt     | 2023-11-18 22:23:06 |
|  3 | test_file | /tmp/test_file | 2023-11-18 22:23:06 |
+----+-----------+----------------+---------------------+
```
> you can use ```ls```, ```ll```, ```l``` instead of ```list```

### restore file from trash
```bash
th restore <id>
th restore <id> <id> <id> ...

# Use an interactive menu
th restore
``` 
> you can use ```r```, ```rs```, ```res``` instead of ```restore```
 
 ### delete file from trash
 ```bash
 th del <id>
 th del <id> <id> <id> ...

# Use an interactive menu
 th del
 ```
 > you can use ```d``` instead of ```del```



## TODO
- [x] ```push``` --> push file to trash
- [x] ```clean``` --> clean trash
- [x] ```list``` --> list trash files
- [x] ```restore``` --> restore file from trash
    - [x]  use ```[start-end]``` to restore a range of files under the interactive menu
- [x] ```del``` --> delete file from trash
    - [x] use ```[start-end]``` to delete a range of files run under the interactive menu
- [ ] ```find``` --> find file in trash
- [ ] ```ui``` --> a tui for trash
