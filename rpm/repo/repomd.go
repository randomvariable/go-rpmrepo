package RPMRepo

import (
	"encoding/xml"
)

type RepositoryMetadata struct {
	XMLName  xml.Name `xml:"repomd"`
	Revision int64    `xml:"revision"`
	Data     []Data
}

type Data struct {
	XMLName         xml.Name `xml:"data"`
	Type            string   `xml:"type,attr"`
	Location        Location
	Checksum        Checksum `xml:"checksum"`
	Timestamp       string   `xml:"timestamp"`
	OpenChecksum    Checksum `xml:"open-checksum"`
	Size            int64    `xml:"size"`
	OpenSize        int64    `xml:"open-size"`
	DatabaseVersion string   `xml:"database_version,omitempty"`
}

type Location struct {
	XMLName xml.Name `xml:"location"`
	Href    string   `xml:"href,attr"`
}

type Checksum struct {
	XMLName   xml.Name `xml:"checksum"`
	Type      string   `xml:"type,attr"`
	Checksum  string   `xml:",chardata"`
	PackageID string   `xml:"pkgid,omitempty"`
}

/*
<?xml version="1.0" encoding="UTF-8"?>
<repomd xmlns="http://linux.duke.edu/metadata/repo" xmlns:rpm="http://linux.duke.edu/metadata/rpm">
  <revision>1499286698</revision>
  <data type="primary">
    <checksum type="sha256">880055a50c05b20641530d09b23f64501a000b2f92fe252417c530178730a95e</checksum>
    <open-checksum type="sha256">9d4a10310a4092bde6231253f62976337f4e4fa3ab5cdda001e141a156a9dffb</open-checksum>
    <location href="repodata/880055a50c05b20641530d09b23f64501a000b2f92fe252417c530178730a95e-primary.xml.gz"/>
    <timestamp>1499286654</timestamp>
    <size>14750183</size>
    <open-size>134514273</open-size>
  </data>
  <data type="filelists">
    <checksum type="sha256">b429c60a20974912547de94994db64e87fd7ec20fe9acca47da30fc5ba42ce96</checksum>
    <open-checksum type="sha256">af9579b55eba28eb19b98e763b7269e47de59645091b75dbb5cb3649afda0d1e</open-checksum>
    <location href="repodata/b429c60a20974912547de94994db64e87fd7ec20fe9acca47da30fc5ba42ce96-filelists.xml.gz"/>
    <timestamp>1499286654</timestamp>
    <size>40663907</size>
    <open-size>582650352</open-size>
  </data>
  <data type="other">
    <checksum type="sha256">070edd188871859175cfae62fb963038f8ae0b108dec1445f71d620ea7f788c9</checksum>
    <open-checksum type="sha256">ad63825978f1dcffbeeba4946e0dc0eb05c749ad57a89a1b938a06e82475f53c</open-checksum>
    <location href="repodata/070edd188871859175cfae62fb963038f8ae0b108dec1445f71d620ea7f788c9-other.xml.gz"/>
    <timestamp>1499286654</timestamp>
    <size>7610619</size>
    <open-size>100013004</open-size>
  </data>
  <data type="primary_db">
    <checksum type="sha256">00bf6ab375bd31ddb0df2789417d17f38ac72a091ea39ec37d7ddeb64ccaffb2</checksum>
    <open-checksum type="sha256">ed291593c267db9c0ff530b504acd00018afd080ceead5a709020014fe783fbf</open-checksum>
    <location href="repodata/00bf6ab375bd31ddb0df2789417d17f38ac72a091ea39ec37d7ddeb64ccaffb2-primary.sqlite.bz2"/>
    <timestamp>1499286679</timestamp>
    <size>28398595</size>
   <open-size>138588160</open-size>
    <database_version>10</database_version>
  </data>
  <data type="filelists_db">
    <checksum type="sha256">192a4d58517814a4130f009abb7fc0899e06d63f78bc55f8e2429250acea6457</checksum>
    <open-checksum type="sha256">7c84685dff9dd1f3f7c9f81aad84ba57f97631822a6fd8d481df12e1b985b6e1</open-checksum>
    <location href="repodata/192a4d58517814a4130f009abb7fc0899e06d63f78bc55f8e2429250acea6457-filelists.sqlite.bz2"/>
    <timestamp>1499286698</timestamp>
    <size>41798451</size>
    <open-size>265166848</open-size>
    <database_version>10</database_version>
  </data>
  <data type="other_db">
    <checksum type="sha256">9e940b3015a8adef4956f3989ebb83c52268fdeabb00de3ebf30b1fd3a808335</checksum>
    <open-checksum type="sha256">399d8fa532962c21d81a3b5d4245e7662b0a44362fe58b3d594eae9b33835891</open-checksum>
    <location href="repodata/9e940b3015a8adef4956f3989ebb83c52268fdeabb00de3ebf30b1fd3a808335-other.sqlite.bz2"/>
    <timestamp>1499286676</timestamp>
    <size>11963226</size>
    <open-size>85381120</open-size>
    <database_version>10</database_version>
  </data>
  <data type="group">
    <checksum type="sha256">e4d5d691a5a9b53222cd8e8c3d5c99f042fdec7a27e39570a08c866e90882ec4</checksum>
    <location href="repodata/e4d5d691a5a9b53222cd8e8c3d5c99f042fdec7a27e39570a08c866e90882ec4-comps-Everything.x86_64.xml"/>
    <timestamp>1499286608</timestamp>
    <size>1895236</size>
  </data>
  <data type="group_gz">
    <checksum type="sha256">9fa6d1959b8b0a97eb8827ee899bba8ebfdb197daa29e621815d73167b2eeb8e</checksum>
    <open-checksum type="sha256">e4d5d691a5a9b53222cd8e8c3d5c99f042fdec7a27e39570a08c866e90882ec4</open-checksum>
    <location href="repodata/9fa6d1959b8b0a97eb8827ee899bba8ebfdb197daa29e621815d73167b2eeb8e-comps-Everything.x86_64.xml.gz"/>
    <timestamp>1499286654</timestamp>
    <size>453974</size>
    <open-size>1895236</open-size>
  </data>
</repomd>


*/
