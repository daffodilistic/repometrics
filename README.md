# repometrics
A sample Golang program to generate metrics from a repository.

## How to run/install
1. Install Golang (v1.18.1 and above). This program was developed on a Windows 
machine, but it should not matter much.
2. Compile the program with `go build`
3. Copy the binary file to a directory of your choice
(typically `/usr/local/bin`).
4. Ensure permissions is set correctly (typically umask `755`) and system `PATH`
variable is set correctly.
5. Run the program with `repometrics <REPOSITORY_DIRECTORY>`
