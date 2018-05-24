package RPMRepo

import "encoding/xml"

type OtherPackage struct {
	XMLName      xml.Name `xml:"package"`
	Name         string   `xml:"name,attr"`
	Architecture string   `xml:"arch,attr"`
	PackageID    string   `xml:"pkgid,attr,omitempty"`
	Version      Version
}

type ChangeLog struct {
	XMLName xml.Name `xml:"changelog"`
	Author  string   `xml:"author,attr"`
	Date    string   `xml:"author,attr"`
	Text    string   `xml:",chardata"`
}

/*
<otherdata xmlns="http://linux.duke.edu/metadata/other" packages="53912">
<package pkgid="65eb286f210cf080b497e4a529760daa7fd7d35c539392e70da1d75e4ae8b7c6" name="0ad" arch="x86_64">
  <version epoch="0" ver="0.0.21" rel="3.fc26"/>
  <changelog author="Igor Gnatenko &lt;i.gnatenko.brain@gmail.com&gt; - 0.0.19" date="1448712000">- 0.0.19</changelog>
  <changelog author="Jonathan Wakely &lt;jwakely@redhat.com&gt; - 0.0.19-2" date="1452772800">- Rebuilt for Boost 1.60</changelog>
  <changelog author="Fedora Release Engineering &lt;releng@fedoraproject.org&gt; - 0.0.19-3" date="1454500800">- Rebuilt for https://fedoraproject.org/wiki/Fedora_24_Mass_Rebuild</changelog>
  <changelog author="Igor Gnatenko &lt;ignatenko@redhat.com&gt; - 0.0.20-1" date="1459598400">- Update to 0.0.20</changelog>
  <changelog author="Peter Robinson &lt;pbrobinson@fedoraproject.org&gt; 0.0.20-2" date="1460203200">- Upstream now supports aarch64 (tests currently fail)</changelog>
  <changelog author="David Tardon &lt;dtardon@redhat.com&gt; - 0.0.20-3" date="1460721600">- rebuild for ICU 57.1</changelog>
  <changelog author="pcpa &lt;paulo.cesar.pereira.de.andrade@gmail.com&gt; - 0.0.20-4" date="1466769600">- Rebuild for miniupnpc 2.0</changelog>
  <changelog author="Igor Gnatenko &lt;i.gnatenko.brain@gmail.com&gt; - 0.0.21-1" date="1478692800">- Update to 0.0.21</changelog>
  <changelog author="Jonathan Wakely &lt;jwakely@redhat.com&gt; - 0.0.21-2" date="1485518400">- Rebuilt for Boost 1.63</changelog>
  <changelog author="Jonathan Wakely &lt;jwakely@redhat.com&gt; - 0.0.21-3" date="1485518401">- Rebuilt for Boost 1.63</changelog>
</package>
*/
