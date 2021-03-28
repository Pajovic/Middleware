# Middleware

Coding challenge project. Gathers information about tournaments and matches and returns sorted last N played matches per tournament

# Requirements and setup

Middleware can be used used in 2 ways:
1. Docker - required Docker installation. New image can be created with make deploy ( default port mapping is 8080:8080 - change it in Makefile if needed ). To run a new container instance execute make run. For developing inside a container, create an image an open the project with Visual Studio Code, then reopen in container. In it, install all necessary or desired plugins through Visual Studio Code. URL: 0.0.0.0:8080/tournaments/topMatches/sid/rsid
2. Without Docker - required Golang installation. In Makefile there are make for build, make clean for deleting binary and make test for all package test execution. All these commands can be used inside a container. URL: 127.0.0.1:8080/tournaments/topMatches/sid/rsid
