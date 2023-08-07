package goaptsearch

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// absolute path to apt lists files
const aptListPath = "/var/lib/apt/lists/"

type APTPackages struct {
	PackageName  string   `json:"Package"`
	Version      string   `json:"Version"`
	Architecture string   `jsone:"Architecture"`
	Depends      []string `json:"Depends"`
	Description  string   `json:"Description"`
	Section      string   `json:"Section"`
	Md5sum       string   `json:"MD5sum"`
	Sha256       string   `json:"SHA256"`
}

// func AptSearch(packageName string) {

// }

func AptListAll() ([]APTPackages, error) {
	allPackagesFiles, errGetRepoFileList := getRepoFileList()
	if errGetRepoFileList != nil {
		return nil, errGetRepoFileList
	}
	var packagesList []APTPackages
	for _, packagesFile := range allPackagesFiles {
		readPackageFile, errOpen := os.ReadFile(filepath.Join(aptListPath, packagesFile))
		if errOpen != nil {
			return nil, errOpen
		}
		lines := strings.Split(string(readPackageFile), "\n")
		var packageNameFromList string
		var versionFromList string
		var architectureFromList string
		var dependsFromList []string
		var descriptionFromList string
		var sectionFromList string
		var md5sumFromList string
		var sha256FromList string
		for _, line := range lines {
			if strings.HasPrefix(line, "Package:") {
				packageNameFromList, _ = strings.CutPrefix(line, "Package:")
			} else if strings.HasPrefix(line, "Version:") {
				versionFromList, _ = strings.CutPrefix(line, "Version:")
			} else if strings.HasPrefix(line, "Architecture:") {
				architectureFromList, _ = strings.CutPrefix(line, "Architecture:")
			} else if strings.HasPrefix(line, "Depends:") {
				dependsList, _ := strings.CutPrefix(line, "Depends:")
				dependsFromList = strings.Split(dependsList, ",")
			} else if strings.HasPrefix(line, "Description:") {
				descriptionFromList, _ = strings.CutPrefix(line, "Description:")
			} else if strings.HasPrefix(line, "Section:") {
				sectionFromList, _ = strings.CutPrefix(line, "Section:")
			} else if strings.HasPrefix(line, "MD5sum:") {
				md5sumFromList, _ = strings.CutPrefix(line, "MD5sum:")
			} else if strings.HasPrefix(line, "SHA256:") {
				sha256FromList, _ = strings.CutPrefix(line, "SHA256:")
			} else if line == "" {
				// information dump because each new line starts a new package
				packagesList = append(packagesList, APTPackages{
					PackageName:  strings.TrimSpace(packageNameFromList),
					Version:      strings.TrimSpace(versionFromList),
					Architecture: strings.TrimSpace(architectureFromList),
					Depends:      dependsFromList,
					Description:  strings.TrimSpace(descriptionFromList),
					Section:      strings.TrimSpace(sectionFromList),
					Md5sum:       strings.TrimSpace(md5sumFromList),
					Sha256:       strings.TrimSpace(sha256FromList),
				})
			}
		}
	}
	return packagesList, nil
}

// getRepoFileList: read files from /var/lib/apt/lists and return only packages
//
// I preferred to use os.ReadDir instead of filepath.Walk because I am not interested in the list of files in the partial directory
func getRepoFileList() ([]string, error) {
	allPackagesFiles, errReadDir := os.ReadDir(aptListPath)
	if errReadDir != nil {
		return nil, errReadDir
	}
	var matchingPackagesFiles []string
	filterPackagesFile, _ := regexp.Compile(`.*\_Packages$`)
	for _, packagesFile := range allPackagesFiles {
		if filterPackagesFile.MatchString(packagesFile.Name()) {
			matchingPackagesFiles = append(matchingPackagesFiles, packagesFile.Name())
		}
	}
	return matchingPackagesFiles, nil
}
