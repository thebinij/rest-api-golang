## Initialize GO APP

go mod init <module_name>

## Install Fiber V2

go get -u <package_name>

## Uninstall the package

go list -m -f={{.Dir}} <package_name>
rm -rf <package_name>
