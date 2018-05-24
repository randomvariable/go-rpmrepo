package RPMRepo

import (
	"encoding/xml"
)

type Package struct {
	XMLName      xml.Name    `xml:"package"`
	Type         string      `xml:"type,attr,omitempty"`
	Name         string      `xml:"name"`
	Architecture string      `xml:"arch"`
	PackageID    string      `xml:"pkgid,omitempty"`
	Version      Version     `xml:"version,omitempty"`
	Checksum     Checksum    `xml:"checksum,omitempty"`
	Summary      string      `xml:"summary,omitempty"`
	Description  Description `xml:"description,omitempty"`
	Packager     string      `xml:"packager,omitempty"`
	URL          string      `xml:"url,omitempty"`
	Time         Time        `xml:"time,omitempty"`
	Size         Size        `xml:"size,omitempty"`
	Location     Location    `xml:"location,omitempty"`
	Format       Format      `xml:"format,omitempty"`
	Files        []File      `xml:",omitempty"`
}

type Description struct {
	XMLName  xml.Name `xml:"description"`
	Language string   `xml:"xml:lang,omitempty"`
	Text     string   `xml:",chardata"`
}

type Version struct {
	XMLName xml.Name `xml:"version"`
	Epoch   int64    `xml:"epoch"`
	Release string   `xml:"rel"`
	Version string   `xml:"ver"`
}

type Time struct {
	XMLName xml.Name `xml:"time"`
	Build   string   `xml:"build,attr"`
	File    string   `xml:"file,attr"`
}

type Size struct {
	XMLName   xml.Name `xml:"size"`
	Archive   int64    `xml:"archive,attr,omitempty"`
	Installed int64    `xml:"installed,attr,omitempty"`
	Package   int64    `xml:"package,attr,omitempty"`
}

type Format struct {
	XMLName     xml.Name    `xml:"format"`
	License     string      `xml:"rpm:license,omitempty"`
	Vendor      string      `xml:"rpm:vendor,omitempty"`
	Group       string      `xml:"rpm:group,omitempty"`
	BuildHost   string      `xml:"rpm:buildhost,omitempty"`
	SourceRPM   string      `xml:"rpm:sourcerpm,omitempty"`
	HeaderRange HeaderRange `xml:"header-range,omitempty"`
	Provides    []RPMEntry  `xml:"rpm:provides>rpm:entry,omitempty"`
	Requires    []RPMEntry  `xml:"rpm:requires>rpm:entry,omitempty"`
	Obsoletes   []RPMEntry  `xml:"rpm:obsoletes>rpm:entry,omitempty"`
	Files       []File
}

type HeaderRange struct {
	XMLName xml.Name `xml:"header-range"`
	Start   int64    `xml:"start,attr,omitempty"`
	End     int64    `xml:"end,attr,omitempty"`
}

type RPMEntry struct {
	XMLName xml.Name `xml:"rpm:entry"`
	Name    string   `xml:"name,attr"`
	Flags   string   `xml:"flags,attr,omitempty"`
	Epoch   int64    `xml:"epoch,attr,omitempty"`
	Release string   `xml:"rel,attr,omitempty"`
	Version string   `xml:"ver,attr,omitempty"`
}

type PackagesMetadata struct {
	XMLName      xml.Name `xml:"metadata"`
	XMLNameSpace string   `xml:"xmlns,attr"`
	PackageCount int64    `xml:"packages,attr"`
	Packages     []Package
}

