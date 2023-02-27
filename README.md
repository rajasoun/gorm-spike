# About 

GORM is an ORM library for Go that supports connection pooling, read/write splitting, and other features that can help improve the performance and scalability of Go applications.

GORM spike to test connection poll and read/write split

> **Connection pooling** is a technique used to manage a pool of database connections and reuse them instead of creating new connections every time an application needs to interact with the database. This can help reduce the overhead of creating new connections and improve the performance of database operations.

## GORM Connection Pooling

In GORM, connection pooling works by maintaining a pool of database connections that can be reused by multiple goroutines.  

1. When an application needs to perform a database operation, it requests a connection from the pool.
2. If there is an available connection in the pool, it is returned to the application. Otherwise, a new connection is created and added to the pool.
3. The size of the connection pool is configurable in GORM, and it can be set to a specific number of connections or allowed to grow dynamically as needed. 
4. When a connection is returned to the pool, it is checked to make sure that it is still valid and usable. If a connection is found to be invalid or has been idle for too long, it is discarded and replaced with a new connection.
5. GORM's connection pooling also includes features like connection timeouts, which allow connections to be closed after a certain period of inactivity, and connection reaping, which periodically checks connections in the pool to make sure they are still valid and removes any that are not.

## GORM Read/Write Splitting

Read/write splitting can help improve the performance of applications that have a high volume of read operations by allowing them to be distributed across multiple read-only database replicas.

GORM also supports read/write splitting, which is a technique used to divide database operations into read and write operations and direct them to different database servers.

1. GORM's read/write splitting works by configuring multiple database connections, with some connections designated for read operations and others designated for write operations.
2. When an application needs to perform a database operation, GORM selects the appropriate connection based on the type of operation.
3. Read operations are distributed across the read connections, while write operations are sent to the write connection(s).
4. GORM's read/write splitting also includes features like load balancing, which distributes read operations across multiple read connections to help evenly distribute the workload.

> Note that read/write splitting may not always be appropriate for every application, as it can add complexity and may not provide a significant performance boost for applications with low read/write ratios.

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
    go test -count=1  -v -run TestCreateUsersWithConnectionPool  github.com/rajasoun/gorm-client/test
    go test -count=1  -v -run TestCreateUsersWithoutConnectionPool  github.com/rajasoun/gorm-client/test
    ```
