# gRPC Service in Go

A gRPC Service written in Go and uses protobufs.

## Machine Used

The code was tested to work on the following machine:

```sh
OS: Arch Linux x86_64
Host: Inspiron 5570
Kernel: 6.9.8-arch1-1
CPU: Intel i5-8250U (8) @ 3.400GHz
```

## Dependencies

- go
- protoc-gen-go
- protoc-gen-go-grpc
- protoc

## Steps to Run

1. Clone the repo

    ```sh
    git clone --depth 1 https://github.com/swarnimcodes/go-grpc-tc.git && cd go-grpc-tc
    ```

2. Install project dependencies

    ```sh
    go mod tidy
    ```

3. Compile the user.proto file and generate corresponding go code

    ```sh
    protoc --go_out=. --go-grpc_out=. user.proto
    ```

4. Run the server

    ```sh
    # run conventionally:
    go run cmd/server/server.go

    # run using live reload
    ## this uses the unix utility called `entr`
    echo "cmd/server/server.go" | entr -r go run cmd/server/server.go 
    ## using this way you can make changes to the server file
    ## and go will automatically compile and run your program
    ## no fancy tools, just in-built utilities to make live reload work
    ```

5. Run the tests

    ```sh
    # run all the test files located anywhere inside the project folder
    go test ./...
    ```

6. Run the client

    ```sh
    go run cmd/client/client/go
    # you can use `entr` here as well.
    # that way you can make changes to the client and
    # get live feedback.
    ```

7. Precautions
    - make sure to re-generate the go code after making changes to the user.proto file
    - rarely the client receives a failure message,  re-run the client/tests and/or increase the timeout and re-run again

8. Vim Sessions
    A session.vim file is provided that includes the setup to code and run the server and the client/tests.
    To run vim and set up the session, run the following command:

    ```sh
    nvim -S ./session.vim # if you use neovim
    vim -S ./session.vim  # if you use vim
    ```

## gRPC Endpoints

1. `GetUserById`
   Input Params: Id [int32]

   Output: User that matches the given Id
2. `GetUsersByIds`
   Input Params: Ids []int32

   Output: All the users that exist with the given Ids from the list
3. `SearchUsers`
   Input Params:
   - SearchByPhoneNumber [bool]: Compulsory param. Tells the server if we are going to search for phone numbers.
   - SearchByMarriageStatus [bool]: Compulsory param. Tells the server if we are going to search for marital status.
   - Phone [int64]: Optional. Necessary when SearchByPhoneNumber is true.
   - Married [bool]: Optional. Necessary when SearchByMarriageStatus is true.
   - City [string]: Optional
   - Fname [string]: Optional

   Output: All the users that match the various combinations formed via the input request.
   Only those users are shown that actually match all the search criteria.
