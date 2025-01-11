# gomni ♾️

`gomni` is a command-line tool for routine work, it is very ease to learn and use! 

`gomni` mean base on `golang` and `omni` for "all"♾️

## Features

### v0.1.0: SSH Client

+ **Connect to SSH remotes** using either hostnames or IP addresses.
+ **Add new SSH remotes** with support for both password and certificate-based authentication.
+ **Delete SSH remotes** by hostname or IP address.
+ **List all available remotes** with their status (online/offline).

## Installation

You can use the quick install script with

```bash
curl -O https://github.com/linyejoe2/gomni/releases/download/v0.1.2/gomni-v0.1.2.tar.gz
curl -o- https://github.com/linyejoe2/gomni/releases/download/v0.1.2/install-completion.sh | bash
# chmod +x install-completion.sh
# ./install-completion.sh
```

Or you can install `gomni` directly by download tarball or zip file from [source](https://github.com/linyejoe2/gomni/releases) page.

## Usage

### SSH Commands

The `ssh` command provides a set of tools for managing your SSH connections.

#### 1. Connect to SSH Remote

You can connect to a remote using its hostname or IP address.

```bash
gomni ssh [host name | ip]
```

💡 Example: Assume you have the following remote list:

NAME|	IP|	USERNAME	|STATUS
-|-|-|-
my-pc|	127.0.0.1	|user|	online

You can connect to my-pc by running either of the following commands:

```bash
gomni ssh my-pc
gomni ssh 127.0.0.1
```

#### 2. Add a New SSH Remote

To add a new SSH remote, you can specify the IP address, hostname, username, and either a password or a certificate file.

```bash
gomni ssh add <ip> -n <hostname> -u <username> -p <password> or -i <certificate file>
```

💡 Example:

+ With password:
```bash
gomni ssh add 127.0.0.1 -n my-pc -u user -p 1234
```
+ With certificate file:
```
gomni ssh add 127.0.0.1 -n my-pc -u user -i ~/.ssh/id_rsa
```

❗ Note:

+ Avoid using reserved words like `add`, `delete`, or `list` as hostname, as they will cause errors.
+ If using a certificate, ensure that the public key is added to the server beforehand.
+ The certificate file refers to the private key, not the public key.

#### 3. Delete an SSH Remote

To delete an SSH remote, use either the hostname or the IP address.

```
gomni ssh delete <remote name | ip>
```

💡 Example:

```
gomni ssh delete my-pc
gomni ssh delete 127.0.0.1
```

#### 4. List Available Remotes

List all available SSH remotes and their status.

```
gomni ssh list
```

#### Flags

All commands support the -h or --help flag to display usage instructions.
