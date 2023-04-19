test -f 洋葱加速器.dmg && rm 洋葱加速器.dmg
create-dmg \
  --volname "洋葱加速器" \
  --volicon "/Users/zz/play/shell/mac/AppIcon.icns" \
  --background "/Users/zz/play/shell/mac/dmg.png" \
  --window-pos 200 120 \
  --window-size 700 520 \
  --icon-size 100 \
  --icon "洋葱加速器.app" 150 180 \
  --app-drop-link 550 180 \
  "洋葱加速器.dmg" \
  "/Users/zz/work/onion-mac/app/"

