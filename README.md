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

1. In GORM, read/write splitting is achieved by using a plugin called gorm.io/plugin/dbresolver which supports different types of database replication schemes like Master-Slave, Master-Master, and Sharding.
2. The dbresolver plugin can be configured to route read operations to read-only database replicas while directing write operations to the primary write database.
3. GORM's read/write splitting feature can help improve the scalability and performance of applications by distributing database operations across multiple servers and reducing the load on the primary write database.

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
