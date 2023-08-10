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

The search can be performed in a precise manner, indicating the exact name of the package to be searched for, or search for all packages that include a certain string in their name.
In the example case a search is performed for **0ad** indicating that this is not a precise search

```go
package main

import (
	"fmt"

	apt "github.com/Sfrisio/go-apt-search"
)

func main() {
	fullPackagesList, errAptListAll := apt.AptListAll()
	if errAptListAll != nil {
		panic(errAptListAll)
	}
	searchResult, errAptSearch := apt.AptSearch("0ad", fullPackagesList, false)
	if errAptSearch != nil {
		panic(errAptSearch)
	}
	for _, singlePackage := range searchResult {
		fmt.Printf("\n### %s ###\n", singlePackage.PackageName)
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

Result:

```
### 0ad ###
Version: 0.0.26-3
Architecture: amd64
Depends: [ 0ad-data (>= 0.0.26)  0ad-data (<= 0.0.26-3)  0ad-data-common (>= 0.0.26)  0ad-data-common (<= 0.0.26-3)  libboost-filesystem1.74.0 (>= 1.74.0)  libc6 (>= 2.34)  libcurl3-gnutls (>= 7.32.0)  libenet7  libfmt9 (>= 9.1.0+ds1)  libfreetype6 (>= 2.2.1)  libgcc-s1 (>= 3.4)  libgloox18 (>= 1.0.24)  libicu72 (>= 72.1~rc-1~)  libminiupnpc17 (>= 1.9.20140610)  libopenal1 (>= 1.14)  libpng16-16 (>= 1.6.2-1)  libsdl2-2.0-0 (>= 2.0.12)  libsodium23 (>= 1.0.14)  libstdc++6 (>= 12)  libvorbisfile3 (>= 1.1.2)  libwxbase3.2-1 (>= 3.2.1+dfsg)  libwxgtk-gl3.2-1 (>= 3.2.1+dfsg)  libwxgtk3.2-1 (>= 3.2.1+dfsg-2)  libx11-6  libxml2 (>= 2.9.0)  zlib1g (>= 1:1.2.0)]
Description: Real-time strategy game of ancient warfare
Section: games
MD5sum: 4d471183a39a3a11d00cd35bf9f6803d
SHA256: 3a2118df47bf3f04285649f0455c2fc6fe2dc7f0b237073038aa00af41f0d5f2

### 0ad-data ###
Version: 0.0.26-1
Architecture: all
Depends: [ 0ad-data (>= 0.0.26)  0ad-data (<= 0.0.26-3)  0ad-data-common (>= 0.0.26)  0ad-data-common (<= 0.0.26-3)  libboost-filesystem1.74.0 (>= 1.74.0)  libc6 (>= 2.34)  libcurl3-gnutls (>= 7.32.0)  libenet7  libfmt9 (>= 9.1.0+ds1)  libfreetype6 (>= 2.2.1)  libgcc-s1 (>= 3.4)  libgloox18 (>= 1.0.24)  libicu72 (>= 72.1~rc-1~)  libminiupnpc17 (>= 1.9.20140610)  libopenal1 (>= 1.14)  libpng16-16 (>= 1.6.2-1)  libsdl2-2.0-0 (>= 2.0.12)  libsodium23 (>= 1.0.14)  libstdc++6 (>= 12)  libvorbisfile3 (>= 1.1.2)  libwxbase3.2-1 (>= 3.2.1+dfsg)  libwxgtk-gl3.2-1 (>= 3.2.1+dfsg)  libwxgtk3.2-1 (>= 3.2.1+dfsg-2)  libx11-6  libxml2 (>= 2.9.0)  zlib1g (>= 1:1.2.0)]
Description: Real-time strategy game of ancient warfare (data files)
Section: games
MD5sum: fc5ed8a20ce1861950c7ed3a5a615be0
SHA256: 53745ae74d05bccf6783400fa98f3932b21729ab9d2e86151aa2c331c3455178

### 0ad-data-common ###
Version: 0.0.26-1
Architecture: all
Depends: [ fonts-dejavu-core | ttf-dejavu-core  fonts-freefont-ttf | ttf-freefont  fonts-texgyre | tex-gyre]
Description: Real-time strategy game of ancient warfare (common data files)
Section: games
MD5sum: 7ce70dc6e6de01134d2e199499fd3925
SHA256: 0a40074c844a304688e503dd0c3f8b04e10e40f6f81b8bad260e07c54aa37864
```

Search filtering by reposipotry

```go
package main

import (
	"fmt"

	apt "github.com/Sfrisio/go-apt-search"
)

func main() {
	availableRepo, errGetAvailableRepo := apt.GetAvailableRepo()
	if errGetAvailableRepo != nil {
		panic(errGetAvailableRepo)
	}
	var myRepo []apt.RepoArchive
	for _, repo := range availableRepo {
		if strings.Contains(repo.Domain, "deb.debian.org") {
			myRepo = append(myRepo, repo)
		}
	}
	searchResult, errAptSearch := apt.AptSearch("synaptic", myRepo, false)
	if errAptSearch != nil {
		panic(errAptSearch)
	}
	for _, singlePackage := range searchResult {
		fmt.Printf("\n### %s ###\n", singlePackage.PackageName)
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

List all available packages:

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
		fmt.Printf("\n###: %s ###\n", singlePackage.PackageName)
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