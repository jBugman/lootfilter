import poe/effect_color.{type Color}
import poe/map_icon/map_icon.{MapIcon}
import poe/map_icon/shape
import poe/map_icon/size.{type Size}
import poe/sound

pub const tink_sound = sound.Alert(16)

pub fn circle(size: Size, color: Color) -> String {
  MapIcon(size, color, shape.Circle)
  |> map_icon.to_string
}
