import actions/color.{type Color}
import actions/map_icon/map_icon.{MapIcon}
import actions/map_icon/shape
import actions/map_icon/size.{type Size}
import actions/sound

pub const tink_sound = sound.Alert(16)

pub fn circle(size: Size, color: Color) -> String {
  MapIcon(size, color, shape.Circle)
  |> map_icon.to_string
}