/*
<metadata xmlns="http://linux.duke.edu/metadata/common" xmlns:rpm="http://linux.duke.edu/metadata/rpm" packages="53912">
<package type="rpm">
  <name>0ad</name>
  <arch>x86_64</arch>
  <version epoch="0" ver="0.0.21" rel="3.fc26"/>
  <checksum type="sha256" pkgid="YES">65eb286f210cf080b497e4a529760daa7fd7d35c539392e70da1d75e4ae8b7c6</checksum>
  <summary>Cross-Platform RTS Game of Ancient Warfare</summary>
  <description>0 A.D. (pronounced "zero ey-dee") is a free, open-source, cross-platform
real-time strategy (RTS) game of ancient warfare. In short, it is a
historically-based war/economy game that allows players to relive or rewrite
the history of Western civilizations, focusing on the years between 500 B.C.
and 500 A.D. The project is highly ambitious, involving state-of-the-art 3D
graphics, detailed artwork, sound, and a flexible and powerful custom-built
game engine.

The game has been in development by Wildfire Games (WFG), a group of volunteer,
hobbyist game developers, since 2001.</description>
  <packager>Fedora Project</packager>
  <url>http://play0ad.com</url>
  <time file="1486412865" build="1485538738"/>
  <size package="3751914" installed="11616900" archive="11645632"/>
  <location href="Packages/0/0ad-0.0.21-3.fc26.x86_64.rpm"/>
  <format>
    <rpm:license>GPLv2+ and BSD and MIT and IBM</rpm:license>
    <rpm:vendor>Fedora Project</rpm:vendor>
    <rpm:group>Amusements/Games</rpm:group>
    <rpm:buildhost>buildvm-09.phx2.fedoraproject.org</rpm:buildhost>
    <rpm:sourcerpm>0ad-0.0.21-3.fc26.src.rpm</rpm:sourcerpm>
    <rpm:header-range start="4424" end="42262"/>
    <rpm:provides>
      <rpm:entry name="0ad" flags="EQ" epoch="0" ver="0.0.21" rel="3.fc26"/>
      <rpm:entry name="0ad(x86-64)" flags="EQ" epoch="0" ver="0.0.21" rel="3.fc26"/>
      <rpm:entry name="appdata()"/>
      <rpm:entry name="appdata(0ad.appdata.xml)"/>
      <rpm:entry name="application()"/>
      <rpm:entry name="application(0ad.desktop)"/>
      <rpm:entry name="libAtlasUI.so()(64bit)"/>
      <rpm:entry name="libCollada.so()(64bit)"/>
    </rpm:provides>
    <rpm:requires>
      <rpm:entry name="/bin/sh"/>
      <rpm:entry name="0ad-data" flags="EQ" epoch="0" ver="0.0.21"/>
      <rpm:entry name="libGL.so.1()(64bit)"/>
      <rpm:entry name="libSDL2-2.0.so.0()(64bit)"/>
      <rpm:entry name="libX11.so.6()(64bit)"/>
      <rpm:entry name="libXcursor.so.1()(64bit)"/>
      <rpm:entry name="libboost_filesystem.so.1.63.0()(64bit)"/>
      <rpm:entry name="libcurl.so.4()(64bit)"/>
      <rpm:entry name="libdl.so.2()(64bit)"/>
      <rpm:entry name="libdl.so.2(GLIBC_2.2.5)(64bit)"/>
      <rpm:entry name="libenet.so.7()(64bit)"/>
      <rpm:entry name="libgcc_s.so.1()(64bit)"/>
      <rpm:entry name="libgcc_s.so.1(GCC_3.0)(64bit)"/>
      <rpm:entry name="libgcc_s.so.1(GCC_3.4)(64bit)"/>
      <rpm:entry name="libgloox.so.13()(64bit)"/>
      <rpm:entry name="libicui18n.so.57()(64bit)"/>
      <rpm:entry name="libicuuc.so.57()(64bit)"/>
      <rpm:entry name="libm.so.6()(64bit)"/>
      <rpm:entry name="libm.so.6(GLIBC_2.2.5)(64bit)"/>
      <rpm:entry name="libminiupnpc.so.16()(64bit)"/>
      <rpm:entry name="libmozjs-38.so()(64bit)"/>
      <rpm:entry name="libmozjs-38.so(js)(64bit)"/>
      <rpm:entry name="libnvtt.so.2.0()(64bit)"/>
      <rpm:entry name="libopenal.so.1()(64bit)"/>
      <rpm:entry name="libpng16.so.16()(64bit)"/>
      <rpm:entry name="libpng16.so.16(PNG16_0)(64bit)"/>
      <rpm:entry name="libpthread.so.0()(64bit)"/>
      <rpm:entry name="libpthread.so.0(GLIBC_2.2.5)(64bit)"/>
      <rpm:entry name="libpthread.so.0(GLIBC_2.3.2)(64bit)"/>
      <rpm:entry name="librt.so.1()(64bit)"/>
      <rpm:entry name="librt.so.1(GLIBC_2.2.5)(64bit)"/>
      <rpm:entry name="libstdc++.so.6()(64bit)"/>
      <rpm:entry name="libstdc++.so.6(CXXABI_1.3)(64bit)"/>
      <rpm:entry name="libstdc++.so.6(CXXABI_1.3.5)(64bit)"/>
      <rpm:entry name="libstdc++.so.6(CXXABI_1.3.8)(64bit)"/>
      <rpm:entry name="libstdc++.so.6(CXXABI_1.3.9)(64bit)"/>
      <rpm:entry name="libstdc++.so.6(GLIBCXX_3.4)(64bit)"/>
      <rpm:entry name="libstdc++.so.6(GLIBCXX_3.4.11)(64bit)"/>
      <rpm:entry name="libstdc++.so.6(GLIBCXX_3.4.14)(64bit)"/>
      <rpm:entry name="libstdc++.so.6(GLIBCXX_3.4.15)(64bit)"/>
      <rpm:entry name="libstdc++.so.6(GLIBCXX_3.4.18)(64bit)"/>
      <rpm:entry name="libstdc++.so.6(GLIBCXX_3.4.20)(64bit)"/>
      <rpm:entry name="libstdc++.so.6(GLIBCXX_3.4.21)(64bit)"/>
      <rpm:entry name="libstdc++.so.6(GLIBCXX_3.4.9)(64bit)"/>
      <rpm:entry name="libvorbisfile.so.3()(64bit)"/>
      <rpm:entry name="libwx_baseu-2.8.so.0()(64bit)"/>
      <rpm:entry name="libwx_baseu-2.8.so.0(WXU_2.8)(64bit)"/>
      <rpm:entry name="libwx_baseu_xml-2.8.so.0()(64bit)"/>
      <rpm:entry name="libwx_baseu_xml-2.8.so.0(WXU_2.8)(64bit)"/>
      <rpm:entry name="libwx_gtk2u_core-2.8.so.0()(64bit)"/>
      <rpm:entry name="libwx_gtk2u_core-2.8.so.0(WXU_2.8)(64bit)"/>
      <rpm:entry name="libwx_gtk2u_gl-2.8.so.0()(64bit)"/>
      <rpm:entry name="libwx_gtk2u_gl-2.8.so.0(WXU_2.8)(64bit)"/>
      <rpm:entry name="libxml2.so.2()(64bit)"/>
      <rpm:entry name="libxml2.so.2(LIBXML2_2.4.30)(64bit)"/>
      <rpm:entry name="libxml2.so.2(LIBXML2_2.5.2)(64bit)"/>
      <rpm:entry name="libxml2.so.2(LIBXML2_2.6.0)(64bit)"/>
      <rpm:entry name="libxml2.so.2(LIBXML2_2.6.21)(64bit)"/>
      <rpm:entry name="libxml2.so.2(LIBXML2_2.9.0)(64bit)"/>
      <rpm:entry name="libz.so.1()(64bit)"/>
      <rpm:entry name="libz.so.1(ZLIB_1.2.0)(64bit)"/>
      <rpm:entry name="rtld(GNU_HASH)"/>
      <rpm:entry name="libc.so.6(GLIBC_2.15)(64bit)"/>
    </rpm:requires>
    <file>/usr/bin/0ad</file>
    <file>/usr/bin/pyrogenesis</file>
  </format>
</package>
*/
