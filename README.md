Package logger

-   This project pushed to github repository and tried to import the package in the dependency project
Problem -1
    go get "<pkg name>" - by default installs in the $GoPath/pkg/mod and $GoPath/pkg/cache
    The dependency project always refers the imported module in the $GoPath/src/<pkg name> folder
    To achieve follow the below steps
	    go env -w GO111MODULE="off"  
    After that execute 'go get' ... will install packages to $GOPATH/src.
        go get -u <pkg-name>
	To re-enable it do 
	    go env -w GO111MODULE="" 
    After re-enabled then again execute the command
        go get -u <pkg-name>
    This time it updates the go.mod in the dependency project

    As the same source downloaded in the workspace pkg/mod and pkg/mod/cache can be removed manually. 


Problem -2 
-   When the go workspace has the projects that are partially in github and partially not in github will throw exception like below when tried to execute the 'go install' command for a specific 
    project that is not part of github
        $ git status - To view the error message 
    fatal: detected dubious ownership in repository at 'C:/'
    To add an exception for this directory, call:
        git config --global --add safe.directory C:/
    Set the environment variable 
        GIT_TEST_DEBUG_UNSAFE_DIRECTORIES=true 
    and run again for more information.

    git config --global --add safe.directory [specific directory]