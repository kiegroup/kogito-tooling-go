#! /bin/bash


APP=../../out/macos/Kogito.app
CONTENTS=$APP/Contents
MACOS=$CONTENTS/MacOs
RESOURCES=$CONTENTS/Resources
NAME=Kogito
DMG=$NAME.dmg
APPLICATIONS=target/Applications


rm -rf ../../out/macos
rm -rf Kogito.dmg

mkdir -p $CONTENTS
mkdir $MACOS
mkdir $RESOURCES

cp ../../build/darwin/kogito $MACOS
cp src/Info.plist $CONTENTS
cp src/Kogito.png $RESOURCES
ln -s /Applications $APPLICATIONS

hdiutil create /tmp/tmp.dmg -ov -volname $NAME -fs HFS+ -srcfolder "target" 
hdiutil convert /tmp/tmp.dmg -format UDZO -o Kogito.dmg
mv Kogito.dmn ../../out
