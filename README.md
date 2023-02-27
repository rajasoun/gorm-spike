# About 

GORM spike to test connection poll and read/write split

# How to run

1. Start mysql server
    ```bash
    cd mysql-writer-reader 
    ./assist.sh up
    cd - 
    ```

2. Run the test
    ```bash
    cd gorm-client
    go test -v
    ```
