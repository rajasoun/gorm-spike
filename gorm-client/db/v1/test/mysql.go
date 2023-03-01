package test

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	dockertest "github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func startMySqlContainer(pool *dockertest.Pool, exposedPort string) (*dockertest.Resource, error) {
	var passwordEnv = "MYSQL_ROOT_PASSWORD=dbpassword"

	return pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "latest",
		Env: []string{
			fmt.Sprint(passwordEnv),
		},

		ExposedPorts: []string{exposedPort},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"3306/tcp": {
				{HostIP: "0.0.0.0", HostPort: exposedPort},
			},
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
}

func prepareDBContainer(db *sql.DB) error {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS dbname")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE USER 'dbuser'@'%' IDENTIFIED BY 'dbpassword'")
	if err != nil {
		return err
	}

	_, err = db.Exec("GRANT ALL PRIVILEGES ON dbname.* TO 'dbuser'@'%'")
	if err != nil {
		return err
	}
	return nil
}

func waitForMySql(pool *dockertest.Pool, exposedPort string) error {
	pool.MaxWait = 120 * time.Second
	return pool.Retry(func() error {
		db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(localhost:%s)/mysql", "dbpassword", exposedPort))
		if err != nil {
			return err
		}
		// Create a new database user with the necessary privileges
		err = prepareDBContainer(db)
		if err != nil {
			return err
		}
		return db.Ping()
	})
}

func newDockerPool() *dockertest.Pool {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	log.Println("Connected to docker")
	return pool
}

func startContainer(pool *dockertest.Pool, exposedPort string) *dockertest.Resource {
	resource, err := startMySqlContainer(pool, exposedPort)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	log.Println("Started resource")
	return resource
}

func setExpiration(resource *dockertest.Resource) {
	if err := resource.Expire(120); err != nil {
		log.Fatalf("Container Expiration Error: %s", err)
	}
	log.Println("Container Expiration Setting Passed")
}

func waitForContainerToStart(pool *dockertest.Pool, exposedPort string) {
	if err := waitForMySql(pool, exposedPort); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	log.Println("Connected to docker")
}

// InitTestDocker function initialize docker with postgres image used for integration tests
func InitTestDocker(exposedPort string) (*dockertest.Pool, *dockertest.Resource) {
	pool := newDockerPool()
	resource := startContainer(pool, exposedPort)
	setExpiration(resource)
	waitForContainerToStart(pool, exposedPort)
	return pool, resource
}
