# SaaS-CLI
SaaS-CLI is a utility that interacts with a Palo Alto Networks Firewall to pull the [SaaS Application Usage Report](https://docs.paloaltonetworks.com/pan-os/10-1/pan-os-admin/monitoring/view-and-manage-reports/view-reports.html#id12e5da3d-d44f-4c9e-9d97-8704151ed103) and outputs all the discovered SaaS applications which you can then review and tag them as Sanctioned if applicable.

## Files contained in this folder:
* cmd - Contains the application entry point code
* pkg - Contains multiple packages used by the application (The secret package is not included in this repository for security reasons).
* saascli-macos - Executable for MacOS systems
* saascli-linux - Executable for Linux systems
* saascli-windows - Executable for Windows systems

## How to use the program
* For Windows users download the "saascli-windows" executable and rename it by adding a ".exe" extension
* For Linux users download the "saascli-linux" executable
* Browse to the folder where the application resides and run the executable
* Windows run "saas-cli-windows.exe"
* Linux run "./saas-cli-linux"

## Encryption/Decryption Key Details
The executable is packaged with a 32-byte array used to encrypt and decrypt the API key. If you want to change this key and recompile the executable use the following steps.
*[Install Go](https://go.dev/doc/install) (if you don't have it already)
* Clone the repository
* Create a secret folder under pkg
* Create a secret.go file inside the secret folder
* Create a function that returns a 32-byte array. Example below:
```
package examplesecret

// Secret: returns a 32 byte array used as the encryption and decryption key. This can be any 32byte array.
func Secret() []byte {
	return []byte("thisshouldbeany32bytearrayyoudig") // CHANGE THIS VALUE TO YOUR OWN VALUE
}
```
