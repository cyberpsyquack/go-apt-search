# go-apt-search

Utility to collect the information of packages available in own repository.
It is currently a primitive and raw version, and I wrote it quickly as I could not find any modules of this type to include in another project of mine (currently in private development).
I will, however, be taking it forward in the next few days, at the same time as my other project which I plan to make public in a first version soon after the summer, in order to make it better in both code and functionality.

For anyone who happens to come here and finds this module useful and would like to help me, any contribution is welcome.

## How it works

**go-apt-search** relies on files in the `/var/lib/apt/lists` path, these files are used to store information about each package resource specified in the system *sources.list*.

Specifically, files from which it extracts information are those in the format *<source_of_repository>_Packages*

Below is an example of how a package, with all its information, is shown in these files

```
Package: ctdb
Source: samba
Version: 2:4.17.9+dfsg-0+deb12u3
Installed-Size: 3763
Maintainer: Debian Samba Maintainers <pkg-samba-maint@lists.alioth.debian.org>
Architecture: amd64
Depends: iproute2, psmisc, samba-libs (= 2:4.17.9+dfsg-0+deb12u3), sudo, tdb-tools, time, libbsd0 (>= 0.0), libc6 (>= 2.34), libpopt0 (>= 1.14), libtalloc2 (>= 2.3.4~), libtdb1 (>= 1.4.7~), libtevent0 (>= 0.13.0~), libtirpc3 (>= 1.0.2)
Recommends: ethtool, python3:any, python3-etcd, librados2 (>= 16.2.11+ds)
Suggests: logrotate, lsof
Description: clustered database to store temporary data
Multi-Arch: foreign
Homepage: https://www.samba.org
Description-md5: 83dff66615250b53a0cd3df6fb3b9ea7
Tag: admin::cluster, admin::file-distribution, implemented-in::c,
 interface::commandline, interface::daemon, network::hiavailability,
 network::load-balancing, protocol::ethernet, protocol::nfs,
 protocol::smb, role::program, suite::samba, use::browsing,
 use::synchronizing, works-with::software:running
Section: net
Priority: optional
Filename: pool/main/s/samba/ctdb_4.17.9+dfsg-0+deb12u3_amd64.deb
Size: 697036
MD5sum: 10d0f2b472e2866adfd94f5244463e49
SHA256: c0b3420a727d0ec23e8a8c161aa6807cd94338dcf3ac06b770299ebf481c5fbb
```

Below is an example of the information on the same package retrieved via **go-apt-search**

```
Package: ctdb
Version: 2:4.17.9+dfsg-0+deb12u3 
Architecture: amd64
Depends: [ iproute2  psmisc  samba-libs (= 2:4.17.9+dfsg-0+deb12u3)  sudo  tdb-tools  time  libbsd0 (>= 0.0)  libc6 (>= 2.34)  libpopt0 (>= 1.14)  libtalloc2 (>= 2.3.4~)  libtdb1 (>= 1.4.7~)  libtevent0 (>= 0.13.0~)  libtirpc3 (>= 1.0.2)]
Description: clustered database to store temporary data
Section: net
MD5sum:  
SHA256: c0b3420a727d0ec23e8a8c161aa6807cd94338dcf3ac06b770299ebf481c5fbb
```

Now only the essential information has been included, but all the others will be added.

## Example

```go
package main

import (
	"fmt"

	apt "github.com/Sfrisio/go-apt-search"
)

func main() {
	fullPackagesList, err := apt.AptListAll()
	if err != nil {
		panic(err)
	}
	for _, singlePackage := range fullPackagesList {
		fmt.Println("##########")
		fmt.Printf("Package: %s\n", singlePackage.PackageName)
		fmt.Printf("Version: %s\n", singlePackage.Version)
		fmt.Printf("Architecture: %s\n", singlePackage.Architecture)
		fmt.Printf("Depends: %s\n", singlePackage.Depends)
		fmt.Printf("Description: %s\n", singlePackage.Description)
		fmt.Printf("Section: %s\n", singlePackage.Section)
		fmt.Printf("MD5sum: %s\n", singlePackage.Md5sum)
		fmt.Printf("SHA256: %s\n", singlePackage.Sha256)
	}
}
```