import poe/effect_color.{type Color}

pub type Beam {
  Beam(Color)
}

pub fn to_string(beam: Beam) -> String {
  let Beam(c) = beam
  "PlayEffect " <> effect_color.to_string(c)
}
