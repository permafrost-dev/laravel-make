package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func getDockerComposeTemplate() string {
	template := `version: '3'

services:
    mysql:
        image: docker.io/mariadb:11
        container_name: {{project.name}}-mysql
        environment:
            MYSQL_ROOT_PASSWORD: '${DB_PASSWORD}'
            MYSQL_DATABASE: '${DB_DATABASE}'
            MYSQL_USER: '${DB_USERNAME}'
            MYSQL_PASSWORD: '${DB_PASSWORD}'
            MYSQL_ALLOW_EMPTY_PASSWORD: 'yes'
        volumes:
            - '{{project.name}}mysql:/var/lib/mysql'
        ports:
            - '${FORWARD_DB_PORT:-3306}:3306'
        networks:
            - {{project.name}}net
        command: '--default-authentication-plugin=mysql_native_password --thread_handling=pool-of-threads --thread-pool-size=128 --thread-pool-dedicated-listener=1 --thread-pool-priority=high  --thread-pool-stall-limit=5'
        restart: always
        healthcheck:
            test: ['CMD', 'mysqladmin', 'ping']
        expose:
            - 3306

    redis:
        image: docker.io/redis:alpine
        container_name: {{project.name}}-redis
        ports:
            - '${REDIS_PORT:-6379}:6379'
        networks:
            - {{project.name}}net


networks:
    {{project.name}}net:

volumes:
    {{project.name}}mysql:`

	template = strings.ReplaceAll(template, "{{project.name}}", os.Args[1])

	return template
}

func createDockerComposeFile() error {
	f, err := os.Create("docker-compose.yml")

	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.WriteString(getDockerComposeTemplate()); err != nil {
		return err
	}

	return nil
}

func updateDotEnvFile() error {
	// read file content, replace a string, and write it back to file

	// open file
	f, err := os.Open(".env")

	if err != nil {
		return err
	}

	defer f.Close()

	// read file content
	b, err := io.ReadAll(f)

	if err != nil {
		return err
	}

	content := string(b)
	content = strings.ReplaceAll(content, "APP_NAME=Laravel", "APP_NAME="+os.Args[1])
	content = strings.ReplaceAll(content, "QUEUE_CONNECTION=sync", "QUEUE_CONNECTION=redis")

	fw, err := os.Create(".env")

	if err != nil {
		return err
	}

	defer fw.Close()

	if _, err := fw.WriteString(content); err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: laravel-make <project-name>")
		os.Exit(1)
	}

	cmd := exec.Command("laravel", "new", os.Args[1])

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Chdir(os.Args[1])

	if err := createDockerComposeFile(); err != nil {
		fmt.Println(err)
		fmt.Println("Failed to create docker-compose.yml file.")
	}

	if err := updateDotEnvFile(); err != nil {
		fmt.Println(err)
		fmt.Println("Failed to update .env file.")
	}

	fmt.Println("Done!")
}
