### Animated GIF in Pixel

```go
import (
    "github.com/aerth/spriteutil"
)

func run(){
// ...
sprite, err := spriteutil.LoadGif(file)
if err != nil { return err }
sprite.Update(dt)
sprite.Draw(win, pixel.IM)
// ...
}
```

![](assets/screenshot.gif)


Image/Font Credits:

  * https://i.giphy.com/media/gz7cJzQSlLmcE/source.gif
  * https://www.spriters-resource.com/images/light/misc/select1.gif
  * https://www.dafont.com/computerfont.font
