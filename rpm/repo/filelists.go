package RPMRepo

import "encoding/xml"

type FileLists struct {
	XMLName      xml.Name `xml:"filelists"`
	XMLNameSpace string   `xml:"xmlns,attr"`
	PackageCount int64    `xml:"packages,attr"`
	Packages     []Package
}

type File struct {
	XMLName xml.Name `xml:"file"`
	Type    string   `xml:"type,attr,omitempty"`
	Path    string   `xml:",chardata"`
}

/*
<?xml version="1.0" encoding="UTF-8"?>
<filelists xmlns="http://linux.duke.edu/metadata/filelists" packages="53912">
<package pkgid="65eb286f210cf080b497e4a529760daa7fd7d35c539392e70da1d75e4ae8b7c6" name="0ad" arch="x86_64">
  <version epoch="0" ver="0.0.21" rel="3.fc26"/>
  <file>/usr/bin/0ad</file>
  <file>/usr/bin/pyrogenesis</file>
  <file type="dir">/usr/lib64/0ad</file>
  <file>/usr/lib64/0ad/libAtlasUI.so</file>
  <file>/usr/lib64/0ad/libCollada.so</file>
  <file type="dir">/usr/share/0ad</file>
*/
