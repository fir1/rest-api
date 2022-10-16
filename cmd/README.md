### /cmd 

Folder will contain main applications for this project.
There might be more than 1 application within **cmd** folder, and each application should have their own folder so the path would be `cmd -> applicationNameFolder`.
The directory name for each application should match the name of executable you wan to have, for example, `cmd/myapp`.

It is important that we don't put a lot of code in the application directory. If you belive that the code can be imported and used in other project, it is better to place these codes in the `pkg` directory.

It is common practice to have a small `main` function that imports and invokes the code from the `/internal` and `pkg` directories.

Please check out some of the successful projects which is written in GoLang.

Examples:
- https://github.com/kubernetes/kubernetes/tree/master/cmd
- https://github.com/ethereum/go-ethereum/tree/master/cmd
- https://github.com/influxdata/influxdb/tree/master/cmd
