#! /bin/bash


APP=target/Kogito.app
CONTENTS=$APP/Contents
MACOS=$CONTENTS/MacOs
RESOURCES=$CONTENTS/Resources
NAME=Kogito
DMG=$NAME.dmg
APPLICATIONS=target/Applications

rm -rf target
rm -rf Kogito.dmg

mkdir -p $CONTENTS
mkdir $MACOS
mkdir $RESOURCES

cp ../../build/darwin/runner $MACOS/kogito
cp src/Info.plist $CONTENTS
cp src/Kogito.png $RESOURCES
ln -s /Applications $APPLICATIONS

hdiutil create /tmp/tmp.dmg -ov -volname $NAME -fs HFS+ -srcfolder "target" 
hdiutil convert /tmp/tmp.dmg -format UDZO -o target/Kogito.dmg
