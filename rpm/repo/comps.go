package RPMRepo

import "encoding/xml"

type Group struct {
	XMLName      xml.Name `xml:"group"`
	ID           string   `xml:"id"`
	Names        []Name
	Descriptions []Description
	Default      bool                 `xml:"default,omitempty"`
	UserVisible  bool                 `xml:"uservisible,omitempty"`
	PackageList  []PackageRequirement `xml:"packagelist>packagereq"`
}

type Name struct {
	XMLName  xml.Name `xml:"name"`
	Language string   `xml:"xml:lang,attr,omitempty"`
	Text     string   `xml:",chardata"`
}

type PackageRequirement struct {
	XMLName xml.Name `xml:"packagereq"`
	Type    string   `xml:"type,attr,omitempty"`
	Package string   `xml:",chardata"`
}

/*
comps>
  <group>
    <id>3d-printing</id>
    <name>3D Printing</name>
    <name xml:lang="bg">3D Печатане</name>
    <name xml:lang="ca">Impressió 3D</name>
    <name xml:lang="cs">3D tisk</name>
    <name xml:lang="de">3D-Druck</name>
    <name xml:lang="en_GB">3D Printing</name>
    <name xml:lang="es">Impresión 3D</name>
    <name xml:lang="fi">3D-tulostaminen</name>
    <name xml:lang="fr">Impression 3D</name>
    <name xml:lang="he">הדפסה בתלת ממד</name>
    <name xml:lang="hu">3D nyomtatás</name>
    <name xml:lang="ia">Impression 3D</name>
    <name xml:lang="id">Pencetakan 3D</name>
    <name xml:lang="it">Stampa 3D</name>
    <name xml:lang="ja">3D プリント</name>
    <name xml:lang="kk">3D баспаға шығару</name>
    <name xml:lang="nl">3D-printen</name>
    <name xml:lang="pa">3ਡੀ ਪਰਿੰਟਿੰਗ</name>
    <name xml:lang="pl">Drukowanie 3D</name>
    <name xml:lang="pt">Impressão 3D</name>
    <name xml:lang="pt_BR">Impressão 3D</name>
    <name xml:lang="ru">3D Печать</name>
    <name xml:lang="sk">3D tlač</name>
    <name xml:lang="sq">Printim 3D</name>
    <name xml:lang="sr">3D штампа</name>
    <name xml:lang="sv">3D-utskrifter</name>
    <name xml:lang="uk">Просторовий друк</name>
    <name xml:lang="zh_CN">3D 打印</name>
    <name xml:lang="zh_TW">3D 印製</name>
    <description>3D printing software</description>
    <description xml:lang="bg">Софтуер за 3D печатане</description>
    <description xml:lang="ca">Programari d'impressió 3D</description>
    <description xml:lang="cs">3Dtiskové aplikace</description>
    <description xml:lang="de">3D-Drucksoftware</description>
    <description xml:lang="en_GB">3D printing software</description>
    <description xml:lang="es">Software de impresión 3D</description>
    <description xml:lang="fi">3d-tulostusohjelmisto</description>
    <description xml:lang="fr">Logiciels d'impression 3D</description>
    <description xml:lang="he">תכנה להדפסה בתלת ממד</description>
    <description xml:lang="hu">3D nyomtatási szoftverek</description>
    <description xml:lang="ia">Software a imprimer 3D</description>
    <description xml:lang="id">Perangkat Lunak Pencetakan 3D</description>
    <description xml:lang="it">Software per stampa 3D</description>
    <description xml:lang="ja">3D プリントソフトウェア</description>
    <description xml:lang="kk">3D баспаға шығару БҚ</description>
    <description xml:lang="nl">3D-printsoftware</description>
    <description xml:lang="pa">3ਡੀ ਪਰਿੰਟਿੰਗ ਸਾਫਟਵੇਅਰ</description>
    <description xml:lang="pl">Oprogramowanie do drukowania 3D</description>
    <description xml:lang="pt">Software de impressão 3D</description>
    <description xml:lang="pt_BR">Software de impressão 3D</description>
    <description xml:lang="ru">Программы для 3D печати</description>
    <description xml:lang="sk">3D tlačové aplikácie</description>
    <description xml:lang="sq">Softuer për printimin 3D</description>
    <description xml:lang="sr">Софтвер за 3D штампу</description>
    <description xml:lang="sv">Programvara för 3D-utskrifter</description>
    <description xml:lang="uk">Програми для просторового друку</description>
    <description xml:lang="zh_CN">3D 打印软件</description>
    <description xml:lang="zh_TW">3D 印製軟體</description>
    <default>false</default>
    <uservisible>true</uservisible>
    <packagelist>
      <packagereq type="default">admesh</packagereq>
      <packagereq type="default">blender</packagereq>
      <packagereq type="default">cura</packagereq>
      <packagereq type="default">cura-lulzbot</packagereq>
      <packagereq type="default">meshlab</packagereq>
      <packagereq type="default">openscad</packagereq>
      <packagereq type="default">openscad-MCAD</packagereq>
      <packagereq type="default">printrun</packagereq>
      <packagereq type="default">slic3r</packagereq>
      <packagereq type="optional">RepetierHost</packagereq>
      <packagereq type="optional">repsnapper</packagereq>
      <packagereq type="optional">sfact</packagereq>
      <packagereq type="optional">skeinforge</packagereq>
    </packagelist>
  </group>
*/
