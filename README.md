#Simple calculator with cli

##Build the calc

- Navigate to the root dir and build the docker image:

        docker build -t calc .


- Run the built image:

       docker run -p 8080:8080 calc

## You can test the calculator with curl requests or use the predefined requests with 'make'
- Add
  
        make add
  

- Subtract
  
        make subtract

- Divide

        make divide

- Multiply

        make multiply

- Validate

        make validate

- Get errors

        make errors

- Running tests

       go test ./...

##Testing the cli
- You have to navigate in the cli folder and execute the binary ./cli which will show the manual

    /.cli