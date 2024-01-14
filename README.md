# Email Domain Checker 🌐

This Go program checks email domain information including MX records, SPF records, and DMARC records. It prompts the user to input a domain and outputs relevant information.

## Usage 🚀

Run the program:

```bash
go run main.go
```

Enter an email domain (e.g., google.com) when prompted, and the program will display relevant information about the email domain.

## Code Explanation 📜

### Main Program (main.go) 🧑‍💻

- The program uses the net package to perform DNS lookups for MX, TXT, and DMARC records.
- User input is obtained using bufio.Scanner reading from os.Stdin.
- The checkDomain function processes the domain information and prints the results.

### checkDomain Function 🕵️‍♂️

- The function looks up MX records to check if the domain has mail servers.
- It looks up TXT records to find SPF (Sender Policy Framework) records.
- It looks up DMARC records by appending "\_dmarc." to the domain and checking TXT records.
- The results are printed to the console in a CSV-like format.

## Closing Notes 📝

If you find any issues or have suggestions for improvement, please feel free to open an issue.

Happy coding! 🚀👨‍💻
