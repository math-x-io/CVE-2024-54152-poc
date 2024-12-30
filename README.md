# ⚙️🔨 IN PROGRESS

## CVE-2024-54152 POC  

Angular Expressions provides expressions for the Angular.JS web framework as a standalone module. Prior to version 1.4.3, an attacker can write a malicious expression that escapes the sandbox to execute arbitrary code on the system. With a more complex (undisclosed) payload, one can get full access to Arbitrary code execution on the system. The problem has been patched in version 1.4.3 of Angular Expressions. Two possible workarounds are available. One may either disable access to `__proto__` globally or make sure that one uses the function with just one argument.

### Proof of Concept

To demonstrate the vulnerability, we have created a simple Node.js application that uses the vulnerable version of Angular Expressions. The application exposes an endpoint `/parse` that evaluates user-provided expressions.

#### Running the Vulnerable Application

1. **Clone the repository**:
    ```sh
    git clone https://github.com/example/CVE-2024-54152-poc.git
    cd CVE-2024-54152-poc
    ```

2. **Build and run the Docker container**:
    ```sh
    docker build -t vulnerable-app .
    docker run -p 8080:8080 vulnerable-app
    ```

3. **Send the malicious payload**:
    You can use either the provided Go or Python script to send the payload.

    - Using Go:
        ```sh
        go run poc.go
        ```

    - Using Python:
        ```sh
        python3 poc.py
        ```

#### Expected Output

The server should execute the payload and return the result of the `id` command, demonstrating arbitrary code execution.

### Mitigation

To mitigate this vulnerability, update Angular Expressions to version 1.4.3 or later. Alternatively, you can disable access to `__proto__` globally or ensure that the function is used with only one argument.

### References

- [Angular Expressions GitHub Repository](https://github.com/peerigon/angular-expressions)
- [CVE-2024-54152 Details](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2024-54152)


